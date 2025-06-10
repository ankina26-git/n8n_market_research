cd ~/myproject/n8n_project

# .env ファイルの内容を環境変数として読み込む
# (direnv などを利用すると自動化できます)
source .env

# 環境変数に設定された値が TF_VAR_ プレフィックスで始まる変数に自動的にマッピングされます
# 例: TF_VAR_GCP_PROJECT_ID -> var.gcp_project_id
# TF_VAR_DB_USER -> var.db_user

# Terraform の初期化
terraform -chdir=terraform init

# Terraform プランの確認
terraform -chdir=terraform plan

# Terraform の適用
terraform -chdir=terraform apply