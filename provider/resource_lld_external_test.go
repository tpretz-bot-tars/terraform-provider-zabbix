package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceLLDExternal(t *testing.T) {
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

resource "zabbix_lld_external" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.external.discovery"
  name   = "LLD External Rule"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_external.testrule", "key", "lld.external.discovery"),
					resource.TestCheckResourceAttr("zabbix_lld_external.testrule", "name", "LLD External Rule"),
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

resource "zabbix_lld_external" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.external.discovery2"
  name   = "LLD External Rule A"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_external.testrule", "key", "lld.external.discovery2"),
					resource.TestCheckResourceAttr("zabbix_lld_external.testrule", "name", "LLD External Rule A"),
				),
			},
		},
	})
}
