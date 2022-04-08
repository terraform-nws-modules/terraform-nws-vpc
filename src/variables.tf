# ---------------------------------------------------------------------------------------------------------------------
# REQUIRED PARAMETERS
# You must provide a value for each of these parameters.
# ---------------------------------------------------------------------------------------------------------------------
variable "name" {
  description = "Your VPC name"
  type        = string
}

variable "cidr4" {
  description = "Your VPC top CIDR"
  type        = string
}

variable "domain" {
  description = "Your VPC network domain name"
  type        = string
}

# ---------------------------------------------------------------------------------------------------------------------
# OPTIONAL PARAMETERS
# These parameters have reasonable defaults.
# ---------------------------------------------------------------------------------------------------------------------
variable "zone" {
  description = "Your zone name"
  type        = string
  default     = "central-0"
}

variable "network_domain" {
  description = "Your VPC network domain name"
  type        = string
  default     = "my.local"
}
