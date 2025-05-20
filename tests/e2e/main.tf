module "vpc" {
  source = "../../modules/vpc"

  aws_vpc_cidr_block = "172.0.0.0/16"
  aws_network_name   = "nimbusdevops-test"

}

module "rds" {
  source = "../../modules/rds"

  db_identifier_prefix   = "e2e-db"
  db_engine              = "mysql"
  db_instance_class      = "db.t3.micro"
  db_allocated_storage   = 20
  db_username            = local.db_credentials["username"]
  db_password            = local.db_credentials["password"]
  database_name          = "e2edb"
  db_skip_final_snapshot = true
  rds_sg_ingress = ["172.0.0.0/16"]

  vpc_id     = module.vpc.vpc_id
  subnet_ids = module.vpc.private_subnet_ids
}

#
#module "alb" {
#  source                    = "../../modules/alb"
#  stack_name                = "nimbusdevops-test"
#  vpc_id                    = null
#  subnet_ids                = []
#  server_port               = 8080
#  alb_sg_ingress_cidr_range = ["0.0.0.0/0"]
#}
#
#module "asg" {
#  source = "../../modules/asg"
#
#  asg_name             = "nimbusdevops-test"
#  server_port          = 8080
#  asg_sg_ingress       = ["0.0.0.0/0"]
#  machine_image        = "ami-04f7a54071e74f488"
#  instance_type        = "t2.micro"
#  max_size             = 2
#  min_size             = 2
#  enable_autoscaling   = true
#  aws_target_group_arn = module.alb.aws_target_group_arn
#}
#
