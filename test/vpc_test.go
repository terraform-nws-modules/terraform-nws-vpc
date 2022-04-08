package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNwsVpcExample(t *testing.T) {
	t.Parallel()

	const (
		zone           = "central-0"
		cidr4          = "10.0.1.0/24"
		exp_domain     = "personal"
		exp_net_domain = "my.local"
		exp_ip         = "185.185.59.183"
	)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"name":           fmt.Sprintf("terratest-vpc-%s", random.UniqueId()),
			"zone":           zone,
			"cidr4":          cidr4,
			"domain":         exp_domain,
			"network_domain": exp_net_domain,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	id := terraform.Output(t, terraformOptions, "vpc_id")

	assert.NotEmpty(t, id)
}
