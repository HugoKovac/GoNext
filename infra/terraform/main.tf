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
  
}

module "frontend" {
  source = "./frontend"

  backend_host = module.backend.backend_host
}
