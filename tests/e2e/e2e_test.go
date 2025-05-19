package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVPCModule(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: ".",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	assert.Contains(t, vpcId, "vpc-", "Expected VPC ID to be returned")

	privateSubnets := terraform.OutputList(t, terraformOptions, "private_subnet_ids")
	assert.Greater(t, len(privateSubnets), 0, "Expected at least one private subnet")

	publicSubnets := terraform.OutputList(t, terraformOptions, "public_subnet_ids")
	assert.Greater(t, len(publicSubnets), 0, "Expected at least one public subnet")

	rdsEndpoint := terraform.Output(t, terraformOptions, "rds_endpoint")
	assert.Contains(t, rdsEndpoint, ".rds.amazonaws.com")

	dbIdentifier := terraform.Output(t, terraformOptions, "db_identifier")
	assert.NotEmpty(t, dbIdentifier)

}
