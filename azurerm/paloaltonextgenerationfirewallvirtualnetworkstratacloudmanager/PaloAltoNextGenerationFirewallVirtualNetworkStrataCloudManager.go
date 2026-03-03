// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkstratacloudmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/paloaltonextgenerationfirewallvirtualnetworkstratacloudmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager azurerm_palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager}.
type PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationNat() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatList
	DestinationNatInput() interface{}
	DnsSettings() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettingsOutputReference
	DnsSettingsInput() *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	Identity() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentityOutputReference
	IdentityInput() *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MarketplaceOfferId() *string
	SetMarketplaceOfferId(val *string)
	MarketplaceOfferIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkProfile() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileOutputReference
	NetworkProfileInput() *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile
	// The tree node.
	Node() constructs.Node
	PlanId() *string
	SetPlanId(val *string)
	PlanIdInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	StrataCloudManagerTenantName() *string
	SetStrataCloudManagerTenantName(val *string)
	StrataCloudManagerTenantNameInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDestinationNat(value interface{})
	PutDnsSettings(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings)
	PutIdentity(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity)
	PutNetworkProfile(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile)
	PutTimeouts(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeouts)
	ResetDestinationNat()
	ResetDnsSettings()
	ResetId()
	ResetIdentity()
	ResetMarketplaceOfferId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlanId()
	ResetTags()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) DestinationNat() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatList {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDestinationNatList
	_jsii_.Get(
		j,
		"destinationNat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) DestinationNatInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationNatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) DnsSettings() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettingsOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettingsOutputReference
	_jsii_.Get(
		j,
		"dnsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) DnsSettingsInput() *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings
	_jsii_.Get(
		j,
		"dnsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Identity() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentityOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) IdentityInput() *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) MarketplaceOfferId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceOfferId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) MarketplaceOfferIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"marketplaceOfferIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) NetworkProfile() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfileOutputReference
	_jsii_.Get(
		j,
		"networkProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) NetworkProfileInput() *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile
	_jsii_.Get(
		j,
		"networkProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) PlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) PlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) StrataCloudManagerTenantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strataCloudManagerTenantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) StrataCloudManagerTenantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strataCloudManagerTenantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) Timeouts() PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeoutsOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager azurerm_palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager} Resource.
func NewPaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager(scope constructs.Construct, id *string, config *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerConfig) PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.62.1/docs/resources/palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager azurerm_palo_alto_next_generation_firewall_virtual_network_strata_cloud_manager} Resource.
func NewPaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_Override(p PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager, scope constructs.Construct, id *string, config *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetMarketplaceOfferId(val *string) {
	if err := j.validateSetMarketplaceOfferIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"marketplaceOfferId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetPlanId(val *string) {
	if err := j.validateSetPlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetStrataCloudManagerTenantName(val *string) {
	if err := j.validateSetStrataCloudManagerTenantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strataCloudManagerTenantName",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager resource upon running "cdktn plan <stack-name>".
func PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validatePaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager.PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) PutDestinationNat(value interface{}) {
	if err := p.validatePutDestinationNatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDestinationNat",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) PutDnsSettings(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerDnsSettings) {
	if err := p.validatePutDnsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putDnsSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) PutIdentity(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerIdentity) {
	if err := p.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putIdentity",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) PutNetworkProfile(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerNetworkProfile) {
	if err := p.validatePutNetworkProfileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNetworkProfile",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) PutTimeouts(value *PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManagerTimeouts) {
	if err := p.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetDestinationNat() {
	_jsii_.InvokeVoid(
		p,
		"resetDestinationNat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetDnsSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetDnsSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetIdentity() {
	_jsii_.InvokeVoid(
		p,
		"resetIdentity",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetMarketplaceOfferId() {
	_jsii_.InvokeVoid(
		p,
		"resetMarketplaceOfferId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetPlanId() {
	_jsii_.InvokeVoid(
		p,
		"resetPlanId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ResetTimeouts() {
	_jsii_.InvokeVoid(
		p,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkStrataCloudManager) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		p,
		"with",
		args,
		&returns,
	)

	return returns
}

