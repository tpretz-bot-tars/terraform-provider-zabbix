package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceLLDHttp(t *testing.T) {
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

resource "zabbix_lld_http" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.http.discovery"
  name   = "LLD HTTP Rule"

  url            = "http://example.com"
  request_method = "get"
  retrieve_mode  = "body"
  status_codes   = "200"
  timeout        = "3s"
  verify_host    = false
  verify_peer    = false
  follow_redirects = true
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "key", "lld.http.discovery"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "name", "LLD HTTP Rule"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "url", "http://example.com"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "request_method", "get"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "retrieve_mode", "body"),
				),
			},
			{
				Config: `
resource "zabbix_hostgroup" "testgrp" {
  name = "test-group"
}
resource "zabbix_template" "testtmpl" {
  groups = [zabbix_hostgroup.testgrp.id]
  host   = "test-template"
}

resource "zabbix_lld_http" "testrule" {
  hostid = zabbix_template.testtmpl.id
  key    = "lld.http.discovery2"
  name   = "LLD HTTP Rule A"

  url            = "http://example.com/abc"
  request_method = "head"
  retrieve_mode  = "headers"
  status_codes   = "200,301"
  timeout        = "5s"
  verify_host    = true
  verify_peer    = true
  follow_redirects = false
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "key", "lld.http.discovery2"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "name", "LLD HTTP Rule A"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "url", "http://example.com/abc"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "request_method", "head"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "retrieve_mode", "headers"),
					resource.TestCheckResourceAttr("zabbix_lld_http.testrule", "follow_redirects", "false"),
				),
			},
		},
	})
}
