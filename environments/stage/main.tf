module "vpc" {
  source = "../../modules/vpc"

  aws_vpc_cidr_block = "10.0.0.0/16"
  aws_network_name   = "nimbusdevops-stage"
}

#module "alb" {
#  source = "../../modules/alb"
#
#  stack_name                = "nimbusdevops-stage"
#  vpc_id                    = module.vpc.vpc_id
#  subnet_ids                = module.vpc.public_subnet_ids
#  server_port               = 8080
#  alb_sg_ingress_cidr_range = ["0.0.0.0/0"]
#}
#
#module "asg" {
#  source = "../../modules/asg"
#
#  asg_name             = "nimbusdevops-stage"
#  server_port          = 8080
#  asg_sg_ingress       = ["0.0.0.0/0"]
#  machine_image        = "ami-04f7a54071e74f488"
#  instance_type        = "t3.micro"
#  max_size             = 4
#  min_size             = 2
#  aws_target_group_arn = [module.alb.aws_target_group_arn]
#  enable_autoscaling   = true
#
#  depends_on = [module.alb]
#}
#
#module "rds" {
#  source = "../../modules/rds"
#
#  db_username            = local.db_credentials["username"]
#  db_password            = local.db_credentials["password"]
#  db_identifier_prefix   = "nimbusdevops-stage"
#  db_engine              = "mysql"
#  db_allocated_storage   = 10
#  db_instance_class      = "db.t3.micro"
#  db_skip_final_snapshot = true
#  database_name          = "mysqlstage"
#}
