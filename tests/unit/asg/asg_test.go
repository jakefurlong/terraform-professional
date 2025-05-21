package rds_test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestASG(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: ".",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// ASG Name
	asg, err := terraform.OutputE(t, terraformOptions, "aws_autoscaling_group_name")
	assert.NoError(t, err, "Failed to get ASG name.")
	t.Logf("ASG Name: %s", asg)
	assert.True(t, strings.HasPrefix(asg, "nimbusdevops-test"))
}
