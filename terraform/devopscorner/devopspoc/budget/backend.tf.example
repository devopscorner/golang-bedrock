# ==========================================================================
#  Budget: backend.tf
# --------------------------------------------------------------------------
#  Description:
#    Store Terraform State to S3
# --------------------------------------------------------------------------
#    - S3 Bucket Path
#    - DynamoDB Table
# ==========================================================================

# --------------------------------------------------------------------------
#  Store Path for Terraform State
# --------------------------------------------------------------------------
terraform {
  backend "s3" {
    region         = var.aws_region
    bucket         = var.tfstate_bucket
    dynamodb_table = var.tfstate_dynamodb_table
    key            = var.tfstate_path
    encrypt        = var.tfstate_encrypt
  }
}
