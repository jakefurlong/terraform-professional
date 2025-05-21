output "vpc_name" {
  value = module.vpc.vpc_name
}

output "internet_gateway_name" {
  value = module.vpc.internet_gateway_name
}

output "public_subnet_ids" {
  value = module.vpc.public_subnet_ids
}

output "private_subnet_ids" {
  value = module.vpc.private_subnet_ids
}

output "route_table_association_ids" {
  value = module.vpc.route_table_association_ids
}

output "security_group_name" {
  value = module.vpc.security_group_name
}

output "aws_route_table_public_name" {
  value = module.vpc.aws_route_table_public_name
}
