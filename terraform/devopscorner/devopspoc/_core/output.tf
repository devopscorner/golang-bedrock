# ==========================================================================
#  Core: output.tf
# --------------------------------------------------------------------------
#  Description
#    Output Terraform Value
# --------------------------------------------------------------------------
#    - VPC ID
#    - VPC CIDR
#    - VPC Name
#    - VPC Security Group ID
#    - Subnet ID EC2 Private A
#    - Subnet ID EC2 Private B
#    - Subnet ID EC2 Private C
#    - Subnet CIDR Block EC2 Private A
#    - Subnet CIDR Block EC2 Private B
#    - Subnet CIDR Block EC2 Private C
#    - Subnet ID EC2 Public A
#    - Subnet ID EC2 Public B
#    - Subnet ID EC2 Public C
#    - Subnet CIDR Block EC2 Public A
#    - Subnet CIDR Block EC2 Public B
#    - Subnet CIDR Block EC2 Public C
#    - Subnet ID EKS Private A
#    - Subnet ID EKS Private B
#    - Subnet ID EKS Private C
#    - Subnet CIDR Block EKS Private A
#    - Subnet CIDR Block EKS Private B
#    - Subnet CIDR Block EKS Private C
#    - Subnet ID EKS Public A
#    - Subnet ID EKS Public B
#    - Subnet ID EKS Public C
#    - Subnet CIDR Block EKS Public B
#    - Subnet CIDR Block EKS Public A
#    - Subnet CIDR Block EKS Public C
# ==========================================================================

# --------------------------------------------------------------------------
#  VPC Output
# --------------------------------------------------------------------------
output "vpc_id" {
  description = "VPC Identity"
  value       = aws_vpc.infra_vpc.id
}
output "vpc_cidr" {
  description = "VPC CIDR Block"
  value       = aws_vpc.infra_vpc.cidr_block
}
output "vpc_name" {
  description = "VPC Name"
  value       = local.vps_tags.Name
}
output "security_group_id" {
  description = "Security Group of VPC Id's"
  value       = aws_security_group.default.id
}

# --------------------------------------------------------------------------
#  EC2 Output
# --------------------------------------------------------------------------
# EC2 Private
output "ec2_private_1a" {
  description = "Private Subnet EC2 Zone A"
  value       = aws_subnet.ec2_private_a.*.id
}
output "ec2_private_1a_cidr" {
  description = "Private Subnet EC2 CIDR Block of Zone A"
  value       = aws_subnet.ec2_private_a.cidr_block
}
output "ec2_private_1b" {
  description = "Private Subnet EC2 Zone B"
  value       = aws_subnet.ec2_private_b.*.id
}
output "ec2_private_1b_cidr" {
  description = "Private Subnet EC2 CIDR Block of Zone B"
  value       = aws_subnet.ec2_private_b.cidr_block
}
output "ec2_private_1c" {
  description = "Private Subnet EC2 Zone C"
  value       = aws_subnet.ec2_private_c.*.id
}
output "ec2_private_1c_cidr" {
  description = "Private Subnet EC2 CIDR Block of Zone C"
  value       = aws_subnet.ec2_private_c.cidr_block
}

# EC2 Public
output "ec2_public_1a" {
  description = "Public Subnet EC2 Zone A"
  value       = aws_subnet.ec2_public_a.*.id
}
output "ec2_public_1a_cidr" {
  description = "Public Subnet EC2 CIDR Block of Zone A"
  value       = aws_subnet.ec2_public_a.cidr_block
}
output "ec2_public_1b" {
  description = "Public Subnet EC2 Zone B"
  value       = aws_subnet.ec2_public_b.*.id
}
output "ec2_public_1b_cidr" {
  description = "Public Subnet EC2 CIDR Block of Zone B"
  value       = aws_subnet.ec2_public_b.cidr_block
}
output "ec2_public_1c" {
  description = "Public Subnet EC2 Zone C"
  value       = aws_subnet.ec2_public_c.*.id
}
output "ec2_public_1c_cidr" {
  description = "Public Subnet EC2 CIDR Block of Zone C"
  value       = aws_subnet.ec2_public_c.cidr_block
}

