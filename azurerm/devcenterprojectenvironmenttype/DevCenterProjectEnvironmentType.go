// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devcenterprojectenvironmenttype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/devcenterprojectenvironmenttype/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/dev_center_project_environment_type azurerm_dev_center_project_environment_type}.
type DevCenterProjectEnvironmentType interface {
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
	CreatorRoleAssignmentRoles() *[]*string
	SetCreatorRoleAssignmentRoles(val *[]*string)
	CreatorRoleAssignmentRolesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DeploymentTargetId() *string
	SetDeploymentTargetId(val *string)
	DeploymentTargetIdInput() *string
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
	Identity() DevCenterProjectEnvironmentTypeIdentityOutputReference
	IdentityInput() *DevCenterProjectEnvironmentTypeIdentity
	IdInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DevCenterProjectEnvironmentTypeTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserRoleAssignment() DevCenterProjectEnvironmentTypeUserRoleAssignmentList
	UserRoleAssignmentInput() interface{}
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
	PutIdentity(value *DevCenterProjectEnvironmentTypeIdentity)
	PutTimeouts(value *DevCenterProjectEnvironmentTypeTimeouts)
	PutUserRoleAssignment(value interface{})
	ResetCreatorRoleAssignmentRoles()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	ResetUserRoleAssignment()
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

// The jsii proxy struct for DevCenterProjectEnvironmentType
type jsiiProxy_DevCenterProjectEnvironmentType struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) CreatorRoleAssignmentRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creatorRoleAssignmentRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) CreatorRoleAssignmentRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creatorRoleAssignmentRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) DeploymentTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) DeploymentTargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deploymentTargetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) DevCenterProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devCenterProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) DevCenterProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"devCenterProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Identity() DevCenterProjectEnvironmentTypeIdentityOutputReference {
	var returns DevCenterProjectEnvironmentTypeIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) IdentityInput() *DevCenterProjectEnvironmentTypeIdentity {
	var returns *DevCenterProjectEnvironmentTypeIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) Timeouts() DevCenterProjectEnvironmentTypeTimeoutsOutputReference {
	var returns DevCenterProjectEnvironmentTypeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) UserRoleAssignment() DevCenterProjectEnvironmentTypeUserRoleAssignmentList {
	var returns DevCenterProjectEnvironmentTypeUserRoleAssignmentList
	_jsii_.Get(
		j,
		"userRoleAssignment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType) UserRoleAssignmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userRoleAssignmentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/dev_center_project_environment_type azurerm_dev_center_project_environment_type} Resource.
func NewDevCenterProjectEnvironmentType(scope constructs.Construct, id *string, config *DevCenterProjectEnvironmentTypeConfig) DevCenterProjectEnvironmentType {
	_init_.Initialize()

	if err := validateNewDevCenterProjectEnvironmentTypeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DevCenterProjectEnvironmentType{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.devCenterProjectEnvironmentType.DevCenterProjectEnvironmentType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.80.0/docs/resources/dev_center_project_environment_type azurerm_dev_center_project_environment_type} Resource.
func NewDevCenterProjectEnvironmentType_Override(d DevCenterProjectEnvironmentType, scope constructs.Construct, id *string, config *DevCenterProjectEnvironmentTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.devCenterProjectEnvironmentType.DevCenterProjectEnvironmentType",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetCreatorRoleAssignmentRoles(val *[]*string) {
	if err := j.validateSetCreatorRoleAssignmentRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creatorRoleAssignmentRoles",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetDeploymentTargetId(val *string) {
	if err := j.validateSetDeploymentTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentTargetId",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetDevCenterProjectId(val *string) {
	if err := j.validateSetDevCenterProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"devCenterProjectId",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DevCenterProjectEnvironmentType)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a DevCenterProjectEnvironmentType resource upon running "cdktn plan <stack-name>".
func DevCenterProjectEnvironmentType_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDevCenterProjectEnvironmentType_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.devCenterProjectEnvironmentType.DevCenterProjectEnvironmentType",
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
func DevCenterProjectEnvironmentType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevCenterProjectEnvironmentType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.devCenterProjectEnvironmentType.DevCenterProjectEnvironmentType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevCenterProjectEnvironmentType_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevCenterProjectEnvironmentType_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.devCenterProjectEnvironmentType.DevCenterProjectEnvironmentType",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DevCenterProjectEnvironmentType_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDevCenterProjectEnvironmentType_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.devCenterProjectEnvironmentType.DevCenterProjectEnvironmentType",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DevCenterProjectEnvironmentType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.devCenterProjectEnvironmentType.DevCenterProjectEnvironmentType",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) PutIdentity(value *DevCenterProjectEnvironmentTypeIdentity) {
	if err := d.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIdentity",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) PutTimeouts(value *DevCenterProjectEnvironmentTypeTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) PutUserRoleAssignment(value interface{}) {
	if err := d.validatePutUserRoleAssignmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserRoleAssignment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ResetCreatorRoleAssignmentRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetCreatorRoleAssignmentRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ResetUserRoleAssignment() {
	_jsii_.InvokeVoid(
		d,
		"resetUserRoleAssignment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DevCenterProjectEnvironmentType) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

