# Key pair for EC2 access
resource "aws_key_pair" "deployer" {
  key_name   = "deployer-key"
  public_key = var.public_key
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
