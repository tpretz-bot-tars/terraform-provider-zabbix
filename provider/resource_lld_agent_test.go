package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceLLDAgent(t *testing.T) {
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

resource "zabbix_lld_agent" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.agent.discovery"
  name   = "LLD Agent Rule"
  active = false
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_agent.testrule", "key", "lld.agent.discovery"),
					resource.TestCheckResourceAttr("zabbix_lld_agent.testrule", "name", "LLD Agent Rule"),
					resource.TestCheckResourceAttr("zabbix_lld_agent.testrule", "active", "false"),
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

resource "zabbix_lld_agent" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.agent.discovery2"
  name   = "LLD Agent Rule A"
  active = true
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_agent.testrule", "key", "lld.agent.discovery2"),
					resource.TestCheckResourceAttr("zabbix_lld_agent.testrule", "name", "LLD Agent Rule A"),
					resource.TestCheckResourceAttr("zabbix_lld_agent.testrule", "active", "true"),
				),
			},
		},
	})
}
