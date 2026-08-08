// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorbatchruleset


type CdnFrontdoorBatchRuleSetRuleConditions struct {
	// client_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#client_port CdnFrontdoorBatchRuleSet#client_port}
	ClientPort interface{} `field:"optional" json:"clientPort" yaml:"clientPort"`
	// device_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#device_type CdnFrontdoorBatchRuleSet#device_type}
	DeviceType interface{} `field:"optional" json:"deviceType" yaml:"deviceType"`
	// host_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#host_name CdnFrontdoorBatchRuleSet#host_name}
	HostName interface{} `field:"optional" json:"hostName" yaml:"hostName"`
	// http_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#http_version CdnFrontdoorBatchRuleSet#http_version}
	HttpVersion interface{} `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// post_argument block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#post_argument CdnFrontdoorBatchRuleSet#post_argument}
	PostArgument interface{} `field:"optional" json:"postArgument" yaml:"postArgument"`
	// query_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#query_string CdnFrontdoorBatchRuleSet#query_string}
	QueryString interface{} `field:"optional" json:"queryString" yaml:"queryString"`
	// remote_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#remote_address CdnFrontdoorBatchRuleSet#remote_address}
	RemoteAddress interface{} `field:"optional" json:"remoteAddress" yaml:"remoteAddress"`
	// request_body block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_body CdnFrontdoorBatchRuleSet#request_body}
	RequestBody interface{} `field:"optional" json:"requestBody" yaml:"requestBody"`
	// request_cookies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_cookies CdnFrontdoorBatchRuleSet#request_cookies}
	RequestCookies interface{} `field:"optional" json:"requestCookies" yaml:"requestCookies"`
	// request_file_extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_file_extension CdnFrontdoorBatchRuleSet#request_file_extension}
	RequestFileExtension interface{} `field:"optional" json:"requestFileExtension" yaml:"requestFileExtension"`
	// request_filename block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_filename CdnFrontdoorBatchRuleSet#request_filename}
	RequestFilename interface{} `field:"optional" json:"requestFilename" yaml:"requestFilename"`
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_header CdnFrontdoorBatchRuleSet#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// request_method block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_method CdnFrontdoorBatchRuleSet#request_method}
	RequestMethod interface{} `field:"optional" json:"requestMethod" yaml:"requestMethod"`
	// request_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_path CdnFrontdoorBatchRuleSet#request_path}
	RequestPath interface{} `field:"optional" json:"requestPath" yaml:"requestPath"`
	// request_scheme block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_scheme CdnFrontdoorBatchRuleSet#request_scheme}
	RequestScheme interface{} `field:"optional" json:"requestScheme" yaml:"requestScheme"`
	// request_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#request_url CdnFrontdoorBatchRuleSet#request_url}
	RequestUrl interface{} `field:"optional" json:"requestUrl" yaml:"requestUrl"`
	// server_port block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#server_port CdnFrontdoorBatchRuleSet#server_port}
	ServerPort interface{} `field:"optional" json:"serverPort" yaml:"serverPort"`
	// socket_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#socket_address CdnFrontdoorBatchRuleSet#socket_address}
	SocketAddress interface{} `field:"optional" json:"socketAddress" yaml:"socketAddress"`
	// ssl_protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/cdn_frontdoor_batch_rule_set#ssl_protocol CdnFrontdoorBatchRuleSet#ssl_protocol}
	SslProtocol interface{} `field:"optional" json:"sslProtocol" yaml:"sslProtocol"`
}

