terraform {
  required_providers {
    cockroach = {
      source  = "cockroachdb/cockroach"
      version = "0.3.2"
    }
  }
}

variable "project_prefix" {
  description = "profect prefix"
  type        = string
}

variable "project_env" {
  description = "profect env"
  type        = string
}

variable "sql_user_password" {
  type      = string
  nullable  = false
  sensitive = true
}

variable "serverless_spend_limit" {
  type     = number
  nullable = false
  default  = 0
}

variable "cloud_provider" {
  type     = string
  nullable = false
  default  = "GCP"
}

variable "cloud_provider_regions" {
  type     = list(string)
  nullable = false
  default  = ["asia-southeast1"]
}

provider "cockroach" {
  # export COCKROACH_API_KEY with the cockroach cloud API Key
}

resource "cockroach_cluster" "core_db" {
  name           = "${var.project_prefix}-${var.project_env}-core-db"
  cloud_provider = var.cloud_provider
  serverless = {
    spend_limit = var.serverless_spend_limit
  }
  regions = [for r in var.cloud_provider_regions : { name = r }]
}

resource "cockroach_sql_user" "core_db_user" {
  name       = "${var.project_prefix}-${var.project_env}-core-db"
  password   = var.sql_user_password
  cluster_id = cockroach_cluster.core_db.id
}
