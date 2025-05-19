terraform {
  backend "s3" {
    bucket       = "nimbusdevops"
    key          = "terraform-state/tests/unit/alb/terraform.tfstate"
    region       = "us-west-1"
    encrypt      = true
    use_lockfile = true
  }
}
