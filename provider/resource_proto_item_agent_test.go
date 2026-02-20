package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProtoItemAgent(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }
resource "zabbix_lld_agent" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.agent.discovery"
  name   = "LLD Agent Rule"
}
resource "zabbix_proto_item_agent" "testitem" {
  ruleid = zabbix_lld_agent.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "agent.ping"
  name = "Proto Agent Item"
  valuetype = "unsigned"
  active = false
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("zabbix_proto_item_agent.testitem", "ruleid"),
					resource.TestCheckResourceAttr("zabbix_proto_item_agent.testitem", "key", "agent.ping"),
					resource.TestCheckResourceAttr("zabbix_proto_item_agent.testitem", "active", "false"),
				),
			},
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }
resource "zabbix_lld_agent" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.agent.discovery"
  name   = "LLD Agent Rule"
}
resource "zabbix_proto_item_agent" "testitem" {
  ruleid = zabbix_lld_agent.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "agent.ping2"
  name = "Proto Agent Item A"
  valuetype = "unsigned"
  active = true
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_proto_item_agent.testitem", "key", "agent.ping2"),
					resource.TestCheckResourceAttr("zabbix_proto_item_agent.testitem", "active", "true"),
				),
			},
		},
	})
}
