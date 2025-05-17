package test

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
		TerraformDir: "./",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcID, err := terraform.OutputE(t, terraformOptions, "vpc_id")
	if err != nil {
		t.Fatalf("Failed to get vpc_id output: %v", err)
	}
	t.Logf("vpc_id output: %s", vpcID)

	igwID := terraform.Output(t, terraformOptions, "internet_gateway_id")
	assert.True(t, strings.HasPrefix(igwID, "igw-"))

	rtID := terraform.Output(t, terraformOptions, "route_table_id")
	assert.True(t, strings.HasPrefix(rtID, "rtb-"))

	subnetIDs := terraform.OutputList(t, terraformOptions, "public_subnet_ids")
	assert.Equal(t, len(azs), len(subnetIDs), "Expected one subnet per AZ")
	for _, id := range subnetIDs {
		assert.True(t, strings.HasPrefix(id, "subnet-"))
	}

	assocIDs := terraform.OutputList(t, terraformOptions, "route_table_association_ids")
	assert.Equal(t, 2, len(assocIDs), "Should associate each subnet with the route table")

	sgID := terraform.Output(t, terraformOptions, "security_group_id")
	assert.True(t, strings.HasPrefix(sgID, "sg-"))
}
