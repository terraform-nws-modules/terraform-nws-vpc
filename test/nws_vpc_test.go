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
		cidr       = "10.0.1.0/24"
		exp_domain = "corp.local"
		exp_ip     = "<your_ip"
	)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"vpc-name": fmt.Sprintf("terratest-vpc-%s", random.UniqueId()),
			"cidr":     cidr,
			"domain":   exp_domain,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	id := terraform.Output(t, terraformOptions, "vpc_id")
	publicIp := terraform.Output(t, terraformOptions, "vpc_nat_ip")

	assert.Equal(t, exp_ip, publicIp)
	assert.NotEmpty(t, id)

}
