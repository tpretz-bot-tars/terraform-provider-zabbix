package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProtoItemInternal(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }
resource "zabbix_lld_internal" "rule" { hostid = zabbix_template.testtmpl.id key = "lld.internal.discovery" name = "LLD Internal Rule" }

resource "zabbix_proto_item_internal" "testitem" {
  ruleid = zabbix_lld_internal.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "internal.key"
  name = "Proto Internal Item"
  valuetype = "unsigned"
  delay = "1m"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("zabbix_proto_item_internal.testitem", "ruleid"),
				),
			},
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" { name = "test-group" }
resource "zabbix_template" "testtmpl" { groups = [zabbix_hostgroup.testgrp.id] host = "test-template" }
resource "zabbix_lld_internal" "rule" { hostid = zabbix_template.testtmpl.id key = "lld.internal.discovery" name = "LLD Internal Rule" }

resource "zabbix_proto_item_internal" "testitem" {
  ruleid = zabbix_lld_internal.rule.id
  hostid = zabbix_template.testtmpl.id
  key = "internal.key2"
  name = "Proto Internal Item A"
  valuetype = "unsigned"
  delay = "30s"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_proto_item_internal.testitem", "key", "internal.key2"),
				),
			},
		},
	})
}