# --------------------------------------------------------------------------
#  EKS Output
# --------------------------------------------------------------------------
# EKS Private
output "eks_private_1a" {
  description = "Private Subnet EKS Zone A"
  value       = aws_subnet.eks_private_a.*.id
}
output "eks_private_1a_cidr" {
  description = "Private Subnet EKS CIDR Block of Zone A"
  value       = aws_subnet.eks_private_a.cidr_block
}
output "eks_private_1b" {
  description = "Private Subnet EKS Zone B"
  value       = aws_subnet.eks_private_b.*.id
}
output "eks_private_1b_cidr" {
  description = "Private Subnet EKS CIDR Block of Zone B"
  value       = aws_subnet.eks_private_b.cidr_block
}
output "eks_private_1c" {
  description = "Private Subnet EKS Zone C"
  value       = aws_subnet.eks_private_c.*.id
}
output "eks_private_1c_cidr" {
  description = "Private Subnet EKS CIDR Block of Zone C"
  value       = aws_subnet.eks_private_c.cidr_block
}

# EKS Public
output "eks_public_1a" {
  description = "Public Subnet EKS Zone A"
  value       = aws_subnet.eks_public_a.*.id
}
output "eks_public_1a_cidr" {
  description = "Public Subnet EKS CIDR Block of Zone A"
  value       = aws_subnet.eks_public_a.cidr_block
}
output "eks_public_1b" {
  description = "Public Subnet EKS Zone B"
  value       = aws_subnet.eks_public_b.*.id
}
output "eks_public_1b_cidr" {
  description = "Public Subnet EKS CIDR Block of Zone B"
  value       = aws_subnet.eks_public_b.cidr_block
}
output "eks_public_1c" {
  description = "Public Subnet EKS Zone C"
  value       = aws_subnet.eks_public_c.*.id
}
output "eks_public_1c_cidr" {
  description = "Public Subnet EKS CIDR Block of Zone C"
  value       = aws_subnet.eks_public_c.cidr_block
}

locals {
  summary = <<SUMMARY
VPC Summary:
  VPC Id:            ${aws_vpc.infra_vpc.id}
  Security Group Id: ${aws_security_group.default.id}
Subnet Private:
  EC2 Private 1a:    ${aws_subnet.ec2_private_a.id}
  EC2 Private 1b:    ${aws_subnet.ec2_private_b.id}
  EC2 Private 1c:    ${aws_subnet.ec2_private_c.id}
  EKS Private 1a:    ${aws_subnet.eks_private_a.id}
  EKS Private 1b:    ${aws_subnet.eks_private_b.id}
  EKS Private 1c:    ${aws_subnet.eks_private_c.id}
Subnet Public:
  EC2 Public 1a:     ${aws_subnet.ec2_public_a.id}
  EC2 Public 1b:     ${aws_subnet.ec2_public_b.id}
  EC2 Public 1c:     ${aws_subnet.ec2_public_c.id}
  EKS Public 1a:     ${aws_subnet.eks_public_a.id}
  EKS Public 1b:     ${aws_subnet.eks_public_b.id}
  EKS Public 1c:     ${aws_subnet.eks_public_c.id}
CIDR Block Private:
  EC2 CIDR 1a:       ${aws_subnet.ec2_private_a.cidr_block}
  EC2 CIDR 1b:       ${aws_subnet.ec2_private_b.cidr_block}
  EC2 CIDR 1c:       ${aws_subnet.ec2_private_c.cidr_block}
  EKS CIDR 1a:       ${aws_subnet.eks_private_a.cidr_block}
  EKS CIDR 1b:       ${aws_subnet.eks_private_b.cidr_block}
  EKS CIDR 1c:       ${aws_subnet.eks_private_c.cidr_block}
CIDR Block Public:
  EC2 CIDR 1a:       ${aws_subnet.ec2_public_a.cidr_block}
  EC2 CIDR 1b:       ${aws_subnet.ec2_public_b.cidr_block}
  EC2 CIDR 1c:       ${aws_subnet.ec2_public_c.cidr_block}
  EKS CIDR 1a:       ${aws_subnet.eks_public_a.cidr_block}
  EKS CIDR 1b:       ${aws_subnet.eks_public_b.cidr_block}
  EKS CIDR 1c:       ${aws_subnet.eks_public_c.cidr_block}
SUMMARY
}

output "summary" {
  description = "Summary Core Infrastructure Configuration"
  value       = local.summary
}
