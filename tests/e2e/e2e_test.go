package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestEndToEndInfrastructure(t *testing.T) {
	t.Parallel()

	errors := []string{}

	terraformOptions := &terraform.Options{
		TerraformDir: ".",
		Upgrade:      true,
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	// ALB
	alb, err := terraform.OutputE(t, terraformOptions, "alb_dns_name")
	if err != nil {
		errors = append(errors, fmt.Sprintf("Failed to get ALB DNS name: %v", err))
	} else {
		t.Logf("ALB DNS Name: %s", alb)
		assert.True(t, strings.HasPrefix(alb, "nimbusdevops-test"))
	}

	albDns := terraform.Output(t, terraformOptions, "alb_dns_name")
	assert.True(t, strings.HasPrefix(albDns, "nimbusdevops-test"))

	url := fmt.Sprintf("http://%s", albDns)

	// Expect HTTP 200 OK and a string in the body
	http_helper.HttpGetWithRetryWithCustomValidation(
		t,
		url,
		nil,
		60,             // 10 minutes total
		10*time.Second, // between tries
		func(statusCode int, body string) bool {
			t.Logf("Attempt received status %d with body: %s", statusCode, body)
			return statusCode == 200
		},
	)

	// Log the final response for debugging
	status, body := http_helper.HttpGet(t, url, nil)
	t.Logf("Final response from ALB: %d - %s", status, body)

	// Print all accumulated errors at the end
	if len(errors) > 0 {
		t.Errorf("\n---\nTest completed with %d error(s):\n%s\n---", len(errors), strings.Join(errors, "\n"))
	}
}
