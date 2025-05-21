module "vpc" {
  source = "../../../modules/vpc"

  aws_vpc_cidr_block = "10.0.0.0/16"
  aws_network_name   = "nimbusdevops-test"
}