resource "aws_s3_bucket" "b" {
  bucket = var.bucket_name
}

resource "aws_s3_bucket_public_access_block" "acl" {
  bucket = aws_s3_bucket.b.id
  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
  depends_on = [ aws_s3_bucket.b ]
}

resource "aws_s3_bucket_policy" "p" {
  bucket = aws_s3_bucket.b.id
  policy = jsonencode({
    Version = "2012-10-17"
    Id      = "MYBUCKETPOLICY"
    Statement = [
      {
        Sid       = "Statement1"
        Effect    = "Allow"
        Principal = "*"
        Action    = "s3:GetObject"
        Resource  = "${aws_s3_bucket.b.arn}/*"
      }
    ]
  })
  depends_on = [ aws_s3_bucket_public_access_block.acl ]
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
  }

  triggers = {
    always_run = timestamp()
  }

  depends_on = [ aws_s3_bucket_policy.p ]
}

resource "aws_s3_bucket_website_configuration" "website" {
  bucket = aws_s3_bucket.b.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }

  depends_on = [ null_resource.build_and_sync ]
}

