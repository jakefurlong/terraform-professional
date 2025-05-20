locals {
  selected_private_subnets = length(var.private_subnet_ids) > 0 ? var.private_subnet_ids : data.aws_subnets.default_private.ids
}