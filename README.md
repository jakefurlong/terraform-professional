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
### Pre-commit notes

From the root directory, perform the following:
*Note:* if you're running Windows, you'll need to use Git Bash or WSL for /bin/bash requirement

```
pip install pre-commit
pre-commit install
pre-commit run --all-files # or dont' run this, it will run on commit...
```

Look at unit test for PR
Look at validating unit instead of running unit tests
Look into pre-commit hook
Check out Atlantis
