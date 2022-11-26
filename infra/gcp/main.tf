provider "google" {
  # Project ID
  project     = var.project_id
  credentials = file(var.gcp_credentials_file_path)
  zone        = "asia-northeast1"
}

# Profect prefix
variable "project_prefix" {
  description = "GCP profect prefix"
  type        = string
}

# Profect env
variable "project_env" {
  description = "GCP profect env"
  type        = string
}

# GCP credentials File Path
variable "gcp_credentials_file_path" {
  description = "File Path to GCP Credentials"
  type        = string
}

# Project ID
variable "project_id" {
  description = "GCP project id"
  type        = string
}

# APIキーのシークレット
variable "secret_core_api_key" {
  description = "Secret: API Key"
  type        = string
}

# データベースURLのシークレット
variable "secret_core_database_url" {
  description = "Secret: Database URL"
  type        = string
}
