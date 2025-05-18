# Professional Terraform Layout

Purpose: provide an example of *one* way to set up your Terraform environment that is production-ready.

## Local testing

```
cd /path/to/root/module
terraform init
go mod init github.com/jakefurlong/terraform-professional
go mod tidy
go test -v
```

Testing CI...