package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProtoItemAggregate(t *testing.T) {
	t.Skip("Zabbix 6.0.44 rejects aggregate prototypes (type 8)")
	id := resource.UniqueId()
	groupName := "test-group-" + id
	tmplHost := "test-template-" + id

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
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
resource "zabbix_lld_simple" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.simple.discovery"
  name   = "LLD Simple Rule"
}

resource "zabbix_proto_item_aggregate" "testitem" {
  ruleid = zabbix_lld_simple.rule.id
  hostid = zabbix_template.testtmpl.id
  key    = "aggregate.key"
  name   = "Proto Aggregate Item"
  valuetype = "unsigned"
  delay = "1m"
  type = "0"
}
`, groupName, tmplHost),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_proto_item_aggregate.testitem", "delay", "1m"),
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
resource "zabbix_lld_simple" "rule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.simple.discovery"
  name   = "LLD Simple Rule"
}

resource "zabbix_proto_item_aggregate" "testitem" {
  ruleid = zabbix_lld_simple.rule.id
  hostid = zabbix_template.testtmpl.id
  key    = "aggregate.key2"
  name   = "Proto Aggregate Item A"
  valuetype = "unsigned"
  delay = "30s"
  type = "0"
}
`, groupName, tmplHost),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_proto_item_aggregate.testitem", "key", "aggregate.key2"),
					resource.TestCheckResourceAttr("zabbix_proto_item_aggregate.testitem", "delay", "30s"),
				),
			},
		},
	})
}
