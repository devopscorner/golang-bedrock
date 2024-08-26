# ==========================================================================
#  Core: core-vars.tf
# --------------------------------------------------------------------------
#  Description
#    Core Infrastructure Spesific Variable
# --------------------------------------------------------------------------
#    - Core Prefix Name
#    - Core VPC CIDR Block
#    - Core VPC Peer
#    - Core VPC Peer Owner ID
#    - Core VPC Gateway Propagating
#    - Core VPC CIDR Secondary Zone A
#    - Core VPC CIDR Secondary Zone B
#    - Core Prefix EC2
#    - Core Prefix EKS
#    - Core Prefix NAT EC2
#    - Core Prefix NAT EKS
# ==========================================================================

# --------------------------------------------------------------------------
#  Prefix Infra
# --------------------------------------------------------------------------
variable "coreinfra" {
  description = "Core Infrastrucre Name Prefix"
  type        = string
  default     = "devopscorner-tf"
}

# --------------------------------------------------------------------------
#  VPC
# --------------------------------------------------------------------------
variable "vpc_cidr" {
  description = "Core Infrastrucre CIDR Block"
  type        = map(string)
  default = {
    default = "10.16.0.0/16"
    lab     = "10.16.0.0/16"
    staging = "10.32.0.0/16"
    nonprod = "10.32.0.0/16"
    prod    = "10.48.0.0/16"
  }
}

variable "vpc_peer" {
  description = "Core Infrastrucre VPC Peers ID"
  type        = map(string)
  default = {
    default = "vpc-1234567890"
    lab     = "vpc-1234567890"
    staging = "vpc-1234567890"
    nonprod = "vpc-1234567890"
    prod    = "vpc-0987654321"
  }
}

variable "peer_owner_id" {
  description = "Core Infrastrucre VPC Peers Owner ID"
  type        = map(string)
  default = {
    default = "1234567890"
    lab     = "1234567890"
    staging = "1234567890"
    nonprod = "1234567890"
    prod    = "0987654321"
  }
}

variable "propagating_vgws" {
  description = "Core Infrastrucre VPC Gateway Propagating"
  type        = map(string)
  default = {
    default = "vgw-1234567890"
    lab     = "vgw-1234567890"
    staging = "vgw-1234567890"
    nonprod = "vgw-1234567890"
    prod    = "vgw-0987654321"
  }
}

# variable "vpc_cidr_secondary_a" {
#   description = "Core Infrastrucre VPC CIDR Secondary Zone A"
#   type = map(string)
#   default = {
#     default = "11.16.0.0/16"
#     lab     = "11.16.0.0/16"
#     staging = "11.32.0.0/16"
#     nonprod = "11.32.0.0/16"
#     prod    = "11.48.0.0/16"
#   }
# }

# variable "vpc_cidr_secondary_b" {
#   description = "Core Infrastrucre VPC CIDR Secondary Zone B"
#   type = map(string)
#   default = {
#     default = "12.16.0.0/16"
#     lab     = "12.16.0.0/16"
#     staging = "12.32.0.0/16"
#     nonprod = "12.32.0.0/16"
#     prod    = "12.48.0.0/16"
#   }
# }

# --------------------------------------------------------------------------
#  Infra Prefix
# --------------------------------------------------------------------------
# EC2 Prefix
variable "ec2_prefix" {
  description = "EC2 Prefix Name"
  type        = string
  default     = "ec2"
}

# EKS Prefix
variable "eks_prefix" {
  description = "EKS Prefix Name"
  type        = string
  default     = "eks"
}

# --------------------------------------------------------------------------
#  NAT Gateway
# --------------------------------------------------------------------------
# NAT EC2 Prefix
variable "nat_ec2_prefix" {
  description = "NAT EC2 Prefix Name"
  type        = string
  default     = "natgw_ec2"
}

# NAT EKS Prefix
variable "nat_eks_prefix" {
  description = "NAT EKS Prefix Name"
  type        = string
  default     = "natgw_eks"
}

# --------------------------------------------------------------------------
#  Subnet
# --------------------------------------------------------------------------
## EC2 Private
variable "ec2_private_a" {
  description = "Private Subnet for EC2 Zone A"
  type        = map(string)
  default = {
    default = "10.16.16.0/21"
    lab     = "10.16.16.0/21"
    staging = "10.32.16.0/21"
    nonprod = "10.32.16.0/21"
    prod    = "10.48.16.0/21"
  }
}

variable "ec2_private_b" {
  description = "Private Subnet for EC2 Zone B"
  type        = map(string)
  default = {
    default = "10.16.24.0/21"
    lab     = "10.16.24.0/21"
    staging = "10.32.24.0/21"
    nonprod = "10.32.24.0/21"
    prod    = "10.48.24.0/21"
  }
}

