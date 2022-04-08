output "vpc_id" {
  description = "VPC internal UUID"
  value       = nws_ec2_vpc.vpc.id
}
