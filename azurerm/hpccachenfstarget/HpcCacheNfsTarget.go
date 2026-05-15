// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hpccachenfstarget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/hpccachenfstarget/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/hpc_cache_nfs_target azurerm_hpc_cache_nfs_target}.
type HpcCacheNfsTarget interface {
	cdktn.TerraformResource
	CacheName() *string
	SetCacheName(val *string)
	CacheNameInput() *string
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
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamespaceJunction() HpcCacheNfsTargetNamespaceJunctionList
	NamespaceJunctionInput() interface{}
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	TargetHostName() *string
	SetTargetHostName(val *string)
	TargetHostNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() HpcCacheNfsTargetTimeoutsOutputReference
	TimeoutsInput() interface{}
	UsageModel() *string
	SetUsageModel(val *string)
	UsageModelInput() *string
	VerificationTimerInSeconds() *float64
	SetVerificationTimerInSeconds(val *float64)
	VerificationTimerInSecondsInput() *float64
	WriteBackTimerInSeconds() *float64
	SetWriteBackTimerInSeconds(val *float64)
	WriteBackTimerInSecondsInput() *float64
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
	PutNamespaceJunction(value interface{})
	PutTimeouts(value *HpcCacheNfsTargetTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetVerificationTimerInSeconds()
	ResetWriteBackTimerInSeconds()
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

// The jsii proxy struct for HpcCacheNfsTarget
type jsiiProxy_HpcCacheNfsTarget struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_HpcCacheNfsTarget) CacheName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) CacheNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) NamespaceJunction() HpcCacheNfsTargetNamespaceJunctionList {
	var returns HpcCacheNfsTargetNamespaceJunctionList
	_jsii_.Get(
		j,
		"namespaceJunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) NamespaceJunctionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namespaceJunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) TargetHostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetHostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) TargetHostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetHostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) Timeouts() HpcCacheNfsTargetTimeoutsOutputReference {
	var returns HpcCacheNfsTargetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) UsageModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) UsageModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) VerificationTimerInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"verificationTimerInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) VerificationTimerInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"verificationTimerInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) WriteBackTimerInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeBackTimerInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheNfsTarget) WriteBackTimerInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeBackTimerInSecondsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/hpc_cache_nfs_target azurerm_hpc_cache_nfs_target} Resource.
func NewHpcCacheNfsTarget(scope constructs.Construct, id *string, config *HpcCacheNfsTargetConfig) HpcCacheNfsTarget {
	_init_.Initialize()

	if err := validateNewHpcCacheNfsTargetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HpcCacheNfsTarget{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.hpcCacheNfsTarget.HpcCacheNfsTarget",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/hpc_cache_nfs_target azurerm_hpc_cache_nfs_target} Resource.
func NewHpcCacheNfsTarget_Override(h HpcCacheNfsTarget, scope constructs.Construct, id *string, config *HpcCacheNfsTargetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.hpcCacheNfsTarget.HpcCacheNfsTarget",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetCacheName(val *string) {
	if err := j.validateSetCacheNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetTargetHostName(val *string) {
	if err := j.validateSetTargetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetHostName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetUsageModel(val *string) {
	if err := j.validateSetUsageModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageModel",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetVerificationTimerInSeconds(val *float64) {
	if err := j.validateSetVerificationTimerInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verificationTimerInSeconds",
		val,
	)
}

func (j *jsiiProxy_HpcCacheNfsTarget)SetWriteBackTimerInSeconds(val *float64) {
	if err := j.validateSetWriteBackTimerInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeBackTimerInSeconds",
		val,
	)
}

// Generates CDKTN code for importing a HpcCacheNfsTarget resource upon running "cdktn plan <stack-name>".
func HpcCacheNfsTarget_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateHpcCacheNfsTarget_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.hpcCacheNfsTarget.HpcCacheNfsTarget",
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
func HpcCacheNfsTarget_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHpcCacheNfsTarget_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.hpcCacheNfsTarget.HpcCacheNfsTarget",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HpcCacheNfsTarget_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHpcCacheNfsTarget_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.hpcCacheNfsTarget.HpcCacheNfsTarget",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HpcCacheNfsTarget_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHpcCacheNfsTarget_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.hpcCacheNfsTarget.HpcCacheNfsTarget",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HpcCacheNfsTarget_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.hpcCacheNfsTarget.HpcCacheNfsTarget",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) PutNamespaceJunction(value interface{}) {
	if err := h.validatePutNamespaceJunctionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putNamespaceJunction",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) PutTimeouts(value *HpcCacheNfsTargetTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) ResetVerificationTimerInSeconds() {
	_jsii_.InvokeVoid(
		h,
		"resetVerificationTimerInSeconds",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) ResetWriteBackTimerInSeconds() {
	_jsii_.InvokeVoid(
		h,
		"resetWriteBackTimerInSeconds",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheNfsTarget) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheNfsTarget) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		h,
		"with",
		args,
		&returns,
	)

	return returns
}

