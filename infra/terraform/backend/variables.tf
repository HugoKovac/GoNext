variable "public_key" {
  description = "Public key for SSH access"
  type = string
}

variable "private_key_path" {
  description = "Path to private key for SSH access"
  default = "~/.ssh/id_rsa"
  type = string
}


variable "binary_name" {
  description = "Name for the binary on the remote server"
  type        = string
  default     = "app-server"
}

variable "service_description" {
  description = "Description for the systemd service"
  type        = string
  default     = "Go Next Backend Service"
}

variable "service_user" {
  description = "User to run the service as"
  type        = string
  default     = "ubuntu"
}

variable "service_port" {
  description = "Port the service runs on"
  type        = number
  default     = 8080
}

variable "source_dir" {
  description = "Path to the Go source code directory"
  type        = string
  default     = "../../backend"
}

variable "go_main_path" {
  description = "Go main to compile"
  type        = string
  default     = "./cmd/http/main.go"
}

variable "build_env" {
  description = "Environment variables for Go build"
  type        = map(string)
  default = {
    GOOS   = "linux"
    GOARCH = "amd64"
    CGO_ENABLED = "0"
  }
}

variable "domain" {
  type        = string
}

variable "frontend_endpoint" {
  type = string
}

variable "backend_domain" {
  type = string
}

variable "frontend_domain" {
  type = string
}
