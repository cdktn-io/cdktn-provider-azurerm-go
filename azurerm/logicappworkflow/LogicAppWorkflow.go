// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package logicappworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/logicappworkflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.76.0/docs/resources/logic_app_workflow azurerm_logic_app_workflow}.
type LogicAppWorkflow interface {
	cdktn.TerraformResource
	AccessControl() LogicAppWorkflowAccessControlOutputReference
	AccessControlInput() *LogicAppWorkflowAccessControl
	AccessEndpoint() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectorEndpointIpAddresses() *[]*string
	ConnectorOutboundIpAddresses() *[]*string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
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
	Identity() LogicAppWorkflowIdentityOutputReference
	IdentityInput() *LogicAppWorkflowIdentity
	IdInput() *string
	IntegrationServiceEnvironmentId() *string
	SetIntegrationServiceEnvironmentId(val *string)
	IntegrationServiceEnvironmentIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LogicAppIntegrationAccountId() *string
	SetLogicAppIntegrationAccountId(val *string)
	LogicAppIntegrationAccountIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LogicAppWorkflowTimeoutsOutputReference
	TimeoutsInput() interface{}
	WorkflowEndpointIpAddresses() *[]*string
	WorkflowOutboundIpAddresses() *[]*string
	WorkflowParameters() *map[string]*string
	SetWorkflowParameters(val *map[string]*string)
	WorkflowParametersInput() *map[string]*string
	WorkflowSchema() *string
	SetWorkflowSchema(val *string)
	WorkflowSchemaInput() *string
	WorkflowVersion() *string
	SetWorkflowVersion(val *string)
	WorkflowVersionInput() *string
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
	PutAccessControl(value *LogicAppWorkflowAccessControl)
	PutIdentity(value *LogicAppWorkflowIdentity)
	PutTimeouts(value *LogicAppWorkflowTimeouts)
	ResetAccessControl()
	ResetEnabled()
	ResetId()
	ResetIdentity()
	ResetIntegrationServiceEnvironmentId()
	ResetLogicAppIntegrationAccountId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetTags()
	ResetTimeouts()
	ResetWorkflowParameters()
	ResetWorkflowSchema()
	ResetWorkflowVersion()
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

// The jsii proxy struct for LogicAppWorkflow
type jsiiProxy_LogicAppWorkflow struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_LogicAppWorkflow) AccessControl() LogicAppWorkflowAccessControlOutputReference {
	var returns LogicAppWorkflowAccessControlOutputReference
	_jsii_.Get(
		j,
		"accessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) AccessControlInput() *LogicAppWorkflowAccessControl {
	var returns *LogicAppWorkflowAccessControl
	_jsii_.Get(
		j,
		"accessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) AccessEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) ConnectorEndpointIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectorEndpointIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) ConnectorOutboundIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectorOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Identity() LogicAppWorkflowIdentityOutputReference {
	var returns LogicAppWorkflowIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) IdentityInput() *LogicAppWorkflowIdentity {
	var returns *LogicAppWorkflowIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) IntegrationServiceEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationServiceEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) IntegrationServiceEnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationServiceEnvironmentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) LogicAppIntegrationAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicAppIntegrationAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) LogicAppIntegrationAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicAppIntegrationAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) Timeouts() LogicAppWorkflowTimeoutsOutputReference {
	var returns LogicAppWorkflowTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowEndpointIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workflowEndpointIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowOutboundIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"workflowOutboundIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"workflowParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"workflowParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowSchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowSchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflow) WorkflowVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.76.0/docs/resources/logic_app_workflow azurerm_logic_app_workflow} Resource.
func NewLogicAppWorkflow(scope constructs.Construct, id *string, config *LogicAppWorkflowConfig) LogicAppWorkflow {
	_init_.Initialize()

	if err := validateNewLogicAppWorkflowParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppWorkflow{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflow",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.76.0/docs/resources/logic_app_workflow azurerm_logic_app_workflow} Resource.
func NewLogicAppWorkflow_Override(l LogicAppWorkflow, scope constructs.Construct, id *string, config *LogicAppWorkflowConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflow",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetIntegrationServiceEnvironmentId(val *string) {
	if err := j.validateSetIntegrationServiceEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationServiceEnvironmentId",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetLogicAppIntegrationAccountId(val *string) {
	if err := j.validateSetLogicAppIntegrationAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logicAppIntegrationAccountId",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetWorkflowParameters(val *map[string]*string) {
	if err := j.validateSetWorkflowParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflowParameters",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetWorkflowSchema(val *string) {
	if err := j.validateSetWorkflowSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflowSchema",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflow)SetWorkflowVersion(val *string) {
	if err := j.validateSetWorkflowVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflowVersion",
		val,
	)
}

// Generates CDKTN code for importing a LogicAppWorkflow resource upon running "cdktn plan <stack-name>".
func LogicAppWorkflow_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateLogicAppWorkflow_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflow",
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
func LogicAppWorkflow_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogicAppWorkflow_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflow",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogicAppWorkflow_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogicAppWorkflow_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflow",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogicAppWorkflow_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogicAppWorkflow_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflow",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogicAppWorkflow_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.logicAppWorkflow.LogicAppWorkflow",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) PutAccessControl(value *LogicAppWorkflowAccessControl) {
	if err := l.validatePutAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAccessControl",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) PutIdentity(value *LogicAppWorkflowIdentity) {
	if err := l.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) PutTimeouts(value *LogicAppWorkflowTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetAccessControl() {
	_jsii_.InvokeVoid(
		l,
		"resetAccessControl",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetIntegrationServiceEnvironmentId() {
	_jsii_.InvokeVoid(
		l,
		"resetIntegrationServiceEnvironmentId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetLogicAppIntegrationAccountId() {
	_jsii_.InvokeVoid(
		l,
		"resetLogicAppIntegrationAccountId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetParameters() {
	_jsii_.InvokeVoid(
		l,
		"resetParameters",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetWorkflowParameters() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkflowParameters",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetWorkflowSchema() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkflowSchema",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) ResetWorkflowVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkflowVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflow) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflow) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		l,
		"with",
		args,
		&returns,
	)

	return returns
}

