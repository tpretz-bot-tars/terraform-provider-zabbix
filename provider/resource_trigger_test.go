package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTrigger(t *testing.T) {
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
resource "zabbix_item_simple" "testitem" {
  hostid = zabbix_template.testtmpl.id
  key = "script[\"abc\"]"

  name = "Ext Item"
  valuetype = "text"
}

resource "zabbix_trigger" "testtrg" {
  name = "test-trigger"
  expression = "{test-template:script[\"abc\"].last()}=0"
  priority = "warn"
  enabled = true
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "name", "test-trigger"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "priority", "warn"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "enabled", "true"),
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
resource "zabbix_item_simple" "testitem" {
  hostid = zabbix_template.testtmpl.id
  key = "script[\"abc\"]"

  name = "Ext Item"
  valuetype = "text"
}

resource "zabbix_trigger" "testtrg" {
  name = "test-trigger-a"
  expression = "{test-template:script[\"abc\"].last()}=1"
  priority = "high"
  enabled = false
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "name", "test-trigger-a"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "priority", "high"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "enabled", "false"),
				),
			},
		},
	})
}
