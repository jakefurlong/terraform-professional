package rds_test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRDSDatabaseAddress(t *testing.T) {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: ".",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// RDS Database Address
	databaseAddress, err := terraform.OutputE(t, terraformOptions, "database_address")
	assert.NoError(t, err, "Failed to get database address.")
	t.Logf("database_address: %s", databaseAddress)
	assert.True(t, strings.HasPrefix(databaseAddress, "nimbusdevops-test"))
	assert.Contains(t, databaseAddress, ".rds.amazonaws.com")
}
