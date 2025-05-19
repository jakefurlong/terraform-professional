output "vpc_id" {
  value = aws_vpc.custom_vpc.id
}

output "internet_gateway_id" {
  value = aws_internet_gateway.custom_igw.id
}

output "route_table_id" {
  value = aws_route_table.public_rt.id
}

output "public_subnet_ids" {
  value = [for s in aws_subnet.public : s.id]
}

output "private_subnet_ids" {
  value = [for s in aws_subnet.private : s.id]
}

output "route_table_association_ids" {
  value = [for rta in aws_route_table_association.public : rta.id]
}

output "security_group_id" {
  value = aws_security_group.default_sg.id
}