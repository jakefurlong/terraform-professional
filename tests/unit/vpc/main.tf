module "vpc" {
  source = "git::https://github.com/jakefurlong/terraform-modules.git//vpc?ref=v1.0.1"

  aws_vpc_cidr_block = "10.0.0.0/16"
  aws_network_name   = "nimbusdevops-test"
}