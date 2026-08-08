// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/streamanalyticsjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/stream_analytics_job azurerm_stream_analytics_job}.
type StreamAnalyticsJob interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CompatibilityLevel() *string
	SetCompatibilityLevel(val *string)
	CompatibilityLevelInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentStoragePolicy() *string
	SetContentStoragePolicy(val *string)
	ContentStoragePolicyInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataLocale() *string
	SetDataLocale(val *string)
	DataLocaleInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventsLateArrivalMaxDelayInSeconds() *float64
	SetEventsLateArrivalMaxDelayInSeconds(val *float64)
	EventsLateArrivalMaxDelayInSecondsInput() *float64
	EventsOutOfOrderMaxDelayInSeconds() *float64
	SetEventsOutOfOrderMaxDelayInSeconds(val *float64)
	EventsOutOfOrderMaxDelayInSecondsInput() *float64
	EventsOutOfOrderPolicy() *string
	SetEventsOutOfOrderPolicy(val *string)
	EventsOutOfOrderPolicyInput() *string
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
	Identity() StreamAnalyticsJobIdentityOutputReference
	IdentityInput() *StreamAnalyticsJobIdentity
	IdInput() *string
	JobId() *string
	JobStorageAccount() StreamAnalyticsJobJobStorageAccountOutputReference
	JobStorageAccountInput() *StreamAnalyticsJobJobStorageAccount
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
	OutputErrorPolicy() *string
	SetOutputErrorPolicy(val *string)
	OutputErrorPolicyInput() *string
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
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	StreamAnalyticsClusterId() *string
	SetStreamAnalyticsClusterId(val *string)
	StreamAnalyticsClusterIdInput() *string
	StreamingUnits() *float64
	SetStreamingUnits(val *float64)
	StreamingUnitsInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StreamAnalyticsJobTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransformationQuery() *string
	SetTransformationQuery(val *string)
	TransformationQueryInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutIdentity(value *StreamAnalyticsJobIdentity)
	PutJobStorageAccount(value *StreamAnalyticsJobJobStorageAccount)
	PutTimeouts(value *StreamAnalyticsJobTimeouts)
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
	ResetCompatibilityLevel()
	ResetContentStoragePolicy()
	ResetDataLocale()
	ResetEventsLateArrivalMaxDelayInSeconds()
	ResetEventsOutOfOrderMaxDelayInSeconds()
	ResetEventsOutOfOrderPolicy()
	ResetId()
	ResetIdentity()
	ResetJobStorageAccount()
	ResetOutputErrorPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSkuName()
	ResetStreamAnalyticsClusterId()
	ResetStreamingUnits()
	ResetTags()
	ResetTimeouts()
	ResetType()
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

// The jsii proxy struct for StreamAnalyticsJob
type jsiiProxy_StreamAnalyticsJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_StreamAnalyticsJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) CompatibilityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) CompatibilityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compatibilityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) ContentStoragePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentStoragePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) ContentStoragePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentStoragePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) DataLocale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLocale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) DataLocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataLocaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) EventsLateArrivalMaxDelayInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsLateArrivalMaxDelayInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) EventsLateArrivalMaxDelayInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsLateArrivalMaxDelayInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) EventsOutOfOrderMaxDelayInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsOutOfOrderMaxDelayInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) EventsOutOfOrderMaxDelayInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"eventsOutOfOrderMaxDelayInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) EventsOutOfOrderPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventsOutOfOrderPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) EventsOutOfOrderPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventsOutOfOrderPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Identity() StreamAnalyticsJobIdentityOutputReference {
	var returns StreamAnalyticsJobIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) IdentityInput() *StreamAnalyticsJobIdentity {
	var returns *StreamAnalyticsJobIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) JobStorageAccount() StreamAnalyticsJobJobStorageAccountOutputReference {
	var returns StreamAnalyticsJobJobStorageAccountOutputReference
	_jsii_.Get(
		j,
		"jobStorageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) JobStorageAccountInput() *StreamAnalyticsJobJobStorageAccount {
	var returns *StreamAnalyticsJobJobStorageAccount
	_jsii_.Get(
		j,
		"jobStorageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) OutputErrorPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputErrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) OutputErrorPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputErrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) StreamAnalyticsClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsClusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) StreamAnalyticsClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamAnalyticsClusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) StreamingUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"streamingUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) StreamingUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"streamingUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Timeouts() StreamAnalyticsJobTimeoutsOutputReference {
	var returns StreamAnalyticsJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TransformationQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TransformationQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StreamAnalyticsJob) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/stream_analytics_job azurerm_stream_analytics_job} Resource.
