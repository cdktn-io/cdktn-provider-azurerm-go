// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthcaremedtechservicefhirdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/healthcaremedtechservicefhirdestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/healthcare_medtech_service_fhir_destination azurerm_healthcare_medtech_service_fhir_destination}.
type HealthcareMedtechServiceFhirDestination interface {
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
	DestinationFhirMappingJson() *string
	SetDestinationFhirMappingJson(val *string)
	DestinationFhirMappingJsonInput() *string
	DestinationFhirServiceId() *string
	SetDestinationFhirServiceId(val *string)
	DestinationFhirServiceIdInput() *string
	DestinationIdentityResolutionType() *string
	SetDestinationIdentityResolutionType(val *string)
	DestinationIdentityResolutionTypeInput() *string
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
	MedtechServiceId() *string
	SetMedtechServiceId(val *string)
	MedtechServiceIdInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() HealthcareMedtechServiceFhirDestinationTimeoutsOutputReference
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
	PutTimeouts(value *HealthcareMedtechServiceFhirDestinationTimeouts)
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

// The jsii proxy struct for HealthcareMedtechServiceFhirDestination
type jsiiProxy_HealthcareMedtechServiceFhirDestination struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) DestinationFhirMappingJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFhirMappingJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) DestinationFhirMappingJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFhirMappingJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) DestinationFhirServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFhirServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) DestinationFhirServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFhirServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) DestinationIdentityResolutionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIdentityResolutionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) DestinationIdentityResolutionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationIdentityResolutionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) MedtechServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"medtechServiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) MedtechServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"medtechServiceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) Timeouts() HealthcareMedtechServiceFhirDestinationTimeoutsOutputReference {
	var returns HealthcareMedtechServiceFhirDestinationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/healthcare_medtech_service_fhir_destination azurerm_healthcare_medtech_service_fhir_destination} Resource.
func NewHealthcareMedtechServiceFhirDestination(scope constructs.Construct, id *string, config *HealthcareMedtechServiceFhirDestinationConfig) HealthcareMedtechServiceFhirDestination {
	_init_.Initialize()

	if err := validateNewHealthcareMedtechServiceFhirDestinationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthcareMedtechServiceFhirDestination{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.healthcareMedtechServiceFhirDestination.HealthcareMedtechServiceFhirDestination",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.78.0/docs/resources/healthcare_medtech_service_fhir_destination azurerm_healthcare_medtech_service_fhir_destination} Resource.
func NewHealthcareMedtechServiceFhirDestination_Override(h HealthcareMedtechServiceFhirDestination, scope constructs.Construct, id *string, config *HealthcareMedtechServiceFhirDestinationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.healthcareMedtechServiceFhirDestination.HealthcareMedtechServiceFhirDestination",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetDestinationFhirMappingJson(val *string) {
	if err := j.validateSetDestinationFhirMappingJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFhirMappingJson",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetDestinationFhirServiceId(val *string) {
	if err := j.validateSetDestinationFhirServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFhirServiceId",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetDestinationIdentityResolutionType(val *string) {
	if err := j.validateSetDestinationIdentityResolutionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationIdentityResolutionType",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetMedtechServiceId(val *string) {
	if err := j.validateSetMedtechServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"medtechServiceId",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HealthcareMedtechServiceFhirDestination)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a HealthcareMedtechServiceFhirDestination resource upon running "cdktn plan <stack-name>".
func HealthcareMedtechServiceFhirDestination_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateHealthcareMedtechServiceFhirDestination_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.healthcareMedtechServiceFhirDestination.HealthcareMedtechServiceFhirDestination",
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
func HealthcareMedtechServiceFhirDestination_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareMedtechServiceFhirDestination_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.healthcareMedtechServiceFhirDestination.HealthcareMedtechServiceFhirDestination",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareMedtechServiceFhirDestination_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareMedtechServiceFhirDestination_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.healthcareMedtechServiceFhirDestination.HealthcareMedtechServiceFhirDestination",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HealthcareMedtechServiceFhirDestination_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHealthcareMedtechServiceFhirDestination_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.healthcareMedtechServiceFhirDestination.HealthcareMedtechServiceFhirDestination",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HealthcareMedtechServiceFhirDestination_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.healthcareMedtechServiceFhirDestination.HealthcareMedtechServiceFhirDestination",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) PutTimeouts(value *HealthcareMedtechServiceFhirDestinationTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthcareMedtechServiceFhirDestination) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

