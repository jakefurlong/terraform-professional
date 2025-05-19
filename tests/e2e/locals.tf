locals {
  db_credentials = jsondecode(data.aws_secretsmanager_secret_version.db_sandbox.secret_string)
}