variable "ec2_private_c" {
  description = "Private Subnet for EC2 Zone C"
  type        = map(string)
  default = {
    default = "10.16.32.0/21"
    lab     = "10.16.32.0/21"
    staging = "10.32.32.0/21"
    nonprod = "10.32.32.0/21"
    prod    = "10.48.32.0/21"
  }
}

## EC2 Public
variable "ec2_public_a" {
  description = "Public Subnet for EC2 Zone A"
  type        = map(string)
  default = {
    default = "10.16.40.0/21"
    lab     = "10.16.40.0/21"
    staging = "10.32.40.0/21"
    nonprod = "10.32.40.0/21"
    prod    = "10.48.40.0/21"
  }
}

variable "ec2_public_b" {
  description = "Public Subnet for EC2 Zone B"
  type        = map(string)
  default = {
    default = "10.16.48.0/21"
    lab     = "10.16.48.0/21"
    staging = "10.32.48.0/21"
    nonprod = "10.32.48.0/21"
    prod    = "10.48.48.0/21"
  }
}

variable "ec2_public_c" {
  description = "Public Subnet for EC2 Zone C"
  type        = map(string)
  default = {
    default = "10.16.56.0/21"
    lab     = "10.16.56.0/21"
    staging = "10.32.56.0/21"
    nonprod = "10.32.56.0/21"
    prod    = "10.48.56.0/21"
  }
}

## EKS Private
variable "eks_private_a" {
  description = "Private Subnet for EKS Zone A"
  type        = map(string)
  default = {
    default = "10.16.64.0/21"
    lab     = "10.16.64.0/21"
    staging = "10.32.64.0/21"
    nonprod = "10.32.64.0/21"
    prod    = "10.48.64.0/21"
  }
}

variable "eks_private_b" {
  description = "Private Subnet for EKS Zone B"
  type        = map(string)
  default = {
    default = "10.16.72.0/21"
    lab     = "10.16.72.0/21"
    staging = "10.32.72.0/21"
    nonprod = "10.32.72.0/21"
    prod    = "10.48.72.0/21"
  }
}

variable "eks_private_c" {
  description = "Private Subnet for EKS Zone C"
  type        = map(string)
  default = {
    default = "10.16.80.0/21"
    lab     = "10.16.80.0/21"
    staging = "10.32.80.0/21"
    nonprod = "10.32.80.0/21"
    prod    = "10.48.80.0/21"
  }
}

## EKS Public
variable "eks_public_a" {
  description = "Public Subnet for EKS Zone A"
  type        = map(string)
  default = {
    default = "10.16.88.0/21"
    lab     = "10.16.88.0/21"
    staging = "10.32.88.0/21"
    nonprod = "10.32.88.0/21"
    prod    = "10.48.88.0/21"
  }
}

variable "eks_public_b" {
  description = "Public Subnet for EKS Zone B"
  type        = map(string)
  default = {
    default = "10.16.96.0/21"
    lab     = "10.16.96.0/21"
    staging = "10.32.96.0/21"
    nonprod = "10.32.96.0/21"
    prod    = "10.48.96.0/21"
  }
}

variable "eks_public_c" {
  description = "Public Subnet for EKS Zone C"
  type        = map(string)
  default = {
    default = "10.16.104.0/21"
    lab     = "10.16.104.0/21"
    staging = "10.32.104.0/21"
    nonprod = "10.32.104.0/21"
    prod    = "10.48.104.0/21"
  }
}

# --------------------------------------------------------------------------
#  Routing Table
# --------------------------------------------------------------------------
# EC2 RT Prefix
variable "ec2_rt_prefix" {
  description = "NAT EC2 Routing Table Prefix Name"
  type        = string
  default     = "ec2-rt"
}

# EKS RT Prefix
variable "eks_rt_prefix" {
  description = "NAT EKS Routing Table Prefix Name"
  type        = string
  default     = "eks-rt"
}

# --------------------------------------------------------------------------
#  Internet Gateway
# --------------------------------------------------------------------------
# IGW Prefix
variable "igw_prefix" {
  description = "IGW Prefix Name"
  type        = string
  default     = "igw"
}

# IGW RT Prefix
variable "igw_rt_prefix" {
  description = "IGW Routing Table Prefix Name"
  type        = string
  default     = "igw-rt"
}

# --------------------------------------------------------------------------
#  NAT Gateway
# --------------------------------------------------------------------------
# NAT Prefix
variable "nat_prefix" {
  description = "NAT Prefix Name"
  type        = string
  default     = "nat"
}

# NAT RT Prefix
variable "nat_rt_prefix" {
  description = "NAT Routing Table Prefix Name"
  type        = string
  default     = "nat-rt"
}