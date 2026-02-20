package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceLLDTrapper(t *testing.T) {
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

resource "zabbix_lld_trapper" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.trapper.discovery"
  name   = "LLD Trapper Rule"
  delay  = "3600"
}
`, groupName, tmplHost),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "key", "lld.trapper.discovery"),
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "name", "LLD Trapper Rule"),
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

resource "zabbix_lld_trapper" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.trapper.discovery2"
  name   = "LLD Trapper Rule A"
  delay  = "3600"
}
`, groupName, tmplHost),
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "key", "lld.trapper.discovery2"),
					resource.TestCheckResourceAttr("zabbix_lld_trapper.testrule", "name", "LLD Trapper Rule A"),
				),
			},
		},
	})
}
