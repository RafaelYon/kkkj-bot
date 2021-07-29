variable "aws_region" {
  type        = string
  description = "The AWS region of resources"
  default     = "us-east-2"
}

variable "app_name" {
  type        = string
  description = "The Application name"
  default     = "kkkj-bot"
}

variable "ec2_security_group_id" {
  type        = string
  description = "The EC2 security group ID"
}

variable "ec2_key_pair" {
  type        = string
  description = "The EC2 key pair name"
}

variable "ec2_type" {
  type        = string
  description = "The EC2 instance type"
  default     = "t2.micro"
}
