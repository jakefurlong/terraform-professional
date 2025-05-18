package alb_test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestALB(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: ".",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// ALB
	alb, err := terraform.OutputE(t, terraformOptions, "alb_dns_name")
	assert.NoError(t, err, "Failed to get ALB DNS name.")
	t.Logf("ALB DNS Name: %s", alb)
	assert.True(t, strings.HasPrefix(alb, "nimbusdevops-test"))
}
