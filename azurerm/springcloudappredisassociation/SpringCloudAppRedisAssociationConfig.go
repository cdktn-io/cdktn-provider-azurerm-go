// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package springcloudappredisassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SpringCloudAppRedisAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/spring_cloud_app_redis_association#name SpringCloudAppRedisAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/spring_cloud_app_redis_association#redis_access_key SpringCloudAppRedisAssociation#redis_access_key}.
	RedisAccessKey *string `field:"required" json:"redisAccessKey" yaml:"redisAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/spring_cloud_app_redis_association#redis_cache_id SpringCloudAppRedisAssociation#redis_cache_id}.
	RedisCacheId *string `field:"required" json:"redisCacheId" yaml:"redisCacheId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/spring_cloud_app_redis_association#spring_cloud_app_id SpringCloudAppRedisAssociation#spring_cloud_app_id}.
	SpringCloudAppId *string `field:"required" json:"springCloudAppId" yaml:"springCloudAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/spring_cloud_app_redis_association#id SpringCloudAppRedisAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/spring_cloud_app_redis_association#ssl_enabled SpringCloudAppRedisAssociation#ssl_enabled}.
	SslEnabled interface{} `field:"optional" json:"sslEnabled" yaml:"sslEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/spring_cloud_app_redis_association#timeouts SpringCloudAppRedisAssociation#timeouts}
	Timeouts *SpringCloudAppRedisAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

