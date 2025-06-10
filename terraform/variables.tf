variable "gcp_project_id" {
  description = "GCP プロジェクト ID"
  type        = string
}

variable "gcp_region" {
  description = "GCP リージョン"
  type        = string
  default     = "asia-northeast1" # 東京リージョン
}

variable "service_prefix" {
  description = "サービス名のプレフィックス (例: autocoredx)"
  type        = string
}

variable "db_user" {
  description = "Cloud SQL データベースユーザー名"
  type        = string
  sensitive   = true # 機密情報として扱う
}

variable "db_password" {
  description = "Cloud SQL データベースパスワード"
  type        = string
  sensitive   = true # 機密情報として扱う
}