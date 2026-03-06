// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstanceactivedirectoryadministrator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/mssqlmanagedinstanceactivedirectoryadministrator/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.63.0/docs/resources/mssql_managed_instance_active_directory_administrator azurerm_mssql_managed_instance_active_directory_administrator}.
type MssqlManagedInstanceActiveDirectoryAdministrator interface {
	cdktn.TerraformResource
	AzureadAuthenticationOnly() interface{}
	SetAzureadAuthenticationOnly(val interface{})
	AzureadAuthenticationOnlyInput() interface{}
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
	LoginUsername() *string
	SetLoginUsername(val *string)
	LoginUsernameInput() *string
	ManagedInstanceId() *string
	SetManagedInstanceId(val *string)
	ManagedInstanceIdInput() *string
	// The tree node.
	Node() constructs.Node
	ObjectId() *string
	SetObjectId(val *string)
	ObjectIdInput() *string
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
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MssqlManagedInstanceActiveDirectoryAdministratorTimeoutsOutputReference
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
	PutTimeouts(value *MssqlManagedInstanceActiveDirectoryAdministratorTimeouts)
	ResetAzureadAuthenticationOnly()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for MssqlManagedInstanceActiveDirectoryAdministrator
type jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) AzureadAuthenticationOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureadAuthenticationOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) AzureadAuthenticationOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureadAuthenticationOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) LoginUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) LoginUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ManagedInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedInstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ManagedInstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedInstanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) Timeouts() MssqlManagedInstanceActiveDirectoryAdministratorTimeoutsOutputReference {
	var returns MssqlManagedInstanceActiveDirectoryAdministratorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.63.0/docs/resources/mssql_managed_instance_active_directory_administrator azurerm_mssql_managed_instance_active_directory_administrator} Resource.
func NewMssqlManagedInstanceActiveDirectoryAdministrator(scope constructs.Construct, id *string, config *MssqlManagedInstanceActiveDirectoryAdministratorConfig) MssqlManagedInstanceActiveDirectoryAdministrator {
	_init_.Initialize()

	if err := validateNewMssqlManagedInstanceActiveDirectoryAdministratorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.mssqlManagedInstanceActiveDirectoryAdministrator.MssqlManagedInstanceActiveDirectoryAdministrator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.63.0/docs/resources/mssql_managed_instance_active_directory_administrator azurerm_mssql_managed_instance_active_directory_administrator} Resource.
func NewMssqlManagedInstanceActiveDirectoryAdministrator_Override(m MssqlManagedInstanceActiveDirectoryAdministrator, scope constructs.Construct, id *string, config *MssqlManagedInstanceActiveDirectoryAdministratorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.mssqlManagedInstanceActiveDirectoryAdministrator.MssqlManagedInstanceActiveDirectoryAdministrator",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetAzureadAuthenticationOnly(val interface{}) {
	if err := j.validateSetAzureadAuthenticationOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azureadAuthenticationOnly",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetLoginUsername(val *string) {
	if err := j.validateSetLoginUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loginUsername",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetManagedInstanceId(val *string) {
	if err := j.validateSetManagedInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedInstanceId",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetObjectId(val *string) {
	if err := j.validateSetObjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectId",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

// Generates CDKTN code for importing a MssqlManagedInstanceActiveDirectoryAdministrator resource upon running "cdktn plan <stack-name>".
func MssqlManagedInstanceActiveDirectoryAdministrator_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateMssqlManagedInstanceActiveDirectoryAdministrator_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.mssqlManagedInstanceActiveDirectoryAdministrator.MssqlManagedInstanceActiveDirectoryAdministrator",
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
func MssqlManagedInstanceActiveDirectoryAdministrator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlManagedInstanceActiveDirectoryAdministrator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.mssqlManagedInstanceActiveDirectoryAdministrator.MssqlManagedInstanceActiveDirectoryAdministrator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlManagedInstanceActiveDirectoryAdministrator_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlManagedInstanceActiveDirectoryAdministrator_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.mssqlManagedInstanceActiveDirectoryAdministrator.MssqlManagedInstanceActiveDirectoryAdministrator",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlManagedInstanceActiveDirectoryAdministrator_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlManagedInstanceActiveDirectoryAdministrator_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.mssqlManagedInstanceActiveDirectoryAdministrator.MssqlManagedInstanceActiveDirectoryAdministrator",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MssqlManagedInstanceActiveDirectoryAdministrator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.mssqlManagedInstanceActiveDirectoryAdministrator.MssqlManagedInstanceActiveDirectoryAdministrator",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) PutTimeouts(value *MssqlManagedInstanceActiveDirectoryAdministratorTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ResetAzureadAuthenticationOnly() {
	_jsii_.InvokeVoid(
		m,
		"resetAzureadAuthenticationOnly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlManagedInstanceActiveDirectoryAdministrator) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

