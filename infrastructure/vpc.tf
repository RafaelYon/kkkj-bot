data "aws_security_group" "default" {
  id = var.ec2_security_group_id
}