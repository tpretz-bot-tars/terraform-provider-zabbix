package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProtoItemTrapper(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }
resource "zabbix_lld_trapper" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.trapper.discovery"
  name   = "LLD Trapper Rule"
}
resource "zabbix_proto_item_trapper" "testitem" {
  ruleid = zabbix_lld_trapper.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "trapper.key"
  name = "Proto Trapper Item"
  valuetype = "text"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("zabbix_proto_item_trapper.testitem", "ruleid"),
					resource.TestCheckResourceAttr("zabbix_proto_item_trapper.testitem", "key", "trapper.key"),
				),
			},
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }
resource "zabbix_lld_trapper" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.trapper.discovery"
  name   = "LLD Trapper Rule"
}
resource "zabbix_proto_item_trapper" "testitem" {
  ruleid = zabbix_lld_trapper.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "trapper.key2"
  name = "Proto Trapper Item A"
  valuetype = "unsigned"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_proto_item_trapper.testitem", "key", "trapper.key2"),
					resource.TestCheckResourceAttr("zabbix_proto_item_trapper.testitem", "valuetype", "unsigned"),
				),
			},
		},
	})
}
