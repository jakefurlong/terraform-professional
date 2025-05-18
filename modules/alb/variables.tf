variable "vpc_id" {
  description = "VPC ID to use. Defaults to the default VPC if not set."
  type        = string
  default     = null
}

variable "security_group_name" {
  description = "Name of the Security Group"
  type        = string
}

variable "load_balancer_name" {
  description = "Name of the ALB"
  type        = string
}

variable "target_group_name" {
  description = "name of the Target Group"
  type        = string
}

variable "server_port" {
  description = "The port the server will use for HTTP requests"
  default     = 8080
  type        = number
}

variable "subnet_ids" {
  description = "List of subnet IDs to use. Defaults to subnets in default VPC if not set."
  type        = list(string)
  default     = []
}

variable "alb_sg_ingress_cidr_range" {
  description = "Allow CIDR range for ALB ingress"
  type        = list(string)
}