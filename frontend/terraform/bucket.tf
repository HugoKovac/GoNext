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

resource "aws_s3_object" "static_site_upload_object" {
  for_each = fileset(var.static_website_build_dir, "**")
  bucket = aws_s3_bucket.b.id
  key = each.value
  source = "${var.static_website_build_dir}/${each.value}"
  etag = filemd5("${var.static_website_build_dir}/${each.value}")
   content_type = lookup({
    ".html" = "text/html"
    ".css"  = "text/css"
    ".js"   = "application/javascript"
    ".json" = "application/json"
    ".png"  = "image/png"
    ".jpg"  = "image/jpeg"
    ".jpeg" = "image/jpeg"
    ".gif"  = "image/gif"
    ".svg"  = "image/svg+xml"
    ".ico"  = "image/x-icon"
    ".txt"  = "text/plain"
    ".pdf"  = "application/pdf"
    ".woff" = "font/woff"
    ".woff2" = "font/woff2"
    ".ttf"  = "font/ttf"
    ".otf"  = "font/otf"
  }, ".${reverse(split(".", each.value))[0]}", "application/octet-stream")
  depends_on = [ aws_s3_bucket.b ]
}

resource "aws_s3_bucket_website_configuration" "website" {
  bucket = aws_s3_bucket.b.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "index.html"
  }

  depends_on = [ aws_s3_object.static_site_upload_object ]
}

