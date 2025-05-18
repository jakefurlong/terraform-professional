data "aws_secretsmanager_secret_version" "db_sandbox" {
  secret_id = "sandbox/jfurlong/db_creds"
}