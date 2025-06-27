resource "aws_s3_bucket" "b" {
  bucket = var.bucket_name
}

resource "aws_s3_bucket_public_access_block" "acl" {
  bucket                  = aws_s3_bucket.b.id
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
  depends_on              = [aws_s3_bucket.b]
}

resource "aws_s3_bucket_policy" "p" {
  bucket = aws_s3_bucket.b.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowCloudFrontServicePrincipalReadOnly"
        Effect    = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action    = "s3:GetObject"
        Resource  = "${aws_s3_bucket.b.arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = aws_cloudfront_distribution.frontend.arn
          }
        }
      }
    ]
  })
  depends_on = [aws_s3_bucket_public_access_block.acl, aws_cloudfront_distribution.frontend]
}

resource "local_file" "env_production" {
  content  = <<EOT
VITE_API_URL=https://${var.backend_domain}
EOT
  filename = "../../frontend/go-next/.env.production"
}

# build + upload
resource "null_resource" "build_and_sync" {
  provisioner "local-exec" {
    command = <<EOT
      cd ../../frontend/go-next
      npm ci
      npm run build
      aws s3 sync dist/ s3://${aws_s3_bucket.b.bucket}/ --delete
    EOT
  } # port to change on backend and up here

  triggers = {
    always_run = timestamp()
  }

  depends_on = [aws_s3_bucket_policy.p, local_file.env_production]
}

resource "aws_s3_bucket_website_configuration" "website" {
  bucket = aws_s3_bucket.b.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }

  depends_on = [null_resource.build_and_sync]
}

output "frontend_endpoint" {
  value = aws_s3_bucket_website_configuration.website.website_endpoint
}
