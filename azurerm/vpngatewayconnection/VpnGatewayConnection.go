// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpngatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/vpngatewayconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/vpn_gateway_connection azurerm_vpn_gateway_connection}.
type VpnGatewayConnection interface {
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
	IdInput() *string
	InternetSecurityEnabled() interface{}
	SetInternetSecurityEnabled(val interface{})
	InternetSecurityEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	RemoteVpnSiteId() *string
	SetRemoteVpnSiteId(val *string)
	RemoteVpnSiteIdInput() *string
	Routing() VpnGatewayConnectionRoutingOutputReference
	RoutingInput() *VpnGatewayConnectionRouting
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VpnGatewayConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrafficSelectorPolicy() VpnGatewayConnectionTrafficSelectorPolicyList
	TrafficSelectorPolicyInput() interface{}
	VpnGatewayId() *string
	SetVpnGatewayId(val *string)
	VpnGatewayIdInput() *string
	VpnLink() VpnGatewayConnectionVpnLinkList
	VpnLinkInput() interface{}
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
	PutRouting(value *VpnGatewayConnectionRouting)
	PutTimeouts(value *VpnGatewayConnectionTimeouts)
	PutTrafficSelectorPolicy(value interface{})
	PutVpnLink(value interface{})
	ResetId()
	ResetInternetSecurityEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRouting()
	ResetTimeouts()
	ResetTrafficSelectorPolicy()
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

// The jsii proxy struct for VpnGatewayConnection
type jsiiProxy_VpnGatewayConnection struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_VpnGatewayConnection) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) InternetSecurityEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) InternetSecurityEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internetSecurityEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RemoteVpnSiteId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVpnSiteId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RemoteVpnSiteIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVpnSiteIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Routing() VpnGatewayConnectionRoutingOutputReference {
	var returns VpnGatewayConnectionRoutingOutputReference
	_jsii_.Get(
		j,
		"routing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) RoutingInput() *VpnGatewayConnectionRouting {
	var returns *VpnGatewayConnectionRouting
	_jsii_.Get(
		j,
		"routingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) Timeouts() VpnGatewayConnectionTimeoutsOutputReference {
	var returns VpnGatewayConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TrafficSelectorPolicy() VpnGatewayConnectionTrafficSelectorPolicyList {
	var returns VpnGatewayConnectionTrafficSelectorPolicyList
	_jsii_.Get(
		j,
		"trafficSelectorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) TrafficSelectorPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSelectorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpnGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnLink() VpnGatewayConnectionVpnLinkList {
	var returns VpnGatewayConnectionVpnLinkList
	_jsii_.Get(
		j,
		"vpnLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnection) VpnLinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpnLinkInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/vpn_gateway_connection azurerm_vpn_gateway_connection} Resource.
func NewVpnGatewayConnection(scope constructs.Construct, id *string, config *VpnGatewayConnectionConfig) VpnGatewayConnection {
	_init_.Initialize()

	if err := validateNewVpnGatewayConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnGatewayConnection{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/vpn_gateway_connection azurerm_vpn_gateway_connection} Resource.
func NewVpnGatewayConnection_Override(v VpnGatewayConnection, scope constructs.Construct, id *string, config *VpnGatewayConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetInternetSecurityEnabled(val interface{}) {
	if err := j.validateSetInternetSecurityEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internetSecurityEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetRemoteVpnSiteId(val *string) {
	if err := j.validateSetRemoteVpnSiteIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteVpnSiteId",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnection)SetVpnGatewayId(val *string) {
	if err := j.validateSetVpnGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpnGatewayId",
		val,
	)
}

// Generates CDKTN code for importing a VpnGatewayConnection resource upon running "cdktn plan <stack-name>".
func VpnGatewayConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVpnGatewayConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
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
func VpnGatewayConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnGatewayConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnGatewayConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnGatewayConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VpnGatewayConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVpnGatewayConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VpnGatewayConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.vpnGatewayConnection.VpnGatewayConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnGatewayConnection) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VpnGatewayConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnGatewayConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnGatewayConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnGatewayConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnGatewayConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnGatewayConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnGatewayConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnGatewayConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VpnGatewayConnection) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutRouting(value *VpnGatewayConnectionRouting) {
	if err := v.validatePutRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putRouting",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutTimeouts(value *VpnGatewayConnectionTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutTrafficSelectorPolicy(value interface{}) {
	if err := v.validatePutTrafficSelectorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTrafficSelectorPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) PutVpnLink(value interface{}) {
	if err := v.validatePutVpnLinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putVpnLink",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetInternetSecurityEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetInternetSecurityEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetRouting() {
	_jsii_.InvokeVoid(
		v,
		"resetRouting",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) ResetTrafficSelectorPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetTrafficSelectorPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnection) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

