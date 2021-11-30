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

	cfg := make(map[string]string)

	cfg["name"] = fmt.Sprintf("terratest-vpc-%s", random.UniqueId())
	cfg["cidr"] = "10.0.1.0/24"
	cfg["domain"] = "corp.local"
	cfg["ip"] = "<your ip>"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"vpc-name": cfg["name"],
			"cidr":     cfg["cidr"],
			"domain":   cfg["domain"],
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	id := terraform.Output(t, terraformOptions, "vpc_id")
	publicIp := terraform.Output(t, terraformOptions, "vpc_nat_ip")

	assert.Equal(t, cfg["ip"], publicIp)
	assert.NotEmpty(t, id)

}
