package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourcesTemplateHostgroupHostProxy(t *testing.T) {
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

resource "zabbix_host" "testhost" {
  host = "test-host"
  interface {
    type = "agent"
    main = true
    ip   = "127.0.0.1"
    dns  = ""
    port = 10050
  }
  groups    = [zabbix_hostgroup.testgrp.id]
  templates = [zabbix_template.testtmpl.id]
}

# Data sources

data "zabbix_hostgroup" "by_name" {
  name = zabbix_hostgroup.testgrp.name
}

data "zabbix_template" "by_host" {
  host = zabbix_template.testtmpl.host
}

data "zabbix_host" "by_host" {
  host = zabbix_host.testhost.host
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.zabbix_hostgroup.by_name", "id"),
					resource.TestCheckResourceAttr("data.zabbix_hostgroup.by_name", "name", "test-group"),

					resource.TestCheckResourceAttrSet("data.zabbix_template.by_host", "id"),
					resource.TestCheckResourceAttr("data.zabbix_template.by_host", "host", "test-template"),

					resource.TestCheckResourceAttrSet("data.zabbix_host.by_host", "id"),
					resource.TestCheckResourceAttr("data.zabbix_host.by_host", "host", "test-host"),
				),
			},
		},
	})
}
