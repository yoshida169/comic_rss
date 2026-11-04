terraform {
  required_version = ">= 1.9"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
  }

  # 当面はローカルstate。チーム開発/CI導入時にS3バックエンドへ移行する
  # backend "s3" {
  #   bucket       = "<state用バケット名>"
  #   key          = "comic-rss/terraform.tfstate"
  #   region       = "ap-northeast-1"
  #   use_lockfile = true
  # }
}
