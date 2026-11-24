// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package horizondbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v19/horizondbcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster azurerm_horizondb_cluster}.
type HorizondbCluster interface {
	cdktn.TerraformResource
	AdministratorLogin() *string
	SetAdministratorLogin(val *string)
	AdministratorLoginInput() *string
	AdministratorPassword() *string
	SetAdministratorPassword(val *string)
	AdministratorPasswordInput() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	AdministratorPasswordWo() *string
	// Deprecated: Write-only: the provider never returns this value; reading it always yields null by protocol contract. The getter remains for compatibility and will be removed in a future prebuilt-provider major.
	SetAdministratorPasswordWo(val *string)
	AdministratorPasswordWoInput() *string
	AdministratorPasswordWoVersion() *float64
	SetAdministratorPasswordWoVersion(val *float64)
	AdministratorPasswordWoVersionInput() *float64
	BackupRetentionInDays() *float64
	SetBackupRetentionInDays(val *float64)
	BackupRetentionInDaysInput() *float64
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
	GeoRedundantBackupEnabled() interface{}
	SetGeoRedundantBackupEnabled(val interface{})
	GeoRedundantBackupEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() HorizondbClusterIdentityOutputReference
	IdentityInput() *HorizondbClusterIdentity
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
	PointInTimeRestore() HorizondbClusterPointInTimeRestoreOutputReference
	PointInTimeRestoreInput() *HorizondbClusterPointInTimeRestore
	PostgresqlVersion() *string
	SetPostgresqlVersion(val *string)
	PostgresqlVersionInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicNetworkAccess() *string
	SetPublicNetworkAccess(val *string)
	PublicNetworkAccessInput() *string
	// Experimental.
	RawOverrides() interface{}
	ReadWritePool() HorizondbClusterReadWritePoolOutputReference
	ReadWritePoolInput() *HorizondbClusterReadWritePool
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
	Timeouts() HorizondbClusterTimeoutsOutputReference
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
	PutIdentity(value *HorizondbClusterIdentity)
	PutPointInTimeRestore(value *HorizondbClusterPointInTimeRestore)
	PutReadWritePool(value *HorizondbClusterReadWritePool)
	PutTimeouts(value *HorizondbClusterTimeouts)
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
	ResetAdministratorPassword()
	ResetAdministratorPasswordWo()
	ResetAdministratorPasswordWoVersion()
	ResetBackupRetentionInDays()
	ResetGeoRedundantBackupEnabled()
	ResetId()
	ResetIdentity()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRestore()
	ResetPostgresqlVersion()
	ResetPublicNetworkAccess()
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

// The jsii proxy struct for HorizondbCluster
type jsiiProxy_HorizondbCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_HorizondbCluster) AdministratorLogin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) AdministratorLoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) AdministratorPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) AdministratorPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) AdministratorPasswordWo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordWo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) AdministratorPasswordWoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"administratorPasswordWoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) AdministratorPasswordWoVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"administratorPasswordWoVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) AdministratorPasswordWoVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"administratorPasswordWoVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) BackupRetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) BackupRetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) GeoRedundantBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) GeoRedundantBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"geoRedundantBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Identity() HorizondbClusterIdentityOutputReference {
	var returns HorizondbClusterIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) IdentityInput() *HorizondbClusterIdentity {
	var returns *HorizondbClusterIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) PointInTimeRestore() HorizondbClusterPointInTimeRestoreOutputReference {
	var returns HorizondbClusterPointInTimeRestoreOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) PointInTimeRestoreInput() *HorizondbClusterPointInTimeRestore {
	var returns *HorizondbClusterPointInTimeRestore
	_jsii_.Get(
		j,
		"pointInTimeRestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) PostgresqlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) PostgresqlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgresqlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) PublicNetworkAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) PublicNetworkAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicNetworkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) ReadWritePool() HorizondbClusterReadWritePoolOutputReference {
	var returns HorizondbClusterReadWritePoolOutputReference
	_jsii_.Get(
		j,
		"readWritePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) ReadWritePoolInput() *HorizondbClusterReadWritePool {
	var returns *HorizondbClusterReadWritePool
	_jsii_.Get(
		j,
		"readWritePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) Timeouts() HorizondbClusterTimeoutsOutputReference {
	var returns HorizondbClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HorizondbCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster azurerm_horizondb_cluster} Resource.
func NewHorizondbCluster(scope constructs.Construct, id *string, config *HorizondbClusterConfig) HorizondbCluster {
	_init_.Initialize()

	if err := validateNewHorizondbClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_HorizondbCluster{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.15.0/docs/resources/horizondb_cluster azurerm_horizondb_cluster} Resource.
func NewHorizondbCluster_Override(h HorizondbCluster, scope constructs.Construct, id *string, config *HorizondbClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbCluster",
		[]interface{}{scope, id, config},
		h,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetAdministratorLogin(val *string) {
	if err := j.validateSetAdministratorLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorLogin",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetAdministratorPassword(val *string) {
	if err := j.validateSetAdministratorPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPassword",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetAdministratorPasswordWo(val *string) {
	if err := j.validateSetAdministratorPasswordWoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPasswordWo",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetAdministratorPasswordWoVersion(val *float64) {
	if err := j.validateSetAdministratorPasswordWoVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"administratorPasswordWoVersion",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetBackupRetentionInDays(val *float64) {
	if err := j.validateSetBackupRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionInDays",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetGeoRedundantBackupEnabled(val interface{}) {
	if err := j.validateSetGeoRedundantBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"geoRedundantBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetPostgresqlVersion(val *string) {
	if err := j.validateSetPostgresqlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgresqlVersion",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetPublicNetworkAccess(val *string) {
	if err := j.validateSetPublicNetworkAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicNetworkAccess",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_HorizondbCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a HorizondbCluster resource upon running "cdktn plan <stack-name>".
func HorizondbCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateHorizondbCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbCluster",
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
func HorizondbCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHorizondbCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HorizondbCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHorizondbCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func HorizondbCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHorizondbCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func HorizondbCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.horizondbCluster.HorizondbCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (h *jsiiProxy_HorizondbCluster) AddMoveTarget(moveTarget *string) {
	if err := h.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (h *jsiiProxy_HorizondbCluster) AddOverride(path *string, value interface{}) {
	if err := h.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (h *jsiiProxy_HorizondbCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HorizondbCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (h *jsiiProxy_HorizondbCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HorizondbCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HorizondbCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HorizondbCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HorizondbCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HorizondbCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HorizondbCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HorizondbCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := h.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (h *jsiiProxy_HorizondbCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (h *jsiiProxy_HorizondbCluster) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := h.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) MoveFromId(id *string) {
	if err := h.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveFromId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HorizondbCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := h.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (h *jsiiProxy_HorizondbCluster) MoveToId(id *string) {
	if err := h.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"moveToId",
		[]interface{}{id},
	)
}

func (h *jsiiProxy_HorizondbCluster) OverrideLogicalId(newLogicalId *string) {
	if err := h.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (h *jsiiProxy_HorizondbCluster) PutIdentity(value *HorizondbClusterIdentity) {
	if err := h.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putIdentity",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizondbCluster) PutPointInTimeRestore(value *HorizondbClusterPointInTimeRestore) {
	if err := h.validatePutPointInTimeRestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putPointInTimeRestore",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizondbCluster) PutReadWritePool(value *HorizondbClusterReadWritePool) {
	if err := h.validatePutReadWritePoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putReadWritePool",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizondbCluster) PutTimeouts(value *HorizondbClusterTimeouts) {
	if err := h.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HorizondbCluster) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := h.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetAdministratorPassword() {
	_jsii_.InvokeVoid(
		h,
		"resetAdministratorPassword",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetAdministratorPasswordWo() {
	_jsii_.InvokeVoid(
		h,
		"resetAdministratorPasswordWo",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetAdministratorPasswordWoVersion() {
	_jsii_.InvokeVoid(
		h,
		"resetAdministratorPasswordWoVersion",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetBackupRetentionInDays() {
	_jsii_.InvokeVoid(
		h,
		"resetBackupRetentionInDays",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetGeoRedundantBackupEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetGeoRedundantBackupEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetId() {
	_jsii_.InvokeVoid(
		h,
		"resetId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetIdentity() {
	_jsii_.InvokeVoid(
		h,
		"resetIdentity",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		h,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetPointInTimeRestore() {
	_jsii_.InvokeVoid(
		h,
		"resetPointInTimeRestore",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetPostgresqlVersion() {
	_jsii_.InvokeVoid(
		h,
		"resetPostgresqlVersion",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetPublicNetworkAccess() {
	_jsii_.InvokeVoid(
		h,
		"resetPublicNetworkAccess",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetTags() {
	_jsii_.InvokeVoid(
		h,
		"resetTags",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		h,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HorizondbCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HorizondbCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

