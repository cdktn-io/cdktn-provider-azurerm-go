// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kustocosmosdbdataconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/kustocosmosdbdataconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.66.0/docs/resources/kusto_cosmosdb_data_connection azurerm_kusto_cosmosdb_data_connection}.
type KustoCosmosdbDataConnection interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CosmosdbContainerId() *string
	SetCosmosdbContainerId(val *string)
	CosmosdbContainerIdInput() *string
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
	KustoDatabaseId() *string
	SetKustoDatabaseId(val *string)
	KustoDatabaseIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedIdentityId() *string
	SetManagedIdentityId(val *string)
	ManagedIdentityIdInput() *string
	MappingRuleName() *string
	SetMappingRuleName(val *string)
	MappingRuleNameInput() *string
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
	RetrievalStartDate() *string
	SetRetrievalStartDate(val *string)
	RetrievalStartDateInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KustoCosmosdbDataConnectionTimeoutsOutputReference
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
	PutTimeouts(value *KustoCosmosdbDataConnectionTimeouts)
	ResetId()
	ResetMappingRuleName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetrievalStartDate()
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

// The jsii proxy struct for KustoCosmosdbDataConnection
type jsiiProxy_KustoCosmosdbDataConnection struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) CosmosdbContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbContainerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) CosmosdbContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cosmosdbContainerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) KustoDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) KustoDatabaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kustoDatabaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) ManagedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) ManagedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) MappingRuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) MappingRuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mappingRuleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) RetrievalStartDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalStartDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) RetrievalStartDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalStartDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) Timeouts() KustoCosmosdbDataConnectionTimeoutsOutputReference {
	var returns KustoCosmosdbDataConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KustoCosmosdbDataConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.66.0/docs/resources/kusto_cosmosdb_data_connection azurerm_kusto_cosmosdb_data_connection} Resource.
func NewKustoCosmosdbDataConnection(scope constructs.Construct, id *string, config *KustoCosmosdbDataConnectionConfig) KustoCosmosdbDataConnection {
	_init_.Initialize()

	if err := validateNewKustoCosmosdbDataConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KustoCosmosdbDataConnection{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.kustoCosmosdbDataConnection.KustoCosmosdbDataConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.66.0/docs/resources/kusto_cosmosdb_data_connection azurerm_kusto_cosmosdb_data_connection} Resource.
func NewKustoCosmosdbDataConnection_Override(k KustoCosmosdbDataConnection, scope constructs.Construct, id *string, config *KustoCosmosdbDataConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.kustoCosmosdbDataConnection.KustoCosmosdbDataConnection",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetCosmosdbContainerId(val *string) {
	if err := j.validateSetCosmosdbContainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cosmosdbContainerId",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetKustoDatabaseId(val *string) {
	if err := j.validateSetKustoDatabaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kustoDatabaseId",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetManagedIdentityId(val *string) {
	if err := j.validateSetManagedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedIdentityId",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetMappingRuleName(val *string) {
	if err := j.validateSetMappingRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mappingRuleName",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetRetrievalStartDate(val *string) {
	if err := j.validateSetRetrievalStartDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrievalStartDate",
		val,
	)
}

func (j *jsiiProxy_KustoCosmosdbDataConnection)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Generates CDKTN code for importing a KustoCosmosdbDataConnection resource upon running "cdktn plan <stack-name>".
func KustoCosmosdbDataConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKustoCosmosdbDataConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kustoCosmosdbDataConnection.KustoCosmosdbDataConnection",
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
func KustoCosmosdbDataConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoCosmosdbDataConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kustoCosmosdbDataConnection.KustoCosmosdbDataConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoCosmosdbDataConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoCosmosdbDataConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kustoCosmosdbDataConnection.KustoCosmosdbDataConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KustoCosmosdbDataConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKustoCosmosdbDataConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kustoCosmosdbDataConnection.KustoCosmosdbDataConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KustoCosmosdbDataConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.kustoCosmosdbDataConnection.KustoCosmosdbDataConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) PutTimeouts(value *KustoCosmosdbDataConnectionTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ResetMappingRuleName() {
	_jsii_.InvokeVoid(
		k,
		"resetMappingRuleName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ResetRetrievalStartDate() {
	_jsii_.InvokeVoid(
		k,
		"resetRetrievalStartDate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KustoCosmosdbDataConnection) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		k,
		"with",
		args,
		&returns,
	)

	return returns
}

