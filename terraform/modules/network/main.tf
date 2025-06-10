# terraform/modules/network/main.tf

resource "google_compute_network" "vpc_network" {
  project                 = var.project_id
  name                    = var.network_name
  auto_create_subnetworks = false # サブネットは手動で作成
}

resource "google_compute_subnetwork" "subnet" {
  project       = var.project_id
  name          = var.subnet_name
  ip_cidr_range = var.subnet_ip_cidr_range
  region        = var.subnet_region
  network       = google_compute_network.vpc_network.id
}

variable "project_id" {
  description = "GCP プロジェクト ID"
  type        = string
}

variable "network_name" {
  description = "VPC ネットワーク名"
  type        = string
}

variable "subnet_name" {
  description = "サブネット名"
  type        = string
}

variable "subnet_ip_cidr_range" {
  description = "サブネットの IP CIDR 範囲"
  type        = string
}

variable "subnet_region" {
  description = "サブネットのリージョン"
  type        = string
}

output "network_id" {
  description = "作成された VPC ネットワークの ID"
  value       = google_compute_network.vpc_network.id
}

output "subnet_id" {
  description = "作成されたサブネットの ID"
  value       = google_compute_subnetwork.subnet.id
}