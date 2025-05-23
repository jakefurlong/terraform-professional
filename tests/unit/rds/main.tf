module "rds" {
  source = "git::https://github.com/jakefurlong/terraform-modules.git//rds?ref=v1.0.1"

  db_username            = local.db_credentials["username"]
  db_password            = local.db_credentials["password"]
  db_identifier_prefix   = "nimbusdevops-test"
  db_engine              = "mysql"
  db_allocated_storage   = 10
  db_instance_class      = "db.t3.micro"
  db_skip_final_snapshot = true
  database_name          = "mysqltest"
  rds_sg_ingress         = ["172.16.0.0/16"]
  vpc_id                 = null
  subnet_ids             = null
}