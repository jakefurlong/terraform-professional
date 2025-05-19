data "aws_secretsmanager_secret_version" "db_stage" {
  secret_id = "stage/db_creds"
}