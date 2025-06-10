# terraform/modules/cloudrun_go/main.tf

resource "google_cloud_run_v2_service" "service" {
  project  = var.project_id
  name     = var.service_name
  location = var.region

  template {
    containers {
      image = var.container_image
      env {
        name  = "DATABASE_URL"
        value = "postgres://${var.db_user}:${var.db_password}@/${var.db_name}?host=/cloudsql/${var.db_connection_name}" # Cloud SQL 接続文字列
      }
      # その他の環境変数
    }
    scaling {
      min_instance_count = 0
      max_instance_count = 1 # 開発用のため最小限
    }
    # Cloud SQL への接続設定
    volumes {
      name = "cloudsql"
      cloud_sql_instance {
        instances = [var.db_connection_name]
      }
    }
  }

  traffic {
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }
}

resource "google_cloud_run_service_iam_member" "noauth_access" {
  project  = var.project_id
  location = var.region
  service  = google_cloud_run_v2_service.service.name
  role     = "roles/run.invoker" # 外部からのアクセスを許可
  member   = "allUsers" # 誰でもアクセスできるように設定 (テスト用、本番環境では認証を設定すべき)
}

variable "project_id" {
  description = "GCP プロジェクト ID"
  type        = string
}

variable "region" {
  description = "GCP リージョン"
  type        = string
}

variable "service_name" {
  description = "Cloud Run サービス名"
  type        = string
}

variable "container_image" {
  description = "デプロイするコンテナイメージのパス"
  type        = string
}

variable "db_user" {
  description = "Cloud SQL ユーザー名"
  type        = string
  sensitive   = true
}

variable "db_password" {
  description = "Cloud SQL パスワード"
  type        = string
  sensitive   = true
}

variable "db_name" {
  description = "Cloud SQL データベース名"
  type        = string
}

variable "db_connection_name" {
  description = "Cloud SQL インスタンスの接続名"
  type        = string
}

output "service_url" {
  description = "デプロイされた Cloud Run サービスの URL"
  value       = google_cloud_run_v2_service.service.uri
}