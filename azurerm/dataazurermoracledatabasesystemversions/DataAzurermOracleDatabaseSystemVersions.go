// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazurermoracledatabasesystemversions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/dataazurermoracledatabasesystemversions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/data-sources/oracle_database_system_versions azurerm_oracle_database_system_versions}.
type DataAzurermOracleDatabaseSystemVersions interface {
	cdktn.TerraformDataSource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseSoftwareImageSupported() interface{}
	SetDatabaseSoftwareImageSupported(val interface{})
	DatabaseSoftwareImageSupportedInput() interface{}
	DatabaseSystemShape() *string
	SetDatabaseSystemShape(val *string)
	DatabaseSystemShapeInput() *string
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
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ShapeFamily() *string
	SetShapeFamily(val *string)
	ShapeFamilyInput() *string
	StorageManagement() *string
	SetStorageManagement(val *string)
	StorageManagementInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataAzurermOracleDatabaseSystemVersionsTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpgradeSupported() interface{}
	SetUpgradeSupported(val interface{})
	UpgradeSupportedInput() interface{}
	Versions() DataAzurermOracleDatabaseSystemVersionsVersionsList
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DataAzurermOracleDatabaseSystemVersionsTimeouts)
	ResetDatabaseSoftwareImageSupported()
	ResetDatabaseSystemShape()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetShapeFamily()
	ResetStorageManagement()
	ResetTimeouts()
	ResetUpgradeSupported()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAzurermOracleDatabaseSystemVersions
type jsiiProxy_DataAzurermOracleDatabaseSystemVersions struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) DatabaseSoftwareImageSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseSoftwareImageSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) DatabaseSoftwareImageSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseSoftwareImageSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) DatabaseSystemShape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSystemShape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) DatabaseSystemShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseSystemShapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ShapeFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ShapeFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shapeFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) StorageManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) StorageManagementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Timeouts() DataAzurermOracleDatabaseSystemVersionsTimeoutsOutputReference {
	var returns DataAzurermOracleDatabaseSystemVersionsTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) UpgradeSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) UpgradeSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upgradeSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) Versions() DataAzurermOracleDatabaseSystemVersionsVersionsList {
	var returns DataAzurermOracleDatabaseSystemVersionsVersionsList
	_jsii_.Get(
		j,
		"versions",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/data-sources/oracle_database_system_versions azurerm_oracle_database_system_versions} Data Source.
func NewDataAzurermOracleDatabaseSystemVersions(scope constructs.Construct, id *string, config *DataAzurermOracleDatabaseSystemVersionsConfig) DataAzurermOracleDatabaseSystemVersions {
	_init_.Initialize()

	if err := validateNewDataAzurermOracleDatabaseSystemVersionsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermOracleDatabaseSystemVersions{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermOracleDatabaseSystemVersions.DataAzurermOracleDatabaseSystemVersions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/data-sources/oracle_database_system_versions azurerm_oracle_database_system_versions} Data Source.
func NewDataAzurermOracleDatabaseSystemVersions_Override(d DataAzurermOracleDatabaseSystemVersions, scope constructs.Construct, id *string, config *DataAzurermOracleDatabaseSystemVersionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermOracleDatabaseSystemVersions.DataAzurermOracleDatabaseSystemVersions",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetDatabaseSoftwareImageSupported(val interface{}) {
	if err := j.validateSetDatabaseSoftwareImageSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseSoftwareImageSupported",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetDatabaseSystemShape(val *string) {
	if err := j.validateSetDatabaseSystemShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseSystemShape",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetShapeFamily(val *string) {
	if err := j.validateSetShapeFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shapeFamily",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetStorageManagement(val *string) {
	if err := j.validateSetStorageManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageManagement",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleDatabaseSystemVersions)SetUpgradeSupported(val interface{}) {
	if err := j.validateSetUpgradeSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeSupported",
		val,
	)
}

// Generates CDKTN code for importing a DataAzurermOracleDatabaseSystemVersions resource upon running "cdktn plan <stack-name>".
func DataAzurermOracleDatabaseSystemVersions_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermOracleDatabaseSystemVersions_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleDatabaseSystemVersions.DataAzurermOracleDatabaseSystemVersions",
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
func DataAzurermOracleDatabaseSystemVersions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleDatabaseSystemVersions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleDatabaseSystemVersions.DataAzurermOracleDatabaseSystemVersions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleDatabaseSystemVersions_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleDatabaseSystemVersions_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleDatabaseSystemVersions.DataAzurermOracleDatabaseSystemVersions",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleDatabaseSystemVersions_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleDatabaseSystemVersions_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleDatabaseSystemVersions.DataAzurermOracleDatabaseSystemVersions",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermOracleDatabaseSystemVersions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.dataAzurermOracleDatabaseSystemVersions.DataAzurermOracleDatabaseSystemVersions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) PutTimeouts(value *DataAzurermOracleDatabaseSystemVersionsTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetDatabaseSoftwareImageSupported() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseSoftwareImageSupported",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetDatabaseSystemShape() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseSystemShape",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetShapeFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetShapeFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetStorageManagement() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageManagement",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ResetUpgradeSupported() {
	_jsii_.InvokeVoid(
		d,
		"resetUpgradeSupported",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleDatabaseSystemVersions) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

