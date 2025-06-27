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
  
}

module "frontend" {
  source = "./frontend"

  backend_domain = var.backend_domain

}

output "frontend_dns_instructions" {
  description = "Run the following command to get the DNS name of the frontend bucket. You will need to create a CNAME record in your DNS provider pointing your frontend domain to this value."
  value = "aws s3api get-bucket-website --bucket $(terraform output -raw frontend_endpoint | cut -d'/' -f1) --query 'WebsiteConfiguration.WebsiteEndpoint' --output text"
}

output "backend_dns_instructions" {
  description = "Run the following command to get the DNS name of the backend ALB. You will need to create a CNAME record in your DNS provider pointing your backend domain to this value."
  value = "terraform output -raw backend_alb_dns_name"
}
