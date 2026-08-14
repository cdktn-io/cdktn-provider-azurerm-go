// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/eventgrideventsubscription/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/eventgrid_event_subscription azurerm_eventgrid_event_subscription}.
type EventgridEventSubscription interface {
	cdktn.TerraformResource
	AdvancedFilter() EventgridEventSubscriptionAdvancedFilterOutputReference
	AdvancedFilteringOnArraysEnabled() interface{}
	SetAdvancedFilteringOnArraysEnabled(val interface{})
	AdvancedFilteringOnArraysEnabledInput() interface{}
	AdvancedFilterInput() *EventgridEventSubscriptionAdvancedFilter
	AzureFunctionEndpoint() EventgridEventSubscriptionAzureFunctionEndpointOutputReference
	AzureFunctionEndpointInput() *EventgridEventSubscriptionAzureFunctionEndpoint
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
	DeadLetterIdentity() EventgridEventSubscriptionDeadLetterIdentityOutputReference
	DeadLetterIdentityInput() *EventgridEventSubscriptionDeadLetterIdentity
	DeliveryIdentity() EventgridEventSubscriptionDeliveryIdentityOutputReference
	DeliveryIdentityInput() *EventgridEventSubscriptionDeliveryIdentity
	DeliveryProperty() EventgridEventSubscriptionDeliveryPropertyList
	DeliveryPropertyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EventDeliverySchema() *string
	SetEventDeliverySchema(val *string)
	EventDeliverySchemaInput() *string
	EventhubId() *string
	SetEventhubId(val *string)
	EventhubIdInput() *string
	ExpirationTimeUtc() *string
	SetExpirationTimeUtc(val *string)
	ExpirationTimeUtcInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HybridConnectionId() *string
	SetHybridConnectionId(val *string)
	HybridConnectionIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IncludedEventTypes() *[]*string
	SetIncludedEventTypes(val *[]*string)
	IncludedEventTypesInput() *[]*string
	Labels() *[]*string
	SetLabels(val *[]*string)
	LabelsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
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
	RetryPolicy() EventgridEventSubscriptionRetryPolicyOutputReference
	RetryPolicyInput() *EventgridEventSubscriptionRetryPolicy
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	ServiceBusQueueId() *string
	SetServiceBusQueueId(val *string)
	ServiceBusQueueIdInput() *string
	ServiceBusTopicId() *string
	SetServiceBusTopicId(val *string)
	ServiceBusTopicIdInput() *string
	StorageBlobDeadLetterDestination() EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
	StorageBlobDeadLetterDestinationInput() *EventgridEventSubscriptionStorageBlobDeadLetterDestination
	StorageQueueEndpoint() EventgridEventSubscriptionStorageQueueEndpointOutputReference
	StorageQueueEndpointInput() *EventgridEventSubscriptionStorageQueueEndpoint
	SubjectFilter() EventgridEventSubscriptionSubjectFilterOutputReference
	SubjectFilterInput() *EventgridEventSubscriptionSubjectFilter
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EventgridEventSubscriptionTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebhookEndpoint() EventgridEventSubscriptionWebhookEndpointOutputReference
	WebhookEndpointInput() *EventgridEventSubscriptionWebhookEndpoint
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
	PutAdvancedFilter(value *EventgridEventSubscriptionAdvancedFilter)
	PutAzureFunctionEndpoint(value *EventgridEventSubscriptionAzureFunctionEndpoint)
	PutDeadLetterIdentity(value *EventgridEventSubscriptionDeadLetterIdentity)
	PutDeliveryIdentity(value *EventgridEventSubscriptionDeliveryIdentity)
	PutDeliveryProperty(value interface{})
	PutRetryPolicy(value *EventgridEventSubscriptionRetryPolicy)
	PutStorageBlobDeadLetterDestination(value *EventgridEventSubscriptionStorageBlobDeadLetterDestination)
	PutStorageQueueEndpoint(value *EventgridEventSubscriptionStorageQueueEndpoint)
	PutSubjectFilter(value *EventgridEventSubscriptionSubjectFilter)
	PutTimeouts(value *EventgridEventSubscriptionTimeouts)
	PutWebhookEndpoint(value *EventgridEventSubscriptionWebhookEndpoint)
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
	ResetAdvancedFilter()
	ResetAdvancedFilteringOnArraysEnabled()
	ResetAzureFunctionEndpoint()
	ResetDeadLetterIdentity()
	ResetDeliveryIdentity()
	ResetDeliveryProperty()
	ResetEventDeliverySchema()
	ResetEventhubId()
	ResetExpirationTimeUtc()
	ResetHybridConnectionId()
	ResetId()
	ResetIncludedEventTypes()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetryPolicy()
	ResetServiceBusQueueId()
	ResetServiceBusTopicId()
	ResetStorageBlobDeadLetterDestination()
	ResetStorageQueueEndpoint()
	ResetSubjectFilter()
	ResetTimeouts()
	ResetWebhookEndpoint()
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

// The jsii proxy struct for EventgridEventSubscription
type jsiiProxy_EventgridEventSubscription struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilter() EventgridEventSubscriptionAdvancedFilterOutputReference {
	var returns EventgridEventSubscriptionAdvancedFilterOutputReference
	_jsii_.Get(
		j,
		"advancedFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilteringOnArraysEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedFilteringOnArraysEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilteringOnArraysEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedFilteringOnArraysEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AdvancedFilterInput() *EventgridEventSubscriptionAdvancedFilter {
	var returns *EventgridEventSubscriptionAdvancedFilter
	_jsii_.Get(
		j,
		"advancedFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AzureFunctionEndpoint() EventgridEventSubscriptionAzureFunctionEndpointOutputReference {
	var returns EventgridEventSubscriptionAzureFunctionEndpointOutputReference
	_jsii_.Get(
		j,
		"azureFunctionEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) AzureFunctionEndpointInput() *EventgridEventSubscriptionAzureFunctionEndpoint {
	var returns *EventgridEventSubscriptionAzureFunctionEndpoint
	_jsii_.Get(
		j,
		"azureFunctionEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeadLetterIdentity() EventgridEventSubscriptionDeadLetterIdentityOutputReference {
	var returns EventgridEventSubscriptionDeadLetterIdentityOutputReference
	_jsii_.Get(
		j,
		"deadLetterIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeadLetterIdentityInput() *EventgridEventSubscriptionDeadLetterIdentity {
	var returns *EventgridEventSubscriptionDeadLetterIdentity
	_jsii_.Get(
		j,
		"deadLetterIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryIdentity() EventgridEventSubscriptionDeliveryIdentityOutputReference {
	var returns EventgridEventSubscriptionDeliveryIdentityOutputReference
	_jsii_.Get(
		j,
		"deliveryIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryIdentityInput() *EventgridEventSubscriptionDeliveryIdentity {
	var returns *EventgridEventSubscriptionDeliveryIdentity
	_jsii_.Get(
		j,
		"deliveryIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryProperty() EventgridEventSubscriptionDeliveryPropertyList {
	var returns EventgridEventSubscriptionDeliveryPropertyList
	_jsii_.Get(
		j,
		"deliveryProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DeliveryPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventDeliverySchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliverySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventDeliverySchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliverySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventhubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) EventhubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventhubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ExpirationTimeUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ExpirationTimeUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationTimeUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) HybridConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) HybridConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridConnectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) IncludedEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) IncludedEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) LabelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) RetryPolicy() EventgridEventSubscriptionRetryPolicyOutputReference {
	var returns EventgridEventSubscriptionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) RetryPolicyInput() *EventgridEventSubscriptionRetryPolicy {
	var returns *EventgridEventSubscriptionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusQueueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusQueueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusQueueIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusQueueIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusTopicId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusTopicId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) ServiceBusTopicIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceBusTopicIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageBlobDeadLetterDestination() EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference {
	var returns EventgridEventSubscriptionStorageBlobDeadLetterDestinationOutputReference
	_jsii_.Get(
		j,
		"storageBlobDeadLetterDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageBlobDeadLetterDestinationInput() *EventgridEventSubscriptionStorageBlobDeadLetterDestination {
	var returns *EventgridEventSubscriptionStorageBlobDeadLetterDestination
	_jsii_.Get(
		j,
		"storageBlobDeadLetterDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageQueueEndpoint() EventgridEventSubscriptionStorageQueueEndpointOutputReference {
	var returns EventgridEventSubscriptionStorageQueueEndpointOutputReference
	_jsii_.Get(
		j,
		"storageQueueEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) StorageQueueEndpointInput() *EventgridEventSubscriptionStorageQueueEndpoint {
	var returns *EventgridEventSubscriptionStorageQueueEndpoint
	_jsii_.Get(
		j,
		"storageQueueEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) SubjectFilter() EventgridEventSubscriptionSubjectFilterOutputReference {
	var returns EventgridEventSubscriptionSubjectFilterOutputReference
	_jsii_.Get(
		j,
		"subjectFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) SubjectFilterInput() *EventgridEventSubscriptionSubjectFilter {
	var returns *EventgridEventSubscriptionSubjectFilter
	_jsii_.Get(
		j,
		"subjectFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) Timeouts() EventgridEventSubscriptionTimeoutsOutputReference {
	var returns EventgridEventSubscriptionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) WebhookEndpoint() EventgridEventSubscriptionWebhookEndpointOutputReference {
	var returns EventgridEventSubscriptionWebhookEndpointOutputReference
	_jsii_.Get(
		j,
		"webhookEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscription) WebhookEndpointInput() *EventgridEventSubscriptionWebhookEndpoint {
	var returns *EventgridEventSubscriptionWebhookEndpoint
	_jsii_.Get(
		j,
		"webhookEndpointInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/eventgrid_event_subscription azurerm_eventgrid_event_subscription} Resource.
func NewEventgridEventSubscription(scope constructs.Construct, id *string, config *EventgridEventSubscriptionConfig) EventgridEventSubscription {
	_init_.Initialize()

	if err := validateNewEventgridEventSubscriptionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridEventSubscription{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/eventgrid_event_subscription azurerm_eventgrid_event_subscription} Resource.
func NewEventgridEventSubscription_Override(e EventgridEventSubscription, scope constructs.Construct, id *string, config *EventgridEventSubscriptionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetAdvancedFilteringOnArraysEnabled(val interface{}) {
	if err := j.validateSetAdvancedFilteringOnArraysEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedFilteringOnArraysEnabled",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetEventDeliverySchema(val *string) {
	if err := j.validateSetEventDeliverySchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventDeliverySchema",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetEventhubId(val *string) {
	if err := j.validateSetEventhubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventhubId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetExpirationTimeUtc(val *string) {
	if err := j.validateSetExpirationTimeUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expirationTimeUtc",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetHybridConnectionId(val *string) {
	if err := j.validateSetHybridConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hybridConnectionId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetIncludedEventTypes(val *[]*string) {
	if err := j.validateSetIncludedEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedEventTypes",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetLabels(val *[]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetServiceBusQueueId(val *string) {
	if err := j.validateSetServiceBusQueueIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceBusQueueId",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscription)SetServiceBusTopicId(val *string) {
	if err := j.validateSetServiceBusTopicIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceBusTopicId",
		val,
	)
}

// Generates CDKTN code for importing a EventgridEventSubscription resource upon running "cdktn plan <stack-name>".
func EventgridEventSubscription_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateEventgridEventSubscription_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
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
func EventgridEventSubscription_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridEventSubscription_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridEventSubscription_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridEventSubscription_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EventgridEventSubscription_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEventgridEventSubscription_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EventgridEventSubscription_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.eventgridEventSubscription.EventgridEventSubscription",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := e.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutAdvancedFilter(value *EventgridEventSubscriptionAdvancedFilter) {
	if err := e.validatePutAdvancedFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAdvancedFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutAzureFunctionEndpoint(value *EventgridEventSubscriptionAzureFunctionEndpoint) {
	if err := e.validatePutAzureFunctionEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAzureFunctionEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeadLetterIdentity(value *EventgridEventSubscriptionDeadLetterIdentity) {
	if err := e.validatePutDeadLetterIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeadLetterIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeliveryIdentity(value *EventgridEventSubscriptionDeliveryIdentity) {
	if err := e.validatePutDeliveryIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeliveryIdentity",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutDeliveryProperty(value interface{}) {
	if err := e.validatePutDeliveryPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putDeliveryProperty",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutRetryPolicy(value *EventgridEventSubscriptionRetryPolicy) {
	if err := e.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutStorageBlobDeadLetterDestination(value *EventgridEventSubscriptionStorageBlobDeadLetterDestination) {
	if err := e.validatePutStorageBlobDeadLetterDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageBlobDeadLetterDestination",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutStorageQueueEndpoint(value *EventgridEventSubscriptionStorageQueueEndpoint) {
	if err := e.validatePutStorageQueueEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStorageQueueEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutSubjectFilter(value *EventgridEventSubscriptionSubjectFilter) {
	if err := e.validatePutSubjectFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSubjectFilter",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutTimeouts(value *EventgridEventSubscriptionTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) PutWebhookEndpoint(value *EventgridEventSubscriptionWebhookEndpoint) {
	if err := e.validatePutWebhookEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putWebhookEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetAdvancedFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetAdvancedFilteringOnArraysEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedFilteringOnArraysEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetAzureFunctionEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetAzureFunctionEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetDeadLetterIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetDeadLetterIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetDeliveryIdentity() {
	_jsii_.InvokeVoid(
		e,
		"resetDeliveryIdentity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetDeliveryProperty() {
	_jsii_.InvokeVoid(
		e,
		"resetDeliveryProperty",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetEventDeliverySchema() {
	_jsii_.InvokeVoid(
		e,
		"resetEventDeliverySchema",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetEventhubId() {
	_jsii_.InvokeVoid(
		e,
		"resetEventhubId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetExpirationTimeUtc() {
	_jsii_.InvokeVoid(
		e,
		"resetExpirationTimeUtc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetHybridConnectionId() {
	_jsii_.InvokeVoid(
		e,
		"resetHybridConnectionId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetIncludedEventTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetIncludedEventTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetServiceBusQueueId() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceBusQueueId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetServiceBusTopicId() {
	_jsii_.InvokeVoid(
		e,
		"resetServiceBusTopicId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetStorageBlobDeadLetterDestination() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageBlobDeadLetterDestination",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetStorageQueueEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetStorageQueueEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetSubjectFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetSubjectFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) ResetWebhookEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetWebhookEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscription) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscription) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

