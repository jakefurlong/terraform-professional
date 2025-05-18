output "alb_dns_name" {
  value       = module.alb.alb_dns_name
  description = "The dns name of the load balancer"
}