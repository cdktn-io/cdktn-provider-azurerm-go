// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot


type WindowsFunctionAppSlotSiteConfigApplicationStack struct {
	// The version of .Net.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#dotnet_version WindowsFunctionAppSlot#dotnet_version}
	DotnetVersion *string `field:"optional" json:"dotnetVersion" yaml:"dotnetVersion"`
	// The version of Java to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#java_version WindowsFunctionAppSlot#java_version}
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// The version of Node to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#node_version WindowsFunctionAppSlot#node_version}
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// The PowerShell Core version to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#powershell_core_version WindowsFunctionAppSlot#powershell_core_version}
	PowershellCoreVersion *string `field:"optional" json:"powershellCoreVersion" yaml:"powershellCoreVersion"`
	// Does the Function App use a custom Application Stack?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#use_custom_runtime WindowsFunctionAppSlot#use_custom_runtime}
	UseCustomRuntime interface{} `field:"optional" json:"useCustomRuntime" yaml:"useCustomRuntime"`
	// Should the DotNet process use an isolated runtime?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/windows_function_app_slot#use_dotnet_isolated_runtime WindowsFunctionAppSlot#use_dotnet_isolated_runtime}
	UseDotnetIsolatedRuntime interface{} `field:"optional" json:"useDotnetIsolatedRuntime" yaml:"useDotnetIsolatedRuntime"`
}

