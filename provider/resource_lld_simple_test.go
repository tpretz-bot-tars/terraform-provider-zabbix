package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceLLDSimple(t *testing.T) {
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

resource "zabbix_lld_simple" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.simple.discovery"
  name   = "LLD Simple Rule"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_simple.testrule", "key", "lld.simple.discovery"),
					resource.TestCheckResourceAttr("zabbix_lld_simple.testrule", "name", "LLD Simple Rule"),
					resource.TestCheckResourceAttrSet("zabbix_lld_simple.testrule", "hostid"),
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

resource "zabbix_lld_simple" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.simple.discovery2"
  name   = "LLD Simple Rule A"
  delay  = "600"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_simple.testrule", "key", "lld.simple.discovery2"),
					resource.TestCheckResourceAttr("zabbix_lld_simple.testrule", "name", "LLD Simple Rule A"),
					resource.TestCheckResourceAttr("zabbix_lld_simple.testrule", "delay", "600"),
				),
			},
		},
	})
}
