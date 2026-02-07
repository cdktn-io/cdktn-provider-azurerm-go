// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkstratacloudmanager

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "destinationNat", GoGetter: "DestinationNat"},
			_jsii_.MemberProperty{JsiiProperty: "destinationNatInput", GoGetter: "DestinationNatInput"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSettings", GoGetter: "DnsSettings"},
			_jsii_.MemberProperty{JsiiProperty: "dnsSettingsInput", GoGetter: "DnsSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identity", GoGetter: "Identity"},
			_jsii_.MemberProperty{JsiiProperty: "identityInput", GoGetter: "IdentityInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "marketplaceOfferId", GoGetter: "MarketplaceOfferId"},
			_jsii_.MemberProperty{JsiiProperty: "marketplaceOfferIdInput", GoGetter: "MarketplaceOfferIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkProfile", GoGetter: "NetworkProfile"},
			_jsii_.MemberProperty{JsiiProperty: "networkProfileInput", GoGetter: "NetworkProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "planId", GoGetter: "PlanId"},
			_jsii_.MemberProperty{JsiiProperty: "planIdInput", GoGetter: "PlanIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putDestinationNat", GoMethod: "PutDestinationNat"},
			_jsii_.MemberMethod{JsiiMethod: "putDnsSettings", GoMethod: "PutDnsSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentity", GoMethod: "PutIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkProfile", GoMethod: "PutNetworkProfile"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationNat", GoMethod: "ResetDestinationNat"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsSettings", GoMethod: "ResetDnsSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentity", GoMethod: "ResetIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "resetMarketplaceOfferId", GoMethod: "ResetMarketplaceOfferId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlanId", GoMethod: "ResetPlanId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "strataCloudManagerTenantName", GoGetter: "StrataCloudManagerTenantName"},
			_jsii_.MemberProperty{JsiiProperty: "strataCloudManagerTenantNameInput", GoGetter: "StrataCloudManagerTenantNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerConfig",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNat",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNat)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatBackendConfig",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatBackendConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatBackendConfigOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatBackendConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddress", GoGetter: "PublicIpAddress"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddressInput", GoGetter: "PublicIpAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatBackendConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatFrontendConfig",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatFrontendConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatFrontendConfigOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatFrontendConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddressId", GoGetter: "PublicIpAddressId"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddressIdInput", GoGetter: "PublicIpAddressIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatFrontendConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatList",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "backendConfig", GoGetter: "BackendConfig"},
			_jsii_.MemberProperty{JsiiProperty: "backendConfigInput", GoGetter: "BackendConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "frontendConfig", GoGetter: "FrontendConfig"},
			_jsii_.MemberProperty{JsiiProperty: "frontendConfigInput", GoGetter: "FrontendConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "putBackendConfig", GoMethod: "PutBackendConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putFrontendConfig", GoMethod: "PutFrontendConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackendConfig", GoMethod: "ResetBackendConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetFrontendConfig", GoMethod: "ResetFrontendConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettingsOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "azureDnsServers", GoGetter: "AzureDnsServers"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServers", GoGetter: "DnsServers"},
			_jsii_.MemberProperty{JsiiProperty: "dnsServersInput", GoGetter: "DnsServersInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsServers", GoMethod: "ResetDnsServers"},
			_jsii_.MemberMethod{JsiiMethod: "resetUseAzureDns", GoMethod: "ResetUseAzureDns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "useAzureDns", GoGetter: "UseAzureDns"},
			_jsii_.MemberProperty{JsiiProperty: "useAzureDnsInput", GoGetter: "UseAzureDnsInput"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentityOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "identityIds", GoGetter: "IdentityIds"},
			_jsii_.MemberProperty{JsiiProperty: "identityIdsInput", GoGetter: "IdentityIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "egressNatIpAddresses", GoGetter: "EgressNatIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "egressNatIpAddressIds", GoGetter: "EgressNatIpAddressIds"},
			_jsii_.MemberProperty{JsiiProperty: "egressNatIpAddressIdsInput", GoGetter: "EgressNatIpAddressIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddresses", GoGetter: "PublicIpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddressIds", GoGetter: "PublicIpAddressIds"},
			_jsii_.MemberProperty{JsiiProperty: "publicIpAddressIdsInput", GoGetter: "PublicIpAddressIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putVnetConfiguration", GoMethod: "PutVnetConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetEgressNatIpAddressIds", GoMethod: "ResetEgressNatIpAddressIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedAddressRanges", GoMethod: "ResetTrustedAddressRanges"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedAddressRanges", GoGetter: "TrustedAddressRanges"},
			_jsii_.MemberProperty{JsiiProperty: "trustedAddressRangesInput", GoGetter: "TrustedAddressRangesInput"},
			_jsii_.MemberProperty{JsiiProperty: "vnetConfiguration", GoGetter: "VnetConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "vnetConfigurationInput", GoGetter: "VnetConfigurationInput"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileVnetConfiguration",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileVnetConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileVnetConfigurationOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileVnetConfigurationOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipOfTrustForUserDefinedRoutes", GoGetter: "IpOfTrustForUserDefinedRoutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustedSubnetId", GoMethod: "ResetTrustedSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUntrustedSubnetId", GoMethod: "ResetUntrustedSubnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSubnetId", GoGetter: "TrustedSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "trustedSubnetIdInput", GoGetter: "TrustedSubnetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "untrustedSubnetId", GoGetter: "UntrustedSubnetId"},
			_jsii_.MemberProperty{JsiiProperty: "untrustedSubnetIdInput", GoGetter: "UntrustedSubnetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkIdInput", GoGetter: "VirtualNetworkIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileVnetConfigurationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeouts",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeoutsOutputReference",
		reflect.TypeOf((*PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