func NewStreamAnalyticsJob(scope constructs.Construct, id *string, config *StreamAnalyticsJobConfig) StreamAnalyticsJob {
	_init_.Initialize()

	if err := validateNewStreamAnalyticsJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StreamAnalyticsJob{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.streamAnalyticsJob.StreamAnalyticsJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/stream_analytics_job azurerm_stream_analytics_job} Resource.
func NewStreamAnalyticsJob_Override(s StreamAnalyticsJob, scope constructs.Construct, id *string, config *StreamAnalyticsJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.streamAnalyticsJob.StreamAnalyticsJob",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetCompatibilityLevel(val *string) {
	if err := j.validateSetCompatibilityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compatibilityLevel",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetContentStoragePolicy(val *string) {
	if err := j.validateSetContentStoragePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentStoragePolicy",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetDataLocale(val *string) {
	if err := j.validateSetDataLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataLocale",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetEventsLateArrivalMaxDelayInSeconds(val *float64) {
	if err := j.validateSetEventsLateArrivalMaxDelayInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsLateArrivalMaxDelayInSeconds",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetEventsOutOfOrderMaxDelayInSeconds(val *float64) {
	if err := j.validateSetEventsOutOfOrderMaxDelayInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsOutOfOrderMaxDelayInSeconds",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetEventsOutOfOrderPolicy(val *string) {
	if err := j.validateSetEventsOutOfOrderPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventsOutOfOrderPolicy",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetOutputErrorPolicy(val *string) {
	if err := j.validateSetOutputErrorPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputErrorPolicy",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetStreamAnalyticsClusterId(val *string) {
	if err := j.validateSetStreamAnalyticsClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamAnalyticsClusterId",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetStreamingUnits(val *float64) {
	if err := j.validateSetStreamingUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingUnits",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetTransformationQuery(val *string) {
	if err := j.validateSetTransformationQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformationQuery",
		val,
	)
}

func (j *jsiiProxy_StreamAnalyticsJob)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

// Generates CDKTN code for importing a StreamAnalyticsJob resource upon running "cdktn plan <stack-name>".
func StreamAnalyticsJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateStreamAnalyticsJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.streamAnalyticsJob.StreamAnalyticsJob",
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
func StreamAnalyticsJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.streamAnalyticsJob.StreamAnalyticsJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamAnalyticsJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.streamAnalyticsJob.StreamAnalyticsJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StreamAnalyticsJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStreamAnalyticsJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.streamAnalyticsJob.StreamAnalyticsJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StreamAnalyticsJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.streamAnalyticsJob.StreamAnalyticsJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := s.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) PutIdentity(value *StreamAnalyticsJobIdentity) {
	if err := s.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) PutJobStorageAccount(value *StreamAnalyticsJobJobStorageAccount) {
	if err := s.validatePutJobStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putJobStorageAccount",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) PutTimeouts(value *StreamAnalyticsJobTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := s.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetCompatibilityLevel() {
	_jsii_.InvokeVoid(
		s,
		"resetCompatibilityLevel",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetContentStoragePolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetContentStoragePolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetDataLocale() {
	_jsii_.InvokeVoid(
		s,
		"resetDataLocale",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetEventsLateArrivalMaxDelayInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetEventsLateArrivalMaxDelayInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetEventsOutOfOrderMaxDelayInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetEventsOutOfOrderMaxDelayInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetEventsOutOfOrderPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetEventsOutOfOrderPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetIdentity() {
	_jsii_.InvokeVoid(
		s,
		"resetIdentity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetJobStorageAccount() {
	_jsii_.InvokeVoid(
		s,
		"resetJobStorageAccount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetOutputErrorPolicy() {
	_jsii_.InvokeVoid(
		s,
		"resetOutputErrorPolicy",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetSkuName() {
	_jsii_.InvokeVoid(
		s,
		"resetSkuName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetStreamAnalyticsClusterId() {
	_jsii_.InvokeVoid(
		s,
		"resetStreamAnalyticsClusterId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetStreamingUnits() {
	_jsii_.InvokeVoid(
		s,
		"resetStreamingUnits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) ResetType() {
	_jsii_.InvokeVoid(
		s,
		"resetType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StreamAnalyticsJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StreamAnalyticsJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

