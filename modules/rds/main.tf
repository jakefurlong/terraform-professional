resource "aws_db_instance" "rds_database" {
  identifier_prefix   = var.db_identifier_prefix
  engine              = var.db_engine
  allocated_storage   = var.db_allocated_storage
  instance_class      = var.db_instance_class
  skip_final_snapshot = var.db_skip_final_snapshot
  db_name             = var.database_name

  username = var.db_username
  password = var.db_password
}

