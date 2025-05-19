terraform {
  backend "s3" {
    bucket       = "nimbusdevops"
    key          = "terraform-state/tests/e2e/terraform.tfstate"
    region       = "us-west-1"
    encrypt      = true
    use_lockfile = true
  }
}
