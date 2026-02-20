package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProtoItemSimple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }

resource "zabbix_lld_simple" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.simple.discovery"
  name   = "LLD Simple Rule"
}

resource "zabbix_proto_item_simple" "testitem" {
  ruleid = zabbix_lld_simple.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "script[\"abc\"]"
  name = "Proto Simple Item"
  valuetype = "text"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("zabbix_proto_item_simple.testitem", "ruleid"),
					resource.TestCheckResourceAttr("zabbix_proto_item_simple.testitem", "name", "Proto Simple Item"),
				),
			},
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }

resource "zabbix_lld_simple" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.simple.discovery"
  name   = "LLD Simple Rule"
}

resource "zabbix_proto_item_simple" "testitem" {
  ruleid = zabbix_lld_simple.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "scriptb[\"abc\"]"
  name = "Proto Simple Item A"
  valuetype = "unsigned"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_proto_item_simple.testitem", "key", "scriptb[\"abc\"]"),
					resource.TestCheckResourceAttr("zabbix_proto_item_simple.testitem", "name", "Proto Simple Item A"),
					resource.TestCheckResourceAttr("zabbix_proto_item_simple.testitem", "valuetype", "unsigned"),
				),
			},
		},
	})
}
