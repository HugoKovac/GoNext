variable "aws_region" {
  description = "AWS Region for resources"
  type        = string
  default = "eu-west-3"

}

variable "public_key" {
  description = "Public key for SSH access"
  type = string
}

