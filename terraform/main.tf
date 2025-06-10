# GCP プロバイダの設定
# Terraform が GCP のリソースを管理するために必要です
provider "google" {
  project = var.gcp_project_id
  region  = var.gcp_region
}


# ネットワークモジュールの呼び出し
# VPC (Virtual Private Cloud) ネットワークとサブネットを作成します
module "network" {
  source        = "./modules/network" # network モジュールのパス
  project_id    = var.gcp_project_id
  network_name  = "${var.service_prefix}-vpc-network"
  subnet_name   = "${var.service_prefix}-subnet"
  subnet_ip_cidr_range = "10.0.0.0/20" # 任意のIPレンジを設定
  subnet_region = var.gcp_region
}

# Cloud SQL モジュールの呼び出し
# PostgreSQL データベースインスタンスを作成します
module "cloudsql" {
  source           = "./modules/cloudsql" # cloudsql モジュールのパス
  project_id       = var.gcp_project_id
  region           = var.gcp_region
  instance_name    = "${var.service_prefix}-cloudsql-instance"
  database_name    = "${var.service_prefix}-database"
  database_user    = var.db_user
  database_password = var.db_password
}

# Cloud Run (Go) モジュールの呼び出し
# Go アプリケーション用の Cloud Run サービスを作成します
module "cloudrun_go" {
  source           = "./modules/cloudrun_go" # cloudrun_go モジュールのパス
  project_id       = var.gcp_project_id
  region           = var.gcp_region
  service_name     = "${var.service_prefix}-go-service"
  container_image  = "gcr.io/${var.gcp_project_id}/go-app:latest" # 後ほどビルドしたイメージのパスに更新します
  # Cloud SQL との接続設定 (必要な場合)
  # depends_on       = [module.cloudsql]
  # db_connection_name = module.cloudsql.connection_name # Cloud SQL からの出力を使用
}

# Cloud Run (Django) モジュールの呼び出し
# Django アプリケーション用の Cloud Run サービスを作成します
module "cloudrun_django" {
  source           = "./modules/cloudrun_django" # cloudrun_django モジュールのパス
  project_id       = var.gcp_project_id
  region           = var.gcp_region
  service_name     = "${var.service_prefix}-django-service"
  container_image  = "gcr.io/${var.gcp_project_id}/django-app:latest" # 後ほどビルドしたイメージのパスに更新します
  # Cloud SQL との接続設定
  depends_on       = [module.cloudsql] # Cloud SQL が先に作成されるように依存関係を設定
  db_connection_name = module.cloudsql.connection_name # Cloud SQL からの出力を使用
}

# Cloud Run (n8n) モジュールの呼び出し
# n8n アプリケーション用の Cloud Run サービスを作成します
module "cloudrun_n8n" {
  source           = "./modules/cloudrun_n8n" # cloudrun_n8n モジュールのパス
  project_id       = var.gcp_project_id
  region           = var.gcp_region
  service_name     = "${var.service_prefix}-n8n-service"
  container_image  = "gcr.io/${var.gcp_project_id}/n8n-app:latest" # 後ほどビルドしたイメージのパスに更新します
  # n8n が Cloud SQL を利用する場合の接続設定
  depends_on       = [module.cloudsql]
  db_connection_name = module.cloudsql.connection_name # Cloud SQL からの出力を使用
}