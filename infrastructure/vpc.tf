data "aws_security_group" "default" {
  id = var.ec2_security_group_id
}

resource "aws_eip" "ec2" {
  instance = aws_instance.web.id
  vpc      = true
}