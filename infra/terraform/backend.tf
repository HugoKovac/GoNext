# Key pair for EC2 access
resource "aws_key_pair" "deployer" {
  key_name   = "deployer-key"
  public_key = var.public_key
}


# Security group allowing SSH and application access
resource "aws_security_group" "allow_access" {
  name_prefix = "app-server-sg"
  
  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  
  ingress {
    description = "Application port"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "GoNextBackend-SG"
  }
}

# EC2 instance
resource "aws_instance" "app_server" {
  ami           = "ami-04a20f4b19f8a7a88"  # Amazon Linux 2
  instance_type = "t2.micro"
  key_name      = aws_key_pair.deployer.key_name
  
  vpc_security_group_ids = [aws_security_group.allow_access.id]
  
  tags = {
    Name = "GoNextBackend"
  }
}

# Build the Go binary
resource "null_resource" "build_binary" {
  # Trigger rebuild when source code changes
  triggers = {
    source_hash = sha256(join("", [
      for f in fileset(var.source_dir, "**/*.go") : filesha256("${var.source_dir}/${f}")
    ]))
    go_mod_hash = fileexists("${var.source_dir}/go.mod") ? filesha256("${var.source_dir}/go.mod") : ""
    go_sum_hash = fileexists("${var.source_dir}/go.sum") ? filesha256("${var.source_dir}/go.sum") : ""
  }

  provisioner "local-exec" {
    command = "cd ${var.source_dir} && go build -o ${var.binary_name} ${var.go_main_path}"
    
    environment = var.build_env
  }
}

# Upload the binary to the EC2 instance
resource "null_resource" "deploy_binary" {
  depends_on = [aws_instance.app_server, null_resource.build_binary]

  # Trigger re-deployment when binary or instance changes
  triggers = {
    build_trigger = null_resource.build_binary.id
    instance_id   = aws_instance.app_server.id
  }

  # Copy binary to server
  provisioner "file" {
    source      = "${var.source_dir}/${var.binary_name}"
    destination = "/tmp/${var.binary_name}"

    connection {
      type        = "ssh"
      user        = var.service_user
      private_key = file(var.private_key_path)
      host        = aws_instance.app_server.public_dns
      timeout     = "5m"
    }
  }

  # Copy systemd service file
  provisioner "file" {
    content = templatefile("${path.module}/service.tpl", {
      binary_name = var.binary_name
      description = var.service_description
      user        = var.service_user
      port        = var.service_port
    })
    destination = "/tmp/${var.binary_name}.service"

    connection {
      type        = "ssh"
      user        = var.service_user
      private_key = file(var.private_key_path)
      host        = aws_instance.app_server.public_dns
      timeout     = "5m"
    }
  }

  # Set up and start the binary
  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/${var.binary_name}",
      "sudo mv /tmp/${var.binary_name} /usr/local/bin/",
      "sudo chown root:root /usr/local/bin/${var.binary_name}",
      
      # Install systemd service file
      "sudo mv /tmp/${var.binary_name}.service /etc/systemd/system/",
      "sudo chown root:root /etc/systemd/system/${var.binary_name}.service",
      
      # Enable and start the service
      "sudo systemctl daemon-reload",
      "sudo systemctl enable ${var.binary_name}",
      "sudo systemctl start ${var.binary_name}",
      "sudo systemctl status ${var.binary_name}"
    ]

    connection {
      type        = "ssh"
      user        = var.service_user
      private_key = file(var.private_key_path)
      host        = aws_instance.app_server.public_dns
      timeout     = "5m"
    }
  }
}