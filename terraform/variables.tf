variable "project" {
  description = "プロジェクト名（リソース名やタグに利用）"
  type        = string
  default     = "comic-rss"
}

variable "env" {
  description = "環境名 (dev / stg / prod)"
  type        = string
  default     = "dev"
}

variable "aws_region" {
  description = "デプロイ先のAWSリージョン"
  type        = string
  default     = "ap-northeast-1"
}
