package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceTrigger(t *testing.T) {
	// Use the documented expression format (Zabbix 6.x): function(/host/key,params)<op><constant>
	// Avoid item keys with quoted parameters inside expressions (e.g. script["abc"]) as Zabbix may reject them.
	id := resource.UniqueId()
	groupName := "test-group-" + id
	hostName := "test-host-" + id

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

resource "zabbix_host" "testhost" {
  host = %q
  groups = [zabbix_hostgroup.testgrp.id]

  interface {
    type = "agent"
    dns  = "localhost"
    port = 10050
  }
}

resource "zabbix_item_trapper" "testitem" {
  hostid = zabbix_host.testhost.id
  key = "trapper.ping"

  name = "Trapper Item"
  valuetype = "unsigned"
}

resource "zabbix_trigger" "testtrg" {
  name = "test-trigger"
  expression = "last(/%s/trapper.ping)=0"
  priority = "warn"
  enabled = true

  depends_on = [zabbix_item_trapper.testitem]
}
`, groupName, hostName, hostName),
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

resource "zabbix_host" "testhost" {
  host = %q
  groups = [zabbix_hostgroup.testgrp.id]

  interface {
    type = "agent"
    dns  = "localhost"
    port = 10050
  }
}

resource "zabbix_item_trapper" "testitem" {
  hostid = zabbix_host.testhost.id
  key = "trapper.ping"

  name = "Trapper Item"
  valuetype = "unsigned"
}

resource "zabbix_trigger" "testtrg" {
  name = "test-trigger-a"
  expression = "last(/%s/trapper.ping)=1"
  priority = "high"
  enabled = false

  depends_on = [zabbix_item_trapper.testitem]
}
`, groupName, hostName, hostName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "name", "test-trigger-a"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "priority", "high"),
					resource.TestCheckResourceAttr("zabbix_trigger.testtrg", "enabled", "false"),
				),
			},
		},
	})
}
