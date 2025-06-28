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

    environment = {
      GOOS        = var.build_env.GOOS
      GOARCH      = var.build_env.GOARCH
      CGO_ENABLED = var.build_env.CGO_ENABLED
    }
  }
}

# Upload the binary to the EC2 instance
resource "null_resource" "deploy_binary" {
  depends_on = [aws_instance.app_server, null_resource.build_binary, aws_db_instance.psql_instance]

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

  provisioner "file" {
    content = templatefile("${path.module}/template/.env.tpl", {
      host      = aws_db_instance.psql_instance.address
      port      = aws_db_instance.psql_instance.port
      username      = aws_db_instance.psql_instance.username
      password  = aws_db_instance.psql_instance.password
      db_name      = aws_db_instance.psql_instance.db_name
      domain = var.domain
      frontend_domain = var.frontend_domain
      allow_origins = var.frontend_domain
    })
    destination = "/tmp/.prodenv"
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
    content = templatefile("${path.module}/template/service.tpl", {
      binary_name = var.binary_name
      description = var.service_description
      user        = var.service_user
      port        = var.service_port
      env_file    = "/tmp/.prodenv"
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
