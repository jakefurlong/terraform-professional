module "asg" {
  source = "../../../modules/asg"

  asg_name                = "nimbusdevops-test"
  server_port             = 8080
  asg_sg_ingress          = ["0.0.0.0/0"]
  machine_image           = "ami-04f7a54071e74f488"
  instance_type           = "t2.micro"
  max_size                = 2
  min_size                = 2
  enable_autoscaling      = true
  aws_vpc_zone_identifier = ["subnet-08902ab9b0c0b2bad", "subnet-002d54f12475a6066"]
}