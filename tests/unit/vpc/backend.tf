terraform {
  backend "s3" {
    bucket         = "your-terraform-state-bucket"
    key            = "terraform-state/tests/unit/vpc/terraform.tfstate"
    region         = "us-west-1"
    encrypt        = true
    use_lockfile   = true
  }
}
