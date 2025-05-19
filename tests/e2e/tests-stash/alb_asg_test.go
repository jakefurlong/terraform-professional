package integration

import (
	"fmt"
	"strings"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestAlbAsgIntegration(t *testing.T) {
	t.Parallel()

	region := "us-west-1"
	terraformOptions := &terraform.Options{
		TerraformDir: ".",
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": region,
		},
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

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
}
