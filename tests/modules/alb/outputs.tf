output "alb_dns_name" {
  description = "The domain name of the load balancer"
  value       = module.alb.alb_dns_name
}

output "aws_target_group_arn" {
  description = "ARN of the Target Group"
  value = null
}