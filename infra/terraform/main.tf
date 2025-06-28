terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

module "backend" {
  source = "./backend"

  public_key = var.public_key
  frontend_endpoint = module.frontend.frontend_endpoint
  backend_domain = var.backend_domain
  frontend_domain = var.frontend_domain
  domain = var.domain
  
}

module "frontend" {
  source = "./frontend"

  backend_domain = var.backend_domain
  frontend_domain = var.frontend_domain
  domain = var.domain

}

output "frontend_endpoint" {
  description = "The endpoint of the frontend application"
  value       = module.frontend.frontend_endpoint
}
