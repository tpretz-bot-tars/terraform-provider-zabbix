package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceLLDInternal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" {
  name = "test-group"
}
resource "zabbix_template" "testtmpl" {
  groups = [zabbix_hostgroup.testgrp.id]
  host   = "test-template"
}

resource "zabbix_lld_internal" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.internal.discovery"
  name   = "LLD Internal Rule"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_internal.testrule", "key", "lld.internal.discovery"),
					resource.TestCheckResourceAttr("zabbix_lld_internal.testrule", "name", "LLD Internal Rule"),
				),
			},
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" {
  name = "test-group"
}
resource "zabbix_template" "testtmpl" {
  groups = [zabbix_hostgroup.testgrp.id]
  host   = "test-template"
}

resource "zabbix_lld_internal" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.internal.discovery2"
  name   = "LLD Internal Rule A"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_internal.testrule", "key", "lld.internal.discovery2"),
					resource.TestCheckResourceAttr("zabbix_lld_internal.testrule", "name", "LLD Internal Rule A"),
				),
			},
		},
	})
}
