# ==========================================================================
#  Core: variable.tf
# --------------------------------------------------------------------------
#  Description:
#    Global Variable
# --------------------------------------------------------------------------
#    - AWS Region
#    - AWS Account ID
#    - AWS Account Profile
#    - Workspace ID
#    - Workspace Environment
#    - Global Tags
#    - Terraform State S3 Bucket Name
#    - Terraform State S3 DynamoDB Table
#    - Terraform State S3 Key (Prefix)
#    - Terraform State S3 Encryption
# ==========================================================================

# --------------------------------------------------------------------------
#  AWS
# --------------------------------------------------------------------------
variable "aws_region" {
  description = "The AWS region to deploy the VPC certificate in"
  type        = string
  default     = "us-west-2"
}

variable "aws_account_id" {
  description = "The AWS Account ID to deploy the VPC certificate in"
  type        = string
  default     = "YOUR_AWS_ACCOUNT"
}

variable "aws_account_profile" {
  description = "The AWS Profile to deploy the VPC certificate in"
  type        = string
  default     = "devopscorner"
}

# --------------------------------------------------------------------------
#  Workspace
# --------------------------------------------------------------------------
variable "workspace_name" {
  description = "Workspace Environment Name"
  type        = string
  default     = "default"
}

variable "workspace_env" {
  description = "Workspace Environment Selection"
  type        = map(string)
  default = {
    default = "default"
    lab     = "rnd"
    staging = "staging"
    nonprod = "nonprod"
    prod    = "prod"
  }
}

# --------------------------------------------------------------------------
#  Environment Resources Tags
# --------------------------------------------------------------------------
variable "environment" {
  description = "Target Environment (tags)"
  type        = map(string)
  default = {
    default = "DEF"
    lab     = "RND"
    staging = "STG"
    nonprod = "NONPROD"
    prod    = "PROD"
  }
}

# --------------------------------------------------------------------------
#  Department Tags
# --------------------------------------------------------------------------
variable "department" {
  description = "Department Owner"
  type        = string
  default     = "DEVOPS"
}

# --------------------------------------------------------------------------
#  Bucket Terraform State
# --------------------------------------------------------------------------
variable "tfstate_bucket" {
  description = "Name of bucket to store tfstate"
  type        = string
  default     = "devopscorner-tf-remote-state"
}

variable "tfstate_dynamodb_table" {
  description = "Name of dynamodb table to store tfstate"
  type        = string
  default     = "devopscorner-tf-state-lock"
}

variable "tfstate_path" {
  description = "Path .tfstate in Bucket"
  type        = string
  default     = "core/terraform.tfstate"
}

variable "tfstate_encrypt" {
  description = "Name of bucket to store tfstate"
  type        = bool
  default     = true
}
