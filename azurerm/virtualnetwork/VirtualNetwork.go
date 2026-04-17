// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package virtualnetwork

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/virtualnetwork/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/virtual_network azurerm_virtual_network}.
type VirtualNetwork interface {
	cdktn.TerraformResource
	AddressSpace() *[]*string
	SetAddressSpace(val *[]*string)
	AddressSpaceInput() *[]*string
	BgpCommunity() *string
	SetBgpCommunity(val *string)
	BgpCommunityInput() *string
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
	DdosProtectionPlan() VirtualNetworkDdosProtectionPlanOutputReference
	DdosProtectionPlanInput() *VirtualNetworkDdosProtectionPlan
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	Encryption() VirtualNetworkEncryptionOutputReference
	EncryptionInput() *VirtualNetworkEncryption
	FlowTimeoutInMinutes() *float64
	SetFlowTimeoutInMinutes(val *float64)
	FlowTimeoutInMinutesInput() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Guid() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAddressPool() VirtualNetworkIpAddressPoolList
	IpAddressPoolInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrivateEndpointVnetPolicies() *string
	SetPrivateEndpointVnetPolicies(val *string)
	PrivateEndpointVnetPoliciesInput() *string
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
	Subnet() VirtualNetworkSubnetList
	SubnetInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualNetworkTimeoutsOutputReference
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
	PutDdosProtectionPlan(value *VirtualNetworkDdosProtectionPlan)
	PutEncryption(value *VirtualNetworkEncryption)
	PutIpAddressPool(value interface{})
	PutSubnet(value interface{})
	PutTimeouts(value *VirtualNetworkTimeouts)
	ResetAddressSpace()
	ResetBgpCommunity()
	ResetDdosProtectionPlan()
	ResetDnsServers()
	ResetEdgeZone()
	ResetEncryption()
	ResetFlowTimeoutInMinutes()
	ResetId()
	ResetIpAddressPool()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateEndpointVnetPolicies()
	ResetSubnet()
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

// The jsii proxy struct for VirtualNetwork
type jsiiProxy_VirtualNetwork struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_VirtualNetwork) AddressSpace() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressSpace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) AddressSpaceInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressSpaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) BgpCommunity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpCommunity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) BgpCommunityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpCommunityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) DdosProtectionPlan() VirtualNetworkDdosProtectionPlanOutputReference {
	var returns VirtualNetworkDdosProtectionPlanOutputReference
	_jsii_.Get(
		j,
		"ddosProtectionPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) DdosProtectionPlanInput() *VirtualNetworkDdosProtectionPlan {
	var returns *VirtualNetworkDdosProtectionPlan
	_jsii_.Get(
		j,
		"ddosProtectionPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Encryption() VirtualNetworkEncryptionOutputReference {
	var returns VirtualNetworkEncryptionOutputReference
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) EncryptionInput() *VirtualNetworkEncryption {
	var returns *VirtualNetworkEncryption
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) FlowTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) FlowTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"flowTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Guid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) IpAddressPool() VirtualNetworkIpAddressPoolList {
	var returns VirtualNetworkIpAddressPoolList
	_jsii_.Get(
		j,
		"ipAddressPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) IpAddressPoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipAddressPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) PrivateEndpointVnetPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointVnetPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) PrivateEndpointVnetPoliciesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointVnetPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Subnet() VirtualNetworkSubnetList {
	var returns VirtualNetworkSubnetList
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) SubnetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) Timeouts() VirtualNetworkTimeoutsOutputReference {
	var returns VirtualNetworkTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetwork) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/virtual_network azurerm_virtual_network} Resource.
func NewVirtualNetwork(scope constructs.Construct, id *string, config *VirtualNetworkConfig) VirtualNetwork {
	_init_.Initialize()

	if err := validateNewVirtualNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualNetwork{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.virtualNetwork.VirtualNetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/virtual_network azurerm_virtual_network} Resource.
func NewVirtualNetwork_Override(v VirtualNetwork, scope constructs.Construct, id *string, config *VirtualNetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.virtualNetwork.VirtualNetwork",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetAddressSpace(val *[]*string) {
	if err := j.validateSetAddressSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressSpace",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetBgpCommunity(val *string) {
	if err := j.validateSetBgpCommunityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpCommunity",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetFlowTimeoutInMinutes(val *float64) {
	if err := j.validateSetFlowTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flowTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetPrivateEndpointVnetPolicies(val *string) {
	if err := j.validateSetPrivateEndpointVnetPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointVnetPolicies",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetwork)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a VirtualNetwork resource upon running "cdktn plan <stack-name>".
func VirtualNetwork_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualNetwork_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetwork.VirtualNetwork",
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
func VirtualNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetwork.VirtualNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetwork.VirtualNetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetwork_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetwork.VirtualNetwork",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualNetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.virtualNetwork.VirtualNetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualNetwork) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualNetwork) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualNetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualNetwork) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualNetwork) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualNetwork) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualNetwork) PutDdosProtectionPlan(value *VirtualNetworkDdosProtectionPlan) {
	if err := v.validatePutDdosProtectionPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDdosProtectionPlan",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetwork) PutEncryption(value *VirtualNetworkEncryption) {
	if err := v.validatePutEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putEncryption",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetwork) PutIpAddressPool(value interface{}) {
	if err := v.validatePutIpAddressPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpAddressPool",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetwork) PutSubnet(value interface{}) {
	if err := v.validatePutSubnetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putSubnet",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetwork) PutTimeouts(value *VirtualNetworkTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetAddressSpace() {
	_jsii_.InvokeVoid(
		v,
		"resetAddressSpace",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetBgpCommunity() {
	_jsii_.InvokeVoid(
		v,
		"resetBgpCommunity",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetDdosProtectionPlan() {
	_jsii_.InvokeVoid(
		v,
		"resetDdosProtectionPlan",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetDnsServers() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsServers",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		v,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetEncryption() {
	_jsii_.InvokeVoid(
		v,
		"resetEncryption",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetFlowTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		v,
		"resetFlowTimeoutInMinutes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetIpAddressPool() {
	_jsii_.InvokeVoid(
		v,
		"resetIpAddressPool",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetPrivateEndpointVnetPolicies() {
	_jsii_.InvokeVoid(
		v,
		"resetPrivateEndpointVnetPolicies",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetSubnet() {
	_jsii_.InvokeVoid(
		v,
		"resetSubnet",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetwork) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		v,
		"with",
		args,
		&returns,
	)

	return returns
}

