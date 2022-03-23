### From March 2022

# We want to deploy a server on AWS in such a way that if someone were to edit the settings and mangle them
# beyond all recognition, we could quickly recover them. Write the Terraform required to deploy a simple EC2
# instance that is available to the internet on ports 22 and 443.

# Make sure you change the subnet ID, key name, and IAM instance profile. But after that, you can run this code:
#   terraform apply

# Assuming you have the private key on your computer at ~/.ssh/example_key_pair_private, you can SSH to the server:
#   ssh -i ~/.ssh/example_key_pair_private ubuntu@PUBLIC_IP_OF_SERVER

# With this terraform, the security group settings and the EC2 instance settings (although there aren't much here),
# would be reproducible easily, and could recreate the EC2 box as needed (or fix whatever was broken).

resource "aws_instance" "ec2_instance" {
  ami                                  = "ami-075c8e2e1712231db" # Ubuntu Linux AMI
  instance_type                        = "t4g.micro"
  key_name                             = "example-key-pair" # Example key pair name
  iam_instance_profile                 = "example-instance-profile" # Example IAM instance profile name
  vpc_security_group_ids               = [aws_security_group.security_group.id]
  subnet_id                            = "subnet-ExampleSubnetID12" # Example subnet ID
  associate_public_ip_address          = true
  instance_initiated_shutdown_behavior = "stop"

  root_block_device {
    volume_size = "16"
  }

  tags = {
    Name = "test-server"
  }

  lifecycle {
    ignore_changes = [key_name]
  }
}

resource "aws_security_group" "security_group" {
  name        = "test-server"
  description = "A security group that is open on ports 22 and 443"
  vpc_id      = var.vpc_id

  tags = {
    Name = "test-server"
  }
}

resource "aws_security_group_rule" "egress" {
  type              = "egress"
  from_port         = 0
  to_port           = 0
  protocol          = "-1"
  cidr_blocks       = ["0.0.0.0/0"]
  description       = "All egress communication"
  security_group_id = aws_security_group.security_group.id
}

resource "aws_security_group_rule" "ingress_22" {
  type              = "ingress"
  from_port         = 22
  to_port           = 22
  protocol          = "TCP"
  cidr_blocks       = ["0.0.0.0/0"]
  description       = "Port 22 open"
  security_group_id = aws_security_group.security_group.id
}

resource "aws_security_group_rule" "ingress_443" {
  type              = "ingress"
  from_port         = 443
  to_port           = 443
  protocol          = "TCP"
  cidr_blocks       = ["0.0.0.0/0"]
  description       = "Port 443 open"
  security_group_id = aws_security_group.security_group.id
}
