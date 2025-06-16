output "instance_id" {
  value = aws_instance.app_server.id
}

# Outputs
output "instance_ip" {
  description = "Public IP of the EC2 instance"
  value       = aws_instance.app_server.public_ip
}

output "instance_dns" {
  description = "Public DNS name of the EC2 instance"
  value       = aws_instance.app_server.public_dns
}

output "ssh_command" {
  description = "SSH command to connect to the instance"
  value       = "ssh -i ${var.private_key_path} ubuntu@${aws_instance.app_server.public_ip}"
}

output "service_url" {
  description = "URL to access the service"
  value       = "http://${aws_instance.app_server.public_ip}:8080"
}