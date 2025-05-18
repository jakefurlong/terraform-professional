terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "6.0.0-beta1"
    }
  }
  required_version = "~> 1.11.4"
}

provider "aws" {
  region = "us-west-1"
}