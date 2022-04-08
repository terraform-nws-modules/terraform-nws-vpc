terraform {
  required_version = ">= 1.1.5"

  required_providers {
    nws = {
      source  = "nws/nwscc"
      version = "0.0.1"
    }
  }
}

module "vpc" {
  source = "../../src"

  name           = var.name
  zone           = var.zone
  cidr4          = var.cidr4
  domain         = var.domain
  network_domain = var.network_domain
}
