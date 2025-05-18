output "database_address" {
    value = aws_db_instance.rds_database.address
    description = "Connect to the database at this endpoint"
}

output "database_port" {
    value = aws_db_instance.rds_database.port
    description = "The port the database is listening on"
}