module "alb" {
  source                    = "../../../modules/alb"
  stack_name                = "nimbusdevops-test"
  vpc_id                    = null
  subnet_ids                = []
  server_port               = 8080
  alb_sg_ingress_cidr_range = ["0.0.0.0/0"]
}

module "asg" {
  source = "../../../modules/asg"

  asg_name             = "nimbusdevops-test"
  server_port          = 8080
  asg_sg_ingress       = ["0.0.0.0/0"]
  machine_image        = "ami-04f7a54071e74f488"
  instance_type        = "t2.micro"
  max_size             = 2
  min_size             = 2
  enable_autoscaling   = true
  aws_target_group_arn = module.alb.aws_target_group_arn
}

output "alb_dns_name" {
  value = module.alb.alb_dns_name
}
