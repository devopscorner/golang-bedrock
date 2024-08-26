# ==========================================================================
#  Core: backend.tf
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
    region         = "us-west-2"
    bucket         = "devopscorner-tf-remote-state"
    dynamodb_table = "devopscorner-tf-state-lock"
    key            = "core/terraform.tfstate"
    encrypt        = true
  }
}
