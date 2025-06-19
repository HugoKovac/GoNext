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

module "frontend" {
  source = "./frontend"
}

module "backend" {
  source = "./backend"

  public_key = var.public_key
}

