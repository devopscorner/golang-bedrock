# ==========================================================================
#  EKS: provider.tf
# --------------------------------------------------------------------------
#  Description:
#    Provider Modules
# --------------------------------------------------------------------------
#    - Terraform Cli Version
#    - AWS Terraform Version
#    - AWS Region
#    - AWS Access Key
#    - AWS Secret Key
#    - AWS Account ID
# ==========================================================================

# --------------------------------------------------------------------------
#  Terraform AWS Version Compability & Remote State Path
# --------------------------------------------------------------------------
terraform {
  required_version = ">= 1.5.0"

  required_providers {
    aws = {
      source = "opentofu/aws"
      # source = "hashicorp/aws"
      version = ">= 4.50.0, < 5.7"
    }
  }
}

# --------------------------------------------------------------------------
#  AWS Provider Properties
# --------------------------------------------------------------------------
provider "aws" {
  region = var.aws_region

  # access_key = file("/etc/terraform/aws/credentials/access-key")
  # secret_key = file("/etc/terraform/aws/credentials/secret-key")

  ## version >= 3.63.0, < 4.0
  # shared_credentials_file  = "$HOME/.aws/devopscorner/credentials"
  # shared_credentials_file  = "/etc/terraform/aws/shared/credentials"
  # profile                  = var.aws_account_profile

  ## version >= 4.0
  # shared_config_files      = ["/etc/terraform/aws/shared/config"]
  # shared_credentials_files = ["/etc/terraform/aws/shared/credentials"]
  # profile                  = var.aws_account_profile

  ## version >= 4.0
  # assume_role {
  #   role_arn     = "arn:aws:iam:aws_region:123456789012:role/ROLE_NAME"
  #   session_name = "SESSION_NAME"
  #   external_id  = "EXTERNAL_ID"
  # }
}
