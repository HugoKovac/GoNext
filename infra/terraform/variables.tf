variable "bucket_name" {
  description = "aws bucket s3 name for frontend static website"
  type = string
  default = "gonext-frontend-bucket"
}

variable "static_website_build_dir" {
  description = "local path for frontend static website"
  type = string
  default = "../../frontend/go-next/dist"
}

variable "public_key" {
  description = "Public key for SSH access"
}
