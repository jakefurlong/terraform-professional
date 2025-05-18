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

	// VPC ID
	vpcID, err := terraform.OutputE(t, terraformOptions, "vpc_id")
	assert.NoError(t, err, "Failed to get vpc_id")
	t.Logf("vpc_id: %s", vpcID)
	assert.True(t, strings.HasPrefix(vpcID, "vpc-"))

	// IGW
	igwID, err := terraform.OutputE(t, terraformOptions, "internet_gateway_id")
	assert.NoError(t, err, "Failed to get internet_gateway_id")
	t.Logf("internet_gateway_id: %s", igwID)
	assert.True(t, strings.HasPrefix(igwID, "igw-"))

	// Route Table
	rtID, err := terraform.OutputE(t, terraformOptions, "route_table_id")
	assert.NoError(t, err, "Failed to get route_table_id")
	t.Logf("route_table_id: %s", rtID)
	assert.True(t, strings.HasPrefix(rtID, "rtb-"))

	// Subnet IDs
	subnetIDs, err := terraform.OutputListE(t, terraformOptions, "public_subnet_ids")
	assert.NoError(t, err, "Failed to get public_subnet_ids")
	t.Logf("public_subnet_ids: %v", subnetIDs)
	assert.Equal(t, len(azs), len(subnetIDs), "Expected one subnet per AZ")
	for _, id := range subnetIDs {
		assert.True(t, strings.HasPrefix(id, "subnet-"))
	}

	// RTAs
	rtaIDs, err := terraform.OutputListE(t, terraformOptions, "route_table_association_ids")
	assert.NoError(t, err, "Failed to get route_table_association_ids")
	t.Logf("route_table_association_ids: %v", rtaIDs)
	assert.Equal(t, len(azs), len(rtaIDs), "Expected one route table association per subnet")

	// SG
	sgID, err := terraform.OutputE(t, terraformOptions, "security_group_id")
	assert.NoError(t, err, "Failed to get security_group_id")
	t.Logf("security_group_id: %s", sgID)
	assert.True(t, strings.HasPrefix(sgID, "sg-"))
}
