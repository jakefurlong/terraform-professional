package vpc_test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestCustomVPC(t *testing.T) {
	t.Parallel()

	azs := []string{"us-west-1b", "us-west-1c"}

	terraformOptions := &terraform.Options{
		TerraformDir: ".",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// VPC Name
	vpcName, err := terraform.OutputE(t, terraformOptions, "vpc_name")
	assert.NoError(t, err, "Failed to get vpc_name")
	t.Logf("vpc_name: %s", vpcName)
	assert.True(t, strings.HasPrefix(vpcName, "nimbusdevops-test"))

	// IGW Name
	igwName, err := terraform.OutputE(t, terraformOptions, "internet_gateway_name")
	assert.NoError(t, err, "Failed to get internet_gateway_name")
	t.Logf("internet_gateway_name: %s", igwName)
	assert.True(t, strings.HasPrefix(igwName, "nimbusdevops-test"))

	// Route Table Name
	rtName, err := terraform.OutputE(t, terraformOptions, "aws_route_table_public_name")
	assert.NoError(t, err, "Failed to get aws_route_table_public_name")
	t.Logf("aws_route_table_public_name: %s", rtName)
	assert.True(t, strings.HasPrefix(rtName, "nimbusdevops-test"))

	// Subnet IDs
	subnetIDs, err := terraform.OutputListE(t, terraformOptions, "public_subnet_ids")
	assert.NoError(t, err, "Failed to get public_subnet_ids")
	t.Logf("public_subnet_ids: %v", subnetIDs)
	assert.Equal(t, len(azs), len(subnetIDs), "Expected one subnet per AZ")
	for _, id := range subnetIDs {
		assert.NotEmpty(t, id, "Subnet ID should not be empty")
	}

	// RTAs
	rtaIDs, err := terraform.OutputListE(t, terraformOptions, "route_table_association_ids")
	assert.NoError(t, err, "Failed to get route_table_association_ids")
	t.Logf("route_table_association_ids: %v", rtaIDs)
	assert.Equal(t, len(azs), len(rtaIDs), "Expected one route table association per subnet")

	// SG Name
	sgName, err := terraform.OutputE(t, terraformOptions, "security_group_name")
	assert.NoError(t, err, "Failed to get security_group_name")
	t.Logf("security_group_name: %s", sgName)
	assert.True(t, strings.HasPrefix(sgName, "nimbusdevops-test"))
}
