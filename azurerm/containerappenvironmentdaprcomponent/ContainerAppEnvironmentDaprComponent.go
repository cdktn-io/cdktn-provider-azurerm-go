// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentdaprcomponent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/containerappenvironmentdaprcomponent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/container_app_environment_dapr_component azurerm_container_app_environment_dapr_component}.
type ContainerAppEnvironmentDaprComponent interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ComponentType() *string
	SetComponentType(val *string)
	ComponentTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerAppEnvironmentId() *string
	SetContainerAppEnvironmentId(val *string)
	ContainerAppEnvironmentIdInput() *string
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
	IgnoreErrors() interface{}
	SetIgnoreErrors(val interface{})
	IgnoreErrorsInput() interface{}
	InitTimeout() *string
	SetInitTimeout(val *string)
	InitTimeoutInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Metadata() ContainerAppEnvironmentDaprComponentMetadataList
	MetadataInput() interface{}
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
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
	Secret() ContainerAppEnvironmentDaprComponentSecretList
	SecretInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerAppEnvironmentDaprComponentTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutMetadata(value interface{})
	PutSecret(value interface{})
	PutTimeouts(value *ContainerAppEnvironmentDaprComponentTimeouts)
	ResetId()
	ResetIgnoreErrors()
	ResetInitTimeout()
	ResetMetadata()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetScopes()
	ResetSecret()
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

// The jsii proxy struct for ContainerAppEnvironmentDaprComponent
type jsiiProxy_ContainerAppEnvironmentDaprComponent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) ComponentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) ComponentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) ContainerAppEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAppEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) ContainerAppEnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerAppEnvironmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) IgnoreErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) IgnoreErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) InitTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) InitTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Metadata() ContainerAppEnvironmentDaprComponentMetadataList {
	var returns ContainerAppEnvironmentDaprComponentMetadataList
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) MetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Secret() ContainerAppEnvironmentDaprComponentSecretList {
	var returns ContainerAppEnvironmentDaprComponentSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Timeouts() ContainerAppEnvironmentDaprComponentTimeoutsOutputReference {
	var returns ContainerAppEnvironmentDaprComponentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/container_app_environment_dapr_component azurerm_container_app_environment_dapr_component} Resource.
func NewContainerAppEnvironmentDaprComponent(scope constructs.Construct, id *string, config *ContainerAppEnvironmentDaprComponentConfig) ContainerAppEnvironmentDaprComponent {
	_init_.Initialize()

	if err := validateNewContainerAppEnvironmentDaprComponentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAppEnvironmentDaprComponent{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.containerAppEnvironmentDaprComponent.ContainerAppEnvironmentDaprComponent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/container_app_environment_dapr_component azurerm_container_app_environment_dapr_component} Resource.
func NewContainerAppEnvironmentDaprComponent_Override(c ContainerAppEnvironmentDaprComponent, scope constructs.Construct, id *string, config *ContainerAppEnvironmentDaprComponentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.containerAppEnvironmentDaprComponent.ContainerAppEnvironmentDaprComponent",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetComponentType(val *string) {
	if err := j.validateSetComponentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentType",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetContainerAppEnvironmentId(val *string) {
	if err := j.validateSetContainerAppEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerAppEnvironmentId",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetIgnoreErrors(val interface{}) {
	if err := j.validateSetIgnoreErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreErrors",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetInitTimeout(val *string) {
	if err := j.validateSetInitTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initTimeout",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_ContainerAppEnvironmentDaprComponent)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a ContainerAppEnvironmentDaprComponent resource upon running "cdktn plan <stack-name>".
func ContainerAppEnvironmentDaprComponent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateContainerAppEnvironmentDaprComponent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.containerAppEnvironmentDaprComponent.ContainerAppEnvironmentDaprComponent",
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
func ContainerAppEnvironmentDaprComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppEnvironmentDaprComponent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.containerAppEnvironmentDaprComponent.ContainerAppEnvironmentDaprComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAppEnvironmentDaprComponent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppEnvironmentDaprComponent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.containerAppEnvironmentDaprComponent.ContainerAppEnvironmentDaprComponent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAppEnvironmentDaprComponent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAppEnvironmentDaprComponent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.containerAppEnvironmentDaprComponent.ContainerAppEnvironmentDaprComponent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerAppEnvironmentDaprComponent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.containerAppEnvironmentDaprComponent.ContainerAppEnvironmentDaprComponent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) PutMetadata(value interface{}) {
	if err := c.validatePutMetadataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetadata",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) PutSecret(value interface{}) {
	if err := c.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecret",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) PutTimeouts(value *ContainerAppEnvironmentDaprComponentTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetIgnoreErrors() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoreErrors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetInitTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetInitTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetSecret() {
	_jsii_.InvokeVoid(
		c,
		"resetSecret",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAppEnvironmentDaprComponent) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

