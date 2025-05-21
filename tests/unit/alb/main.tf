module "alb" {
  source = "../../../modules/alb"

  stack_name                = "nimbusdevops-test"
  aws_vpc_id                = "vpc-837eafe5"
  subnet_ids                = ["subnet-40f07326", "subnet-ba6db1e0"]
  server_port               = 8080
  alb_sg_ingress_cidr_range = ["0.0.0.0/0"]

}