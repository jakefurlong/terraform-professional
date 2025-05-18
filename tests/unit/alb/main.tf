module "alb" {
  source = "../../../modules/alb"

  stack_name = "nimbusdevops-test"
  vpc_id = null
  subnet_ids = []
  server_port = 8080
  alb_sg_ingress_cidr_range = ["0.0.0.0/0"]

}