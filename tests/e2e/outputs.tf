output "vpc_id" {
  value = module.vpc.vpc_id
}

output "private_subnet_ids" {
  value = module.vpc.private_subnet_ids
}

output "public_subnet_ids" {
  value = module.vpc.public_subnet_ids
}

output "rds_endpoint" {
  value = module.rds.db_endpoint
}

output "db_identifier" {
  value = module.rds.db_identifier
}

output "alb_dns_name" {
  value = module.alb.alb_dns_name
}