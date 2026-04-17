// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package botchannelsregistration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v15/botchannelsregistration/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/bot_channels_registration azurerm_bot_channels_registration}.
type BotChannelsRegistration interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CmkKeyVaultUrl() *string
	SetCmkKeyVaultUrl(val *string)
	CmkKeyVaultUrlInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeveloperAppInsightsApiKey() *string
	SetDeveloperAppInsightsApiKey(val *string)
	DeveloperAppInsightsApiKeyInput() *string
	DeveloperAppInsightsApplicationId() *string
	SetDeveloperAppInsightsApplicationId(val *string)
	DeveloperAppInsightsApplicationIdInput() *string
	DeveloperAppInsightsKey() *string
	SetDeveloperAppInsightsKey(val *string)
	DeveloperAppInsightsKeyInput() *string
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IconUrl() *string
	SetIconUrl(val *string)
	IconUrlInput() *string
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
	MicrosoftAppId() *string
	SetMicrosoftAppId(val *string)
	MicrosoftAppIdInput() *string
	MicrosoftAppTenantId() *string
	SetMicrosoftAppTenantId(val *string)
	MicrosoftAppTenantIdInput() *string
	MicrosoftAppType() *string
	SetMicrosoftAppType(val *string)
	MicrosoftAppTypeInput() *string
	MicrosoftAppUserAssignedIdentityId() *string
	SetMicrosoftAppUserAssignedIdentityId(val *string)
	MicrosoftAppUserAssignedIdentityIdInput() *string
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
	PublicNetworkAccessEnabled() interface{}
	SetPublicNetworkAccessEnabled(val interface{})
	PublicNetworkAccessEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	StreamingEndpointEnabled() interface{}
	SetStreamingEndpointEnabled(val interface{})
	StreamingEndpointEnabledInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BotChannelsRegistrationTimeoutsOutputReference
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
	PutTimeouts(value *BotChannelsRegistrationTimeouts)
	ResetCmkKeyVaultUrl()
	ResetDescription()
	ResetDeveloperAppInsightsApiKey()
	ResetDeveloperAppInsightsApplicationId()
	ResetDeveloperAppInsightsKey()
	ResetDisplayName()
	ResetEndpoint()
	ResetIconUrl()
	ResetId()
	ResetMicrosoftAppTenantId()
	ResetMicrosoftAppType()
	ResetMicrosoftAppUserAssignedIdentityId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetStreamingEndpointEnabled()
	ResetTags()
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

// The jsii proxy struct for BotChannelsRegistration
type jsiiProxy_BotChannelsRegistration struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_BotChannelsRegistration) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) CmkKeyVaultUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkKeyVaultUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) CmkKeyVaultUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkKeyVaultUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DeveloperAppInsightsApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DeveloperAppInsightsApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DeveloperAppInsightsApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DeveloperAppInsightsApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DeveloperAppInsightsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DeveloperAppInsightsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"developerAppInsightsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) IconUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) IconUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppTenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppTenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppTenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppTenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppUserAssignedIdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppUserAssignedIdentityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) MicrosoftAppUserAssignedIdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"microsoftAppUserAssignedIdentityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) StreamingEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) StreamingEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamingEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) Timeouts() BotChannelsRegistrationTimeoutsOutputReference {
	var returns BotChannelsRegistrationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BotChannelsRegistration) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/bot_channels_registration azurerm_bot_channels_registration} Resource.
func NewBotChannelsRegistration(scope constructs.Construct, id *string, config *BotChannelsRegistrationConfig) BotChannelsRegistration {
	_init_.Initialize()

	if err := validateNewBotChannelsRegistrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BotChannelsRegistration{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.botChannelsRegistration.BotChannelsRegistration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.69.0/docs/resources/bot_channels_registration azurerm_bot_channels_registration} Resource.
func NewBotChannelsRegistration_Override(b BotChannelsRegistration, scope constructs.Construct, id *string, config *BotChannelsRegistrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.botChannelsRegistration.BotChannelsRegistration",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetCmkKeyVaultUrl(val *string) {
	if err := j.validateSetCmkKeyVaultUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmkKeyVaultUrl",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetDeveloperAppInsightsApiKey(val *string) {
	if err := j.validateSetDeveloperAppInsightsApiKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerAppInsightsApiKey",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetDeveloperAppInsightsApplicationId(val *string) {
	if err := j.validateSetDeveloperAppInsightsApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerAppInsightsApplicationId",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetDeveloperAppInsightsKey(val *string) {
	if err := j.validateSetDeveloperAppInsightsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerAppInsightsKey",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetIconUrl(val *string) {
	if err := j.validateSetIconUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iconUrl",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetMicrosoftAppId(val *string) {
	if err := j.validateSetMicrosoftAppIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftAppId",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetMicrosoftAppTenantId(val *string) {
	if err := j.validateSetMicrosoftAppTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftAppTenantId",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetMicrosoftAppType(val *string) {
	if err := j.validateSetMicrosoftAppTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftAppType",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetMicrosoftAppUserAssignedIdentityId(val *string) {
	if err := j.validateSetMicrosoftAppUserAssignedIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"microsoftAppUserAssignedIdentityId",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetStreamingEndpointEnabled(val interface{}) {
	if err := j.validateSetStreamingEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_BotChannelsRegistration)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a BotChannelsRegistration resource upon running "cdktn plan <stack-name>".
func BotChannelsRegistration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateBotChannelsRegistration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.botChannelsRegistration.BotChannelsRegistration",
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
func BotChannelsRegistration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotChannelsRegistration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.botChannelsRegistration.BotChannelsRegistration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BotChannelsRegistration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotChannelsRegistration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.botChannelsRegistration.BotChannelsRegistration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BotChannelsRegistration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBotChannelsRegistration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.botChannelsRegistration.BotChannelsRegistration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BotChannelsRegistration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.botChannelsRegistration.BotChannelsRegistration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) PutTimeouts(value *BotChannelsRegistrationTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetCmkKeyVaultUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetCmkKeyVaultUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetDeveloperAppInsightsApiKey() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsApiKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetDeveloperAppInsightsApplicationId() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsApplicationId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetDeveloperAppInsightsKey() {
	_jsii_.InvokeVoid(
		b,
		"resetDeveloperAppInsightsKey",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetDisplayName() {
	_jsii_.InvokeVoid(
		b,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetEndpoint() {
	_jsii_.InvokeVoid(
		b,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetIconUrl() {
	_jsii_.InvokeVoid(
		b,
		"resetIconUrl",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetMicrosoftAppTenantId() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftAppTenantId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetMicrosoftAppType() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftAppType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetMicrosoftAppUserAssignedIdentityId() {
	_jsii_.InvokeVoid(
		b,
		"resetMicrosoftAppUserAssignedIdentityId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetStreamingEndpointEnabled() {
	_jsii_.InvokeVoid(
		b,
		"resetStreamingEndpointEnabled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BotChannelsRegistration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BotChannelsRegistration) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		b,
		"with",
		args,
		&returns,
	)

	return returns
}

