// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rds_test

import (
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/rds"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccRDSParameterGroupDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_db_parameter_group.test"
	resourceName := "aws_db_parameter_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, rds.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccParameterGroupDataSourceConfig_nonExistent,
				ExpectError: regexache.MustCompile(`not found`),
			},
			{
				Config: testAccParameterGroupDataSourceConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
					resource.TestCheckResourceAttrPair(datasourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(datasourceName, "family", resourceName, "family"),
					resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
				),
			},
		},
	})
}

const testAccParameterGroupDataSourceConfig_nonExistent = `
data "aws_db_parameter_group" "test" {
  name = "tf-acc-test-does-not-exist"
}
`

func testAccParameterGroupDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_db_parameter_group" "test" {
  name   = %[1]q
  family = "postgres12"

  parameter {
    name         = "client_encoding"
    value        = "UTF8"
    apply_method = "pending-reboot"
  }
}

data "aws_db_parameter_group" "test" {
  name = aws_db_parameter_group.test.name
}
`, rName)
}