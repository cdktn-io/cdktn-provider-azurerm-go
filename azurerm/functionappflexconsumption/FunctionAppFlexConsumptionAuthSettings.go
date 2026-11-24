// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionAuthSettings struct {
	// Should the Authentication / Authorization feature be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#enabled FunctionAppFlexConsumption#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#active_directory FunctionAppFlexConsumption#active_directory}
	ActiveDirectory *FunctionAppFlexConsumptionAuthSettingsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// A map of Login Parameters to send to the OpenID Connect authorization endpoint when a user logs in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#additional_login_parameters FunctionAppFlexConsumption#additional_login_parameters}
	AdditionalLoginParameters *map[string]*string `field:"optional" json:"additionalLoginParameters" yaml:"additionalLoginParameters"`
	// A list of External URLs that can be redirected to as part of logging in or logging out of the Windows Web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#allowed_external_redirect_urls FunctionAppFlexConsumption#allowed_external_redirect_urls}
	AllowedExternalRedirectUrls *[]*string `field:"optional" json:"allowedExternalRedirectUrls" yaml:"allowedExternalRedirectUrls"`
	// The default authentication provider to use when multiple providers are configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#default_provider FunctionAppFlexConsumption#default_provider}
	DefaultProvider *string `field:"optional" json:"defaultProvider" yaml:"defaultProvider"`
	// facebook block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#facebook FunctionAppFlexConsumption#facebook}
	Facebook *FunctionAppFlexConsumptionAuthSettingsFacebook `field:"optional" json:"facebook" yaml:"facebook"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#github FunctionAppFlexConsumption#github}
	Github *FunctionAppFlexConsumptionAuthSettingsGithub `field:"optional" json:"github" yaml:"github"`
	// google block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#google FunctionAppFlexConsumption#google}
	Google *FunctionAppFlexConsumptionAuthSettingsGoogle `field:"optional" json:"google" yaml:"google"`
	// The OpenID Connect Issuer URI that represents the entity which issues access tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#issuer FunctionAppFlexConsumption#issuer}
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// microsoft block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#microsoft FunctionAppFlexConsumption#microsoft}
	Microsoft *FunctionAppFlexConsumptionAuthSettingsMicrosoft `field:"optional" json:"microsoft" yaml:"microsoft"`
	// The RuntimeVersion of the Authentication / Authorization feature in use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#runtime_version FunctionAppFlexConsumption#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
	// The number of hours after session token expiration that a session token can be used to call the token refresh API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#token_refresh_extension_hours FunctionAppFlexConsumption#token_refresh_extension_hours}
	TokenRefreshExtensionHours *float64 `field:"optional" json:"tokenRefreshExtensionHours" yaml:"tokenRefreshExtensionHours"`
	// Should the Windows Web App durably store platform-specific security tokens that are obtained during login flows?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#token_store_enabled FunctionAppFlexConsumption#token_store_enabled}
	TokenStoreEnabled interface{} `field:"optional" json:"tokenStoreEnabled" yaml:"tokenStoreEnabled"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#twitter FunctionAppFlexConsumption#twitter}
	Twitter *FunctionAppFlexConsumptionAuthSettingsTwitter `field:"optional" json:"twitter" yaml:"twitter"`
	// The action to take when an unauthenticated client attempts to access the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/function_app_flex_consumption#unauthenticated_client_action FunctionAppFlexConsumption#unauthenticated_client_action}
	UnauthenticatedClientAction *string `field:"optional" json:"unauthenticatedClientAction" yaml:"unauthenticatedClientAction"`
}

