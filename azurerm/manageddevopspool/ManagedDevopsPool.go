// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/manageddevopspool/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool azurerm_managed_devops_pool}.
type ManagedDevopsPool interface {
	cdktn.TerraformResource
	AzureDevopsOrganization() ManagedDevopsPoolAzureDevopsOrganizationOutputReference
	AzureDevopsOrganizationInput() *ManagedDevopsPoolAzureDevopsOrganization
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
	DevCenterProjectId() *string
	SetDevCenterProjectId(val *string)
	DevCenterProjectIdInput() *string
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
	Identity() ManagedDevopsPoolIdentityOutputReference
	IdentityInput() *ManagedDevopsPoolIdentity
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaximumConcurrency() *float64
	SetMaximumConcurrency(val *float64)
	MaximumConcurrencyInput() *float64
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	StatefulAgent() ManagedDevopsPoolStatefulAgentOutputReference
	StatefulAgentInput() *ManagedDevopsPoolStatefulAgent
	StatelessAgent() ManagedDevopsPoolStatelessAgentOutputReference
	StatelessAgentInput() *ManagedDevopsPoolStatelessAgent
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ManagedDevopsPoolTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualMachineScaleSetFabric() ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference
	VirtualMachineScaleSetFabricInput() *ManagedDevopsPoolVirtualMachineScaleSetFabric
	WorkFolder() *string
	SetWorkFolder(val *string)
	WorkFolderInput() *string
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
	PutAzureDevopsOrganization(value *ManagedDevopsPoolAzureDevopsOrganization)
	PutIdentity(value *ManagedDevopsPoolIdentity)
	PutStatefulAgent(value *ManagedDevopsPoolStatefulAgent)
	PutStatelessAgent(value *ManagedDevopsPoolStatelessAgent)
	PutTimeouts(value *ManagedDevopsPoolTimeouts)
	PutVirtualMachineScaleSetFabric(value *ManagedDevopsPoolVirtualMachineScaleSetFabric)
	ResetId()
	ResetIdentity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStatefulAgent()
	ResetStatelessAgent()
	ResetTags()
	ResetTimeouts()
	ResetWorkFolder()
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

// The jsii proxy struct for ManagedDevopsPool
type jsiiProxy_ManagedDevopsPool struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ManagedDevopsPool) AzureDevopsOrganization() ManagedDevopsPoolAzureDevopsOrganizationOutputReference {
	var returns ManagedDevopsPoolAzureDevopsOrganizationOutputReference
	_jsii_.Get(
		j,
		"azureDevopsOrganization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) AzureDevopsOrganizationInput() *ManagedDevopsPoolAzureDevopsOrganization {
	var returns *ManagedDevopsPoolAzureDevopsOrganization
	_jsii_.Get(
		j,
		"azureDevopsOrganizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) DevCenterProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devCenterProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) DevCenterProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devCenterProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Identity() ManagedDevopsPoolIdentityOutputReference {
	var returns ManagedDevopsPoolIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) IdentityInput() *ManagedDevopsPoolIdentity {
	var returns *ManagedDevopsPoolIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) MaximumConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) MaximumConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) StatefulAgent() ManagedDevopsPoolStatefulAgentOutputReference {
	var returns ManagedDevopsPoolStatefulAgentOutputReference
	_jsii_.Get(
		j,
		"statefulAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) StatefulAgentInput() *ManagedDevopsPoolStatefulAgent {
	var returns *ManagedDevopsPoolStatefulAgent
	_jsii_.Get(
		j,
		"statefulAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) StatelessAgent() ManagedDevopsPoolStatelessAgentOutputReference {
	var returns ManagedDevopsPoolStatelessAgentOutputReference
	_jsii_.Get(
		j,
		"statelessAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) StatelessAgentInput() *ManagedDevopsPoolStatelessAgent {
	var returns *ManagedDevopsPoolStatelessAgent
	_jsii_.Get(
		j,
		"statelessAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) Timeouts() ManagedDevopsPoolTimeoutsOutputReference {
	var returns ManagedDevopsPoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) VirtualMachineScaleSetFabric() ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference {
	var returns ManagedDevopsPoolVirtualMachineScaleSetFabricOutputReference
	_jsii_.Get(
		j,
		"virtualMachineScaleSetFabric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) VirtualMachineScaleSetFabricInput() *ManagedDevopsPoolVirtualMachineScaleSetFabric {
	var returns *ManagedDevopsPoolVirtualMachineScaleSetFabric
	_jsii_.Get(
		j,
		"virtualMachineScaleSetFabricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) WorkFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedDevopsPool) WorkFolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workFolderInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool azurerm_managed_devops_pool} Resource.
func NewManagedDevopsPool(scope constructs.Construct, id *string, config *ManagedDevopsPoolConfig) ManagedDevopsPool {
	_init_.Initialize()

	if err := validateNewManagedDevopsPoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedDevopsPool{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPool",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/managed_devops_pool azurerm_managed_devops_pool} Resource.
func NewManagedDevopsPool_Override(m ManagedDevopsPool, scope constructs.Construct, id *string, config *ManagedDevopsPoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPool",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetDevCenterProjectId(val *string) {
	if err := j.validateSetDevCenterProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devCenterProjectId",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetMaximumConcurrency(val *float64) {
	if err := j.validateSetMaximumConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumConcurrency",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ManagedDevopsPool)SetWorkFolder(val *string) {
	if err := j.validateSetWorkFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workFolder",
		val,
	)
}

// Generates CDKTN code for importing a ManagedDevopsPool resource upon running "cdktn plan <stack-name>".
func ManagedDevopsPool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateManagedDevopsPool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPool",
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
func ManagedDevopsPool_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDevopsPool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPool",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDevopsPool_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDevopsPool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPool",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ManagedDevopsPool_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateManagedDevopsPool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPool",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ManagedDevopsPool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.managedDevopsPool.ManagedDevopsPool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) PutAzureDevopsOrganization(value *ManagedDevopsPoolAzureDevopsOrganization) {
	if err := m.validatePutAzureDevopsOrganizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzureDevopsOrganization",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) PutIdentity(value *ManagedDevopsPoolIdentity) {
	if err := m.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIdentity",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) PutStatefulAgent(value *ManagedDevopsPoolStatefulAgent) {
	if err := m.validatePutStatefulAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStatefulAgent",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) PutStatelessAgent(value *ManagedDevopsPoolStatelessAgent) {
	if err := m.validatePutStatelessAgentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStatelessAgent",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) PutTimeouts(value *ManagedDevopsPoolTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) PutVirtualMachineScaleSetFabric(value *ManagedDevopsPoolVirtualMachineScaleSetFabric) {
	if err := m.validatePutVirtualMachineScaleSetFabricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVirtualMachineScaleSetFabric",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetIdentity() {
	_jsii_.InvokeVoid(
		m,
		"resetIdentity",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetStatefulAgent() {
	_jsii_.InvokeVoid(
		m,
		"resetStatefulAgent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetStatelessAgent() {
	_jsii_.InvokeVoid(
		m,
		"resetStatelessAgent",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) ResetWorkFolder() {
	_jsii_.InvokeVoid(
		m,
		"resetWorkFolder",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedDevopsPool) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedDevopsPool) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		m,
		"with",
		args,
		&returns,
	)

	return returns
}

