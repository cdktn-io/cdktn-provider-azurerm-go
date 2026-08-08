// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/virtualnetworkgatewayconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection}.
type VirtualNetworkGatewayConnection interface {
	cdktn.TerraformResource
	AuthorizationKey() *string
	SetAuthorizationKey(val *string)
	AuthorizationKeyInput() *string
	BgpEnabled() interface{}
	SetBgpEnabled(val interface{})
	BgpEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionMode() *string
	SetConnectionMode(val *string)
	ConnectionModeInput() *string
	ConnectionProtocol() *string
	SetConnectionProtocol(val *string)
	ConnectionProtocolInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomBgpAddresses() VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference
	CustomBgpAddressesInput() *VirtualNetworkGatewayConnectionCustomBgpAddresses
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DpdTimeoutSeconds() *float64
	SetDpdTimeoutSeconds(val *float64)
	DpdTimeoutSecondsInput() *float64
	EgressNatRuleIds() *[]*string
	SetEgressNatRuleIds(val *[]*string)
	EgressNatRuleIdsInput() *[]*string
	ExpressRouteCircuitId() *string
	SetExpressRouteCircuitId(val *string)
	ExpressRouteCircuitIdInput() *string
	ExpressRouteGatewayBypass() interface{}
	SetExpressRouteGatewayBypass(val interface{})
	ExpressRouteGatewayBypassInput() interface{}
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
	IngressNatRuleIds() *[]*string
	SetIngressNatRuleIds(val *[]*string)
	IngressNatRuleIdsInput() *[]*string
	IpsecPolicy() VirtualNetworkGatewayConnectionIpsecPolicyOutputReference
	IpsecPolicyInput() *VirtualNetworkGatewayConnectionIpsecPolicy
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocalAzureIpAddressEnabled() interface{}
	SetLocalAzureIpAddressEnabled(val interface{})
	LocalAzureIpAddressEnabledInput() interface{}
	LocalNetworkGatewayId() *string
	SetLocalNetworkGatewayId(val *string)
	LocalNetworkGatewayIdInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PeerVirtualNetworkGatewayId() *string
	SetPeerVirtualNetworkGatewayId(val *string)
	PeerVirtualNetworkGatewayIdInput() *string
	PrivateLinkFastPathEnabled() interface{}
	SetPrivateLinkFastPathEnabled(val interface{})
	PrivateLinkFastPathEnabledInput() interface{}
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
	RoutingWeight() *float64
	SetRoutingWeight(val *float64)
	RoutingWeightInput() *float64
	SharedKey() *string
	SetSharedKey(val *string)
	SharedKeyInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() VirtualNetworkGatewayConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrafficSelectorPolicy() VirtualNetworkGatewayConnectionTrafficSelectorPolicyList
	TrafficSelectorPolicyInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UsePolicyBasedTrafficSelectors() interface{}
	SetUsePolicyBasedTrafficSelectors(val interface{})
	UsePolicyBasedTrafficSelectorsInput() interface{}
	VirtualNetworkGatewayId() *string
	SetVirtualNetworkGatewayId(val *string)
	VirtualNetworkGatewayIdInput() *string
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
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
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
	PutCustomBgpAddresses(value *VirtualNetworkGatewayConnectionCustomBgpAddresses)
	PutIpsecPolicy(value *VirtualNetworkGatewayConnectionIpsecPolicy)
	PutTimeouts(value *VirtualNetworkGatewayConnectionTimeouts)
	PutTrafficSelectorPolicy(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetAuthorizationKey()
	ResetBgpEnabled()
	ResetConnectionMode()
	ResetConnectionProtocol()
	ResetCustomBgpAddresses()
	ResetDpdTimeoutSeconds()
	ResetEgressNatRuleIds()
	ResetExpressRouteCircuitId()
	ResetExpressRouteGatewayBypass()
	ResetId()
	ResetIngressNatRuleIds()
	ResetIpsecPolicy()
	ResetLocalAzureIpAddressEnabled()
	ResetLocalNetworkGatewayId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeerVirtualNetworkGatewayId()
	ResetPrivateLinkFastPathEnabled()
	ResetRoutingWeight()
	ResetSharedKey()
	ResetTags()
	ResetTimeouts()
	ResetTrafficSelectorPolicy()
	ResetUsePolicyBasedTrafficSelectors()
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

// The jsii proxy struct for VirtualNetworkGatewayConnection
type jsiiProxy_VirtualNetworkGatewayConnection struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) AuthorizationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) AuthorizationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) BgpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) BgpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConnectionProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) CustomBgpAddresses() VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference {
	var returns VirtualNetworkGatewayConnectionCustomBgpAddressesOutputReference
	_jsii_.Get(
		j,
		"customBgpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) CustomBgpAddressesInput() *VirtualNetworkGatewayConnectionCustomBgpAddresses {
	var returns *VirtualNetworkGatewayConnectionCustomBgpAddresses
	_jsii_.Get(
		j,
		"customBgpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) DpdTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) DpdTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dpdTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) EgressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) EgressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteCircuitId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteCircuitIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressRouteCircuitIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteGatewayBypass() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressRouteGatewayBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ExpressRouteGatewayBypassInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"expressRouteGatewayBypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IngressNatRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IngressNatRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingressNatRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IpsecPolicy() VirtualNetworkGatewayConnectionIpsecPolicyOutputReference {
	var returns VirtualNetworkGatewayConnectionIpsecPolicyOutputReference
	_jsii_.Get(
		j,
		"ipsecPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) IpsecPolicyInput() *VirtualNetworkGatewayConnectionIpsecPolicy {
	var returns *VirtualNetworkGatewayConnectionIpsecPolicy
	_jsii_.Get(
		j,
		"ipsecPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalAzureIpAddressEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalAzureIpAddressEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localAzureIpAddressEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalNetworkGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localNetworkGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocalNetworkGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localNetworkGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) PeerVirtualNetworkGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVirtualNetworkGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) PeerVirtualNetworkGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVirtualNetworkGatewayIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) PrivateLinkFastPathEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkFastPathEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) PrivateLinkFastPathEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateLinkFastPathEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) RoutingWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routingWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) RoutingWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"routingWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SharedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) SharedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Timeouts() VirtualNetworkGatewayConnectionTimeoutsOutputReference {
	var returns VirtualNetworkGatewayConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TrafficSelectorPolicy() VirtualNetworkGatewayConnectionTrafficSelectorPolicyList {
	var returns VirtualNetworkGatewayConnectionTrafficSelectorPolicyList
	_jsii_.Get(
		j,
		"trafficSelectorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TrafficSelectorPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trafficSelectorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) UsePolicyBasedTrafficSelectors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePolicyBasedTrafficSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) UsePolicyBasedTrafficSelectorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePolicyBasedTrafficSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) VirtualNetworkGatewayId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkGatewayId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection) VirtualNetworkGatewayIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkGatewayIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection} Resource.
func NewVirtualNetworkGatewayConnection(scope constructs.Construct, id *string, config *VirtualNetworkGatewayConnectionConfig) VirtualNetworkGatewayConnection {
	_init_.Initialize()

	if err := validateNewVirtualNetworkGatewayConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualNetworkGatewayConnection{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/virtual_network_gateway_connection azurerm_virtual_network_gateway_connection} Resource.
func NewVirtualNetworkGatewayConnection_Override(v VirtualNetworkGatewayConnection, scope constructs.Construct, id *string, config *VirtualNetworkGatewayConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetAuthorizationKey(val *string) {
	if err := j.validateSetAuthorizationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationKey",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetBgpEnabled(val interface{}) {
	if err := j.validateSetBgpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgpEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetConnectionMode(val *string) {
	if err := j.validateSetConnectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionMode",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetConnectionProtocol(val *string) {
	if err := j.validateSetConnectionProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionProtocol",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetDpdTimeoutSeconds(val *float64) {
	if err := j.validateSetDpdTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dpdTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetEgressNatRuleIds(val *[]*string) {
	if err := j.validateSetEgressNatRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetExpressRouteCircuitId(val *string) {
	if err := j.validateSetExpressRouteCircuitIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressRouteCircuitId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetExpressRouteGatewayBypass(val interface{}) {
	if err := j.validateSetExpressRouteGatewayBypassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expressRouteGatewayBypass",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetIngressNatRuleIds(val *[]*string) {
	if err := j.validateSetIngressNatRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressNatRuleIds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLocalAzureIpAddressEnabled(val interface{}) {
	if err := j.validateSetLocalAzureIpAddressEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAzureIpAddressEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLocalNetworkGatewayId(val *string) {
	if err := j.validateSetLocalNetworkGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetPeerVirtualNetworkGatewayId(val *string) {
	if err := j.validateSetPeerVirtualNetworkGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVirtualNetworkGatewayId",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetPrivateLinkFastPathEnabled(val interface{}) {
	if err := j.validateSetPrivateLinkFastPathEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkFastPathEnabled",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetRoutingWeight(val *float64) {
	if err := j.validateSetRoutingWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routingWeight",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetSharedKey(val *string) {
	if err := j.validateSetSharedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedKey",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetUsePolicyBasedTrafficSelectors(val interface{}) {
	if err := j.validateSetUsePolicyBasedTrafficSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePolicyBasedTrafficSelectors",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayConnection)SetVirtualNetworkGatewayId(val *string) {
	if err := j.validateSetVirtualNetworkGatewayIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkGatewayId",
		val,
	)
}

// Generates CDKTN code for importing a VirtualNetworkGatewayConnection resource upon running "cdktn plan <stack-name>".
func VirtualNetworkGatewayConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
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
func VirtualNetworkGatewayConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkGatewayConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func VirtualNetworkGatewayConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVirtualNetworkGatewayConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VirtualNetworkGatewayConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.virtualNetworkGatewayConnection.VirtualNetworkGatewayConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) AddMoveTarget(moveTarget *string) {
	if err := v.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := v.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGatewayConnection) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := v.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) MoveFromId(id *string) {
	if err := v.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveFromId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := v.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) MoveToId(id *string) {
	if err := v.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"moveToId",
		[]interface{}{id},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutCustomBgpAddresses(value *VirtualNetworkGatewayConnectionCustomBgpAddresses) {
	if err := v.validatePutCustomBgpAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCustomBgpAddresses",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutIpsecPolicy(value *VirtualNetworkGatewayConnectionIpsecPolicy) {
	if err := v.validatePutIpsecPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpsecPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutTimeouts(value *VirtualNetworkGatewayConnectionTimeouts) {
	if err := v.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) PutTrafficSelectorPolicy(value interface{}) {
	if err := v.validatePutTrafficSelectorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTrafficSelectorPolicy",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := v.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetAuthorizationKey() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthorizationKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetBgpEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetBgpEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetConnectionMode() {
	_jsii_.InvokeVoid(
		v,
		"resetConnectionMode",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetConnectionProtocol() {
	_jsii_.InvokeVoid(
		v,
		"resetConnectionProtocol",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetCustomBgpAddresses() {
	_jsii_.InvokeVoid(
		v,
		"resetCustomBgpAddresses",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetDpdTimeoutSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetDpdTimeoutSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetEgressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetEgressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetExpressRouteCircuitId() {
	_jsii_.InvokeVoid(
		v,
		"resetExpressRouteCircuitId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetExpressRouteGatewayBypass() {
	_jsii_.InvokeVoid(
		v,
		"resetExpressRouteGatewayBypass",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetId() {
	_jsii_.InvokeVoid(
		v,
		"resetId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetIngressNatRuleIds() {
	_jsii_.InvokeVoid(
		v,
		"resetIngressNatRuleIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetIpsecPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetIpsecPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetLocalAzureIpAddressEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalAzureIpAddressEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetLocalNetworkGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetLocalNetworkGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetPeerVirtualNetworkGatewayId() {
	_jsii_.InvokeVoid(
		v,
		"resetPeerVirtualNetworkGatewayId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetPrivateLinkFastPathEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetPrivateLinkFastPathEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetRoutingWeight() {
	_jsii_.InvokeVoid(
		v,
		"resetRoutingWeight",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetSharedKey() {
	_jsii_.InvokeVoid(
		v,
		"resetSharedKey",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetTags() {
	_jsii_.InvokeVoid(
		v,
		"resetTags",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		v,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetTrafficSelectorPolicy() {
	_jsii_.InvokeVoid(
		v,
		"resetTrafficSelectorPolicy",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ResetUsePolicyBasedTrafficSelectors() {
	_jsii_.InvokeVoid(
		v,
		"resetUsePolicyBasedTrafficSelectors",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayConnection) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

