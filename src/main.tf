terraform {
  required_version = ">= 1.1.5"

  required_providers {
    nws = {
      source  = "nws/nwscc"
      version = "0.0.1"
    }
  }
}

data "nws_ec2_zone" "zone" {
  name = var.zone
}

data "nws_ec2_domain" "dom" {
  name = var.domain
}

data "nws_ec2_vpc_offer" "default" {
  name = "defaultVPC"
}

resource "nws_ec2_vpc" "vpc" {
  zone_id        = data.nws_ec2_zone.zone.id
  domain_id      = data.nws_ec2_domain.dom.id
  vpc_offer_id   = data.nws_ec2_vpc_offer.default.id
  name           = var.name
  cidr4          = var.cidr4
  network_domain = var.network_domain
}
