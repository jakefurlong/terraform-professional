output "vpc_id" {
  value = module.vpc.vpc_id
}

output "internet_gateway_id" {
  value = module.vpc.internet_gateway_id
}

output "route_table_id" {
  value = module.vpc.route_table_id
}

output "public_subnet_ids" {
  value = module.vpc.public_subnet_ids
}

output "route_table_association_ids" {
  value = module.vpc.route_table_association_ids
}

output "security_group_id" {
  value = module.vpc.security_group_id
}
