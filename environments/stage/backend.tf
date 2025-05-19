terraform {
  backend "s3" {
    bucket  = "nimbusdevops"
    key     = "terraform-state/environments/stage/terraform.tfstate"
    region  = "us-west-1"
    encrypt = true
  }
}
