# terraform/modules/cloudsql/main.tf

resource "google_sql_database_instance" "main_instance" {
  project          = var.project_id
  name             = var.instance_name
  database_version = "POSTGRES_14" # 必要に応じてバージョンを変更
  region           = var.region
  settings {
    tier = "db-f1-micro" # 開発用のため最小限のティア。本番環境では適切なティアを選択
    ip_configuration {
      ipv4_enabled = true
      # private_network = module.network.id # プライベート IP 接続の場合
      # authorized_networks {
      #   value = "0.0.0.0/0" # 全てのIPからの接続を許可 (テスト用、本番環境では特定のIPに制限すべき)
      # }
    }
  }
}

resource "google_sql_database" "database" {
  project    = var.project_id
  name       = var.database_name
  instance   = google_sql_database_instance.main_instance.name
  charset    = "UTF8"
  collation  = "en_US.UTF8"
}

resource "google_sql_user" "user" {
  project  = var.project_id
  name     = var.database_user
  instance = google_sql_database_instance.main_instance.name
  host     = "%" # 任意のホストからの接続を許可 (テスト用、本番環境では特定のホストに制限すべき)
  password = var.database_password
}

variable "project_id" {
  description = "GCP プロジェクト ID"
  type        = string
}

variable "region" {
  description = "GCP リージョン"
  type        = string
}

variable "instance_name" {
  description = "Cloud SQL インスタンス名"
  type        = string
}

variable "database_name" {
  description = "データベース名"
  type        = string
}

variable "database_user" {
  description = "データベースユーザー名"
  type        = string
}

variable "database_password" {
  description = "データベースパスワード"
  type        = string
  sensitive   = true
}

output "connection_name" {
  description = "Cloud SQL インスタンスの接続名"
  value       = google_sql_database_instance.main_instance.connection_name
}