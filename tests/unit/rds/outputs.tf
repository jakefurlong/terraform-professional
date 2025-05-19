output "database_address" {
  value       = module.rds.database_address
  description = "Connect to the database at this endpoint"
}

output "database_port" {
  value       = module.rds.database_port
  description = "The port the database is listening on"
}