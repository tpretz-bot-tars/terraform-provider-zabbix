package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceLLDTrapper(t *testing.T) {
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

resource "zabbix_lld_trapper" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.trapper.discovery"
  name   = "LLD Trapper Rule"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "key", "lld.trapper.discovery"),
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "name", "LLD Trapper Rule"),
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

resource "zabbix_lld_trapper" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.trapper.discovery2"
  name   = "LLD Trapper Rule A"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "key", "lld.trapper.discovery2"),
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "name", "LLD Trapper Rule A"),
				),
			},
		},
	})
}
