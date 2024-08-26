# ==========================================================================
#  Core: natgw.tf
# --------------------------------------------------------------------------
#  Description
#    NAT Gateway for EC2 & EKS
# --------------------------------------------------------------------------
#    - NAT Gateway from Private Subnet
#    - EIP Enabled
#    - Route Table for Private Subnet with NAT Gateway
# ==========================================================================

# --------------------------------------------------------------------------
#  NAT GW Tags
# --------------------------------------------------------------------------
locals {
  ## EKS
  tags_nat_eks_rt_private = {
    ResourceGroup = "${var.environment[local.env]}-RT-EKS"
  }
  tags_nat_eks = {
    ResourceGroup = "${var.environment[local.env]}-NAT-EKS"
  }
}

# --------------------------------------------------------------------------
#  EIP (enabled)
# --------------------------------------------------------------------------
## EKS
resource "aws_eip" "eks" {
  tags = {
    "Name" = "${var.coreinfra}-${var.workspace_env[local.env]}-eip-eks"
  }

  tags_all = {
    "Name" = "${var.coreinfra}-${var.workspace_env[local.env]}-eip-eks"
  }
}

# --------------------------------------------------------------------------
#  NAT GW
# --------------------------------------------------------------------------
resource "aws_nat_gateway" "eks_ngw" {
  allocation_id = aws_eip.eks.id
  subnet_id     = aws_subnet.eks_public_a.id

  tags = merge(local.tags, local.tags_nat_eks, { Name = "${var.coreinfra}-${var.workspace_env[local.env]}-${var.nat_eks_prefix}" })

  lifecycle {
    ignore_changes = [
      allocation_id
    ]
  }
}

## --------------------------------------------------------------------------
#  Route Table NAT GW
# --------------------------------------------------------------------------
## EKS
resource "aws_route_table" "nat_eks_rt_private_a" {
  vpc_id = aws_vpc.infra_vpc.id
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.eks_ngw.id
  }

  # propagating_vgws = [var.propagating_vgws[local.env]]
  # route{
  #   cidr_block                = var.cidr_block_vpc_peering[local.env]
  #   vpc_peering_connection_id = aws_vpc_peering_connection.vpc_peer.id
  # }

  tags = merge(local.tags, local.tags_nat_eks_rt_private, { Name = "${var.coreinfra}-${var.workspace_env[local.env]}-${var.eks_rt_prefix}-private-${var.aws_region}a" }, local.tags_internal_elb)
}

resource "aws_route_table" "nat_eks_rt_private_b" {
  vpc_id = aws_vpc.infra_vpc.id
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.eks_ngw.id
  }

  # propagating_vgws = [var.propagating_vgws[local.env]]
  # route{
  #   cidr_block                = var.cidr_block_vpc_peering[local.env]
  #   vpc_peering_connection_id = aws_vpc_peering_connection.vpc_peer.id
  # }

  tags = merge(local.tags, local.tags_nat_eks_rt_private, { Name = "${var.coreinfra}-${var.workspace_env[local.env]}-${var.eks_rt_prefix}-private-${var.aws_region}b" }, local.tags_internal_elb)
}

resource "aws_route_table" "nat_eks_rt_private_c" {
  vpc_id = aws_vpc.infra_vpc.id
  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.eks_ngw.id
  }

  # propagating_vgws = [var.propagating_vgws[local.env]]
  # route{
  #   cidr_block                = var.cidr_block_vpc_peering[local.env]
  #   vpc_peering_connection_id = aws_vpc_peering_connection.vpc_peer.id
  # }

  tags = merge(local.tags, local.tags_nat_eks_rt_private, { Name = "${var.coreinfra}-${var.workspace_env[local.env]}-${var.eks_rt_prefix}-private-${var.aws_region}c" }, local.tags_internal_elb)
}

# --------------------------------------------------------------------------
#  Route Table with Private Subnet
# --------------------------------------------------------------------------
## EKS
resource "aws_route_table_association" "nat_eks_rt_private_1a" {
  subnet_id      = aws_subnet.eks_private_a.id
  route_table_id = aws_route_table.nat_eks_rt_private_a.id
}

resource "aws_route_table_association" "nat_eks_rt_private_1b" {
  subnet_id      = aws_subnet.eks_private_b.id
  route_table_id = aws_route_table.nat_eks_rt_private_b.id
}

resource "aws_route_table_association" "nat_eks_rt_private_1c" {
  subnet_id      = aws_subnet.eks_private_c.id
  route_table_id = aws_route_table.nat_eks_rt_private_c.id
}
