module "alb" {
  source = "../../../modules/alb"

  security_group_name = "nimbusdevops-test-sg"
  target_group_name   = "nimbusdevops-test-tg"
  load_balancer_name  = "nimbusdevops-test-alb"
  vpc_id = null
  subnet_ids = []
  server_port = 8080
  alb_sg_ingress_cidr_range = ["0.0.0.0/0"]

}