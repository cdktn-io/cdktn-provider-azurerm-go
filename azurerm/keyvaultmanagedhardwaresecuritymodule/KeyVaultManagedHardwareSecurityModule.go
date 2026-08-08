// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/keyvaultmanagedhardwaresecuritymodule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/key_vault_managed_hardware_security_module azurerm_key_vault_managed_hardware_security_module}.
type KeyVaultManagedHardwareSecurityModule interface {
	cdktn.TerraformResource
	AdminObjectIds() *[]*string
	SetAdminObjectIds(val *[]*string)
	AdminObjectIdsInput() *[]*string
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
	HsmUri() *string
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
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkAcls() KeyVaultManagedHardwareSecurityModuleNetworkAclsOutputReference
	NetworkAclsInput() *KeyVaultManagedHardwareSecurityModuleNetworkAcls
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
	PurgeProtectionEnabled() interface{}
	SetPurgeProtectionEnabled(val interface{})
	PurgeProtectionEnabledInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecurityDomainEncryptedData() *string
	SecurityDomainKeyVaultCertificateIds() *[]*string
	SetSecurityDomainKeyVaultCertificateIds(val *[]*string)
	SecurityDomainKeyVaultCertificateIdsInput() *[]*string
	SecurityDomainQuorum() *float64
	SetSecurityDomainQuorum(val *float64)
	SecurityDomainQuorumInput() *float64
	SkuName() *string
	SetSkuName(val *string)
	SkuNameInput() *string
	SoftDeleteRetentionDays() *float64
	SetSoftDeleteRetentionDays(val *float64)
	SoftDeleteRetentionDaysInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KeyVaultManagedHardwareSecurityModuleTimeoutsOutputReference
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
	PutNetworkAcls(value *KeyVaultManagedHardwareSecurityModuleNetworkAcls)
	PutTimeouts(value *KeyVaultManagedHardwareSecurityModuleTimeouts)
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
	ResetId()
	ResetNetworkAcls()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPublicNetworkAccessEnabled()
	ResetPurgeProtectionEnabled()
	ResetSecurityDomainKeyVaultCertificateIds()
	ResetSecurityDomainQuorum()
	ResetSoftDeleteRetentionDays()
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

// The jsii proxy struct for KeyVaultManagedHardwareSecurityModule
type jsiiProxy_KeyVaultManagedHardwareSecurityModule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) AdminObjectIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminObjectIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) AdminObjectIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminObjectIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) HsmUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hsmUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) NetworkAcls() KeyVaultManagedHardwareSecurityModuleNetworkAclsOutputReference {
	var returns KeyVaultManagedHardwareSecurityModuleNetworkAclsOutputReference
	_jsii_.Get(
		j,
		"networkAcls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) NetworkAclsInput() *KeyVaultManagedHardwareSecurityModuleNetworkAcls {
	var returns *KeyVaultManagedHardwareSecurityModuleNetworkAcls
	_jsii_.Get(
		j,
		"networkAclsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) PublicNetworkAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) PublicNetworkAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicNetworkAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) PurgeProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) PurgeProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"purgeProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SecurityDomainEncryptedData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityDomainEncryptedData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SecurityDomainKeyVaultCertificateIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityDomainKeyVaultCertificateIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SecurityDomainKeyVaultCertificateIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityDomainKeyVaultCertificateIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SecurityDomainQuorum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityDomainQuorum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SecurityDomainQuorumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"securityDomainQuorumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SoftDeleteRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softDeleteRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SoftDeleteRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"softDeleteRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) Timeouts() KeyVaultManagedHardwareSecurityModuleTimeoutsOutputReference {
	var returns KeyVaultManagedHardwareSecurityModuleTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/key_vault_managed_hardware_security_module azurerm_key_vault_managed_hardware_security_module} Resource.
func NewKeyVaultManagedHardwareSecurityModule(scope constructs.Construct, id *string, config *KeyVaultManagedHardwareSecurityModuleConfig) KeyVaultManagedHardwareSecurityModule {
	_init_.Initialize()

	if err := validateNewKeyVaultManagedHardwareSecurityModuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultManagedHardwareSecurityModule{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.keyVaultManagedHardwareSecurityModule.KeyVaultManagedHardwareSecurityModule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.0.1/docs/resources/key_vault_managed_hardware_security_module azurerm_key_vault_managed_hardware_security_module} Resource.
func NewKeyVaultManagedHardwareSecurityModule_Override(k KeyVaultManagedHardwareSecurityModule, scope constructs.Construct, id *string, config *KeyVaultManagedHardwareSecurityModuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.keyVaultManagedHardwareSecurityModule.KeyVaultManagedHardwareSecurityModule",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetAdminObjectIds(val *[]*string) {
	if err := j.validateSetAdminObjectIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminObjectIds",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetPublicNetworkAccessEnabled(val interface{}) {
	if err := j.validateSetPublicNetworkAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetPurgeProtectionEnabled(val interface{}) {
	if err := j.validateSetPurgeProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purgeProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetSecurityDomainKeyVaultCertificateIds(val *[]*string) {
	if err := j.validateSetSecurityDomainKeyVaultCertificateIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDomainKeyVaultCertificateIds",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetSecurityDomainQuorum(val *float64) {
	if err := j.validateSetSecurityDomainQuorumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityDomainQuorum",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetSkuName(val *string) {
	if err := j.validateSetSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skuName",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetSoftDeleteRetentionDays(val *float64) {
	if err := j.validateSetSoftDeleteRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softDeleteRetentionDays",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModule)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

// Generates CDKTN code for importing a KeyVaultManagedHardwareSecurityModule resource upon running "cdktn plan <stack-name>".
func KeyVaultManagedHardwareSecurityModule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.keyVaultManagedHardwareSecurityModule.KeyVaultManagedHardwareSecurityModule",
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
func KeyVaultManagedHardwareSecurityModule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.keyVaultManagedHardwareSecurityModule.KeyVaultManagedHardwareSecurityModule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyVaultManagedHardwareSecurityModule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.keyVaultManagedHardwareSecurityModule.KeyVaultManagedHardwareSecurityModule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyVaultManagedHardwareSecurityModule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.keyVaultManagedHardwareSecurityModule.KeyVaultManagedHardwareSecurityModule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KeyVaultManagedHardwareSecurityModule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.keyVaultManagedHardwareSecurityModule.KeyVaultManagedHardwareSecurityModule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := k.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) PutNetworkAcls(value *KeyVaultManagedHardwareSecurityModuleNetworkAcls) {
	if err := k.validatePutNetworkAclsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putNetworkAcls",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) PutTimeouts(value *KeyVaultManagedHardwareSecurityModuleTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := k.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetNetworkAcls() {
	_jsii_.InvokeVoid(
		k,
		"resetNetworkAcls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetPublicNetworkAccessEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPublicNetworkAccessEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetPurgeProtectionEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPurgeProtectionEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetSecurityDomainKeyVaultCertificateIds() {
	_jsii_.InvokeVoid(
		k,
		"resetSecurityDomainKeyVaultCertificateIds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetSecurityDomainQuorum() {
	_jsii_.InvokeVoid(
		k,
		"resetSecurityDomainQuorum",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetSoftDeleteRetentionDays() {
	_jsii_.InvokeVoid(
		k,
		"resetSoftDeleteRetentionDays",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

