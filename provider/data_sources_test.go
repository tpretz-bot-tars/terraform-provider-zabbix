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
  interfaces {
    type = "agent"
    main = true
    useip = true
    ip = "127.0.0.1"
    dns = ""
    port = "10050"
  }
  groups = [zabbix_hostgroup.testgrp.id]
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

# Proxy datasource is best-effort: default test env may not have any proxys.
# We include a lookup that should return empty id (not error).

data "zabbix_proxy" "by_host" {
  host = "definitely-not-a-proxy"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.zabbix_hostgroup.by_name", "id"),
					resource.TestCheckResourceAttr("data.zabbix_hostgroup.by_name", "name", "test-group"),

					resource.TestCheckResourceAttrSet("data.zabbix_template.by_host", "id"),
					resource.TestCheckResourceAttr("data.zabbix_template.by_host", "host", "test-template"),

					resource.TestCheckResourceAttrSet("data.zabbix_host.by_host", "id"),
					resource.TestCheckResourceAttr("data.zabbix_host.by_host", "host", "test-host"),

					// Proxy should not hard-fail if not found (id should be empty)
					resource.TestCheckResourceAttr("data.zabbix_proxy.by_host", "id", ""),
				),
			},
		},
	})
}
