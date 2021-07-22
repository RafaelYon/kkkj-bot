terraform {
  required_version = "~>1.0.1"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=3.50.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
  default_tags {
    tags = {
      App = var.app_name
    }
  }
}
