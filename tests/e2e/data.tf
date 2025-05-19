data "aws_secretsmanager_secret_version" "db_sandbox" {
  secret_id = "test/db_creds"
}