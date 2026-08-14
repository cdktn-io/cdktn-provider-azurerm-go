// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule


type CdnFrontdoorRuleConditions struct {
	// client_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#client_port CdnFrontdoorRule#client_port}
	ClientPort interface{} `field:"optional" json:"clientPort" yaml:"clientPort"`
	// device_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#device_type CdnFrontdoorRule#device_type}
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// host_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#host_name CdnFrontdoorRule#host_name}
	HostName interface{} `field:"optional" json:"hostName" yaml:"hostName"`
	// http_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#http_version CdnFrontdoorRule#http_version}
	HttpVersion interface{} `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// post_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#post_argument CdnFrontdoorRule#post_argument}
	PostArgument interface{} `field:"optional" json:"postArgument" yaml:"postArgument"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#query_string CdnFrontdoorRule#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// remote_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#remote_address CdnFrontdoorRule#remote_address}
	RemoteAddress interface{} `field:"optional" json:"remoteAddress" yaml:"remoteAddress"`
	// request_body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_body CdnFrontdoorRule#request_body}
	RequestBody interface{} `field:"optional" json:"requestBody" yaml:"requestBody"`
	// request_cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_cookies CdnFrontdoorRule#request_cookies}
	RequestCookies interface{} `field:"optional" json:"requestCookies" yaml:"requestCookies"`
	// request_file_extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_file_extension CdnFrontdoorRule#request_file_extension}
	RequestFileExtension interface{} `field:"optional" json:"requestFileExtension" yaml:"requestFileExtension"`
	// request_filename block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_filename CdnFrontdoorRule#request_filename}
	RequestFilename interface{} `field:"optional" json:"requestFilename" yaml:"requestFilename"`
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_header CdnFrontdoorRule#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_method CdnFrontdoorRule#request_method}
	RequestMethod interface{} `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// request_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_path CdnFrontdoorRule#request_path}
	RequestPath interface{} `field:"optional" json:"requestPath" yaml:"requestPath"`
	// request_scheme block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_scheme CdnFrontdoorRule#request_scheme}
	RequestScheme interface{} `field:"optional" json:"requestScheme" yaml:"requestScheme"`
	// request_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#request_url CdnFrontdoorRule#request_url}
	RequestUrl interface{} `field:"optional" json:"requestUrl" yaml:"requestUrl"`
	// server_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#server_port CdnFrontdoorRule#server_port}
	ServerPort interface{} `field:"optional" json:"serverPort" yaml:"serverPort"`
	// socket_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#socket_address CdnFrontdoorRule#socket_address}
	SocketAddress interface{} `field:"optional" json:"socketAddress" yaml:"socketAddress"`
	// ssl_protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/cdn_frontdoor_rule#ssl_protocol CdnFrontdoorRule#ssl_protocol}
	SslProtocol interface{} `field:"optional" json:"sslProtocol" yaml:"sslProtocol"`
}

