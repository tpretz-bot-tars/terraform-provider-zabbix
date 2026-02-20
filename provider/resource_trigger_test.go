package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTrigger(t *testing.T) {
	t.Skip("Zabbix 6.0.44 rejects trigger expressions referencing script[...] with quotes")
	id := resource.UniqueId()
	groupName := "test-group-" + id
	tmplHost := "test-template-" + id

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "zabbix_hostgroup" "testgrp" {
  name = %q
}
resource "zabbix_template" "testtmpl" {
  groups = [zabbix_hostgroup.testgrp.id]
  host   = %q
}
resource "zabbix_item_simple" "testitem" {
  hostid = zabbix_template.testtmpl.id
  key = "script[\"abc\"]"

  name = "Ext Item"
  valuetype = "text"
}

resource "zabbix_trigger" "testtrg" {
  name = "test-trigger"
  expression = "{%s:script[\"abc\"].last()}=0"
  priority = "warn"
  enabled = true
}
`, groupName, tmplHost, tmplHost),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "name", "test-trigger"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "priority", "warn"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "enabled", "true"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "zabbix_hostgroup" "testgrp" {
  name = %q
}
resource "zabbix_template" "testtmpl" {
  groups = [zabbix_hostgroup.testgrp.id]
  host   = %q
}
resource "zabbix_item_simple" "testitem" {
  hostid = zabbix_template.testtmpl.id
  key = "script[\"abc\"]"

  name = "Ext Item"
  valuetype = "text"
}

resource "zabbix_trigger" "testtrg" {
  name = "test-trigger-a"
  expression = "{%s:script[\"abc\"].last()}=1"
  priority = "high"
  enabled = false
}
`, groupName, tmplHost, tmplHost),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "name", "test-trigger-a"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "priority", "high"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "enabled", "false"),
				),
			},
		},
	})
}
