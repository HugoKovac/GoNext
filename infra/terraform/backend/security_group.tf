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

resource "aws_security_group" "db_sg" {
  name_prefix = "db-sg-"

  ingress {
    from_port       = 5432
    to_port         = 5432
    protocol        = "tcp"
    security_groups = [aws_security_group.allow_access.id]
  }
  
  tags = {
    Name = "database-security-group"
  }
}
