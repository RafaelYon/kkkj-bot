resource "aws_iam_role" "ec2" {
  description           = var.app_name
  assume_role_policy    = file("${path.module}/iam/EC2TrustPolicy.json")
  force_detach_policies = true
}

resource "aws_iam_role_policy" "ec2" {
  role = aws_iam_role.ec2.id

  policy = templatefile("${path.module}/iam/EC2Policy.json", {
    aws_region  = var.aws_region
    aws_account = data.aws_caller_identity.current.account_id
    app_name    = var.app_name
  })
}

resource "aws_iam_instance_profile" "ec2" {
  name_prefix = var.app_name
  role        = aws_iam_role.ec2.name
}

data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"] # Canonical
}

resource "aws_instance" "web" {
  ami                         = data.aws_ami.ubuntu.id
  instance_type               = var.ec2_type
  iam_instance_profile        = aws_iam_instance_profile.ec2.name
  key_name                    = var.ec2_key_pair
  vpc_security_group_ids      = [data.aws_security_group.default.id]
  monitoring                  = false
  associate_public_ip_address = true

  root_block_device {
    delete_on_termination = true
    volume_size           = 10
    volume_type           = "gp2"
  }
}