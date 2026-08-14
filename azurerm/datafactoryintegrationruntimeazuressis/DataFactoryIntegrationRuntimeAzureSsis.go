// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datafactoryintegrationruntimeazuressis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/datafactoryintegrationruntimeazuressis/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/data_factory_integration_runtime_azure_ssis azurerm_data_factory_integration_runtime_azure_ssis}.
type DataFactoryIntegrationRuntimeAzureSsis interface {
	cdktn.TerraformResource
	CatalogInfo() DataFactoryIntegrationRuntimeAzureSsisCatalogInfoOutputReference
	CatalogInfoInput() *DataFactoryIntegrationRuntimeAzureSsisCatalogInfo
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyComputeScale() DataFactoryIntegrationRuntimeAzureSsisCopyComputeScaleOutputReference
	CopyComputeScaleInput() *DataFactoryIntegrationRuntimeAzureSsisCopyComputeScale
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CredentialName() *string
	SetCredentialName(val *string)
	CredentialNameInput() *string
	CustomSetupScript() DataFactoryIntegrationRuntimeAzureSsisCustomSetupScriptOutputReference
	CustomSetupScriptInput() *DataFactoryIntegrationRuntimeAzureSsisCustomSetupScript
	DataFactoryId() *string
	SetDataFactoryId(val *string)
	DataFactoryIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	ExpressCustomSetup() DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference
	ExpressCustomSetupInput() *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup
	ExpressVnetIntegration() DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegrationOutputReference
	ExpressVnetIntegrationInput() *DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegration
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
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxParallelExecutionsPerNode() *float64
	SetMaxParallelExecutionsPerNode(val *float64)
	MaxParallelExecutionsPerNodeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeSize() *string
	SetNodeSize(val *string)
	NodeSizeInput() *string
	NumberOfNodes() *float64
	SetNumberOfNodes(val *float64)
	NumberOfNodesInput() *float64
	PackageStore() DataFactoryIntegrationRuntimeAzureSsisPackageStoreList
	PackageStoreInput() interface{}
	PipelineExternalComputeScale() DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference
	PipelineExternalComputeScaleInput() *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Proxy() DataFactoryIntegrationRuntimeAzureSsisProxyOutputReference
	ProxyInput() *DataFactoryIntegrationRuntimeAzureSsisProxy
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFactoryIntegrationRuntimeAzureSsisTimeoutsOutputReference
	TimeoutsInput() interface{}
	VnetIntegration() DataFactoryIntegrationRuntimeAzureSsisVnetIntegrationOutputReference
	VnetIntegrationInput() *DataFactoryIntegrationRuntimeAzureSsisVnetIntegration
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
	PutCatalogInfo(value *DataFactoryIntegrationRuntimeAzureSsisCatalogInfo)
	PutCopyComputeScale(value *DataFactoryIntegrationRuntimeAzureSsisCopyComputeScale)
	PutCustomSetupScript(value *DataFactoryIntegrationRuntimeAzureSsisCustomSetupScript)
	PutExpressCustomSetup(value *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup)
	PutExpressVnetIntegration(value *DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegration)
	PutPackageStore(value interface{})
	PutPipelineExternalComputeScale(value *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale)
	PutProxy(value *DataFactoryIntegrationRuntimeAzureSsisProxy)
	PutTimeouts(value *DataFactoryIntegrationRuntimeAzureSsisTimeouts)
	PutVnetIntegration(value *DataFactoryIntegrationRuntimeAzureSsisVnetIntegration)
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
	ResetCatalogInfo()
	ResetCopyComputeScale()
	ResetCredentialName()
	ResetCustomSetupScript()
	ResetDescription()
	ResetEdition()
	ResetExpressCustomSetup()
	ResetExpressVnetIntegration()
	ResetId()
	ResetLicenseType()
	ResetMaxParallelExecutionsPerNode()
	ResetNumberOfNodes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackageStore()
	ResetPipelineExternalComputeScale()
	ResetProxy()
	ResetTimeouts()
	ResetVnetIntegration()
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

// The jsii proxy struct for DataFactoryIntegrationRuntimeAzureSsis
type jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CatalogInfo() DataFactoryIntegrationRuntimeAzureSsisCatalogInfoOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisCatalogInfoOutputReference
	_jsii_.Get(
		j,
		"catalogInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CatalogInfoInput() *DataFactoryIntegrationRuntimeAzureSsisCatalogInfo {
	var returns *DataFactoryIntegrationRuntimeAzureSsisCatalogInfo
	_jsii_.Get(
		j,
		"catalogInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CopyComputeScale() DataFactoryIntegrationRuntimeAzureSsisCopyComputeScaleOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisCopyComputeScaleOutputReference
	_jsii_.Get(
		j,
		"copyComputeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CopyComputeScaleInput() *DataFactoryIntegrationRuntimeAzureSsisCopyComputeScale {
	var returns *DataFactoryIntegrationRuntimeAzureSsisCopyComputeScale
	_jsii_.Get(
		j,
		"copyComputeScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CredentialName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CredentialNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CustomSetupScript() DataFactoryIntegrationRuntimeAzureSsisCustomSetupScriptOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisCustomSetupScriptOutputReference
	_jsii_.Get(
		j,
		"customSetupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) CustomSetupScriptInput() *DataFactoryIntegrationRuntimeAzureSsisCustomSetupScript {
	var returns *DataFactoryIntegrationRuntimeAzureSsisCustomSetupScript
	_jsii_.Get(
		j,
		"customSetupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) DataFactoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) DataFactoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ExpressCustomSetup() DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetupOutputReference
	_jsii_.Get(
		j,
		"expressCustomSetup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ExpressCustomSetupInput() *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup {
	var returns *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup
	_jsii_.Get(
		j,
		"expressCustomSetupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ExpressVnetIntegration() DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegrationOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegrationOutputReference
	_jsii_.Get(
		j,
		"expressVnetIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ExpressVnetIntegrationInput() *DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegration {
	var returns *DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegration
	_jsii_.Get(
		j,
		"expressVnetIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) MaxParallelExecutionsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelExecutionsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) MaxParallelExecutionsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelExecutionsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) NodeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) NodeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) NumberOfNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PackageStore() DataFactoryIntegrationRuntimeAzureSsisPackageStoreList {
	var returns DataFactoryIntegrationRuntimeAzureSsisPackageStoreList
	_jsii_.Get(
		j,
		"packageStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PackageStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"packageStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PipelineExternalComputeScale() DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScaleOutputReference
	_jsii_.Get(
		j,
		"pipelineExternalComputeScale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PipelineExternalComputeScaleInput() *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale {
	var returns *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale
	_jsii_.Get(
		j,
		"pipelineExternalComputeScaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Proxy() DataFactoryIntegrationRuntimeAzureSsisProxyOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisProxyOutputReference
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ProxyInput() *DataFactoryIntegrationRuntimeAzureSsisProxy {
	var returns *DataFactoryIntegrationRuntimeAzureSsisProxy
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) Timeouts() DataFactoryIntegrationRuntimeAzureSsisTimeoutsOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) VnetIntegration() DataFactoryIntegrationRuntimeAzureSsisVnetIntegrationOutputReference {
	var returns DataFactoryIntegrationRuntimeAzureSsisVnetIntegrationOutputReference
	_jsii_.Get(
		j,
		"vnetIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) VnetIntegrationInput() *DataFactoryIntegrationRuntimeAzureSsisVnetIntegration {
	var returns *DataFactoryIntegrationRuntimeAzureSsisVnetIntegration
	_jsii_.Get(
		j,
		"vnetIntegrationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/data_factory_integration_runtime_azure_ssis azurerm_data_factory_integration_runtime_azure_ssis} Resource.
func NewDataFactoryIntegrationRuntimeAzureSsis(scope constructs.Construct, id *string, config *DataFactoryIntegrationRuntimeAzureSsisConfig) DataFactoryIntegrationRuntimeAzureSsis {
	_init_.Initialize()

	if err := validateNewDataFactoryIntegrationRuntimeAzureSsisParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsis",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.1.0/docs/resources/data_factory_integration_runtime_azure_ssis azurerm_data_factory_integration_runtime_azure_ssis} Resource.
func NewDataFactoryIntegrationRuntimeAzureSsis_Override(d DataFactoryIntegrationRuntimeAzureSsis, scope constructs.Construct, id *string, config *DataFactoryIntegrationRuntimeAzureSsisConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsis",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetCredentialName(val *string) {
	if err := j.validateSetCredentialNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetDataFactoryId(val *string) {
	if err := j.validateSetDataFactoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetMaxParallelExecutionsPerNode(val *float64) {
	if err := j.validateSetMaxParallelExecutionsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelExecutionsPerNode",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetNodeSize(val *string) {
	if err := j.validateSetNodeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSize",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetNumberOfNodes(val *float64) {
	if err := j.validateSetNumberOfNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a DataFactoryIntegrationRuntimeAzureSsis resource upon running "cdktn plan <stack-name>".
func DataFactoryIntegrationRuntimeAzureSsis_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataFactoryIntegrationRuntimeAzureSsis_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsis",
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
func DataFactoryIntegrationRuntimeAzureSsis_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryIntegrationRuntimeAzureSsis_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsis",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryIntegrationRuntimeAzureSsis_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryIntegrationRuntimeAzureSsis_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsis",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryIntegrationRuntimeAzureSsis_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryIntegrationRuntimeAzureSsis_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsis",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFactoryIntegrationRuntimeAzureSsis_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.dataFactoryIntegrationRuntimeAzureSsis.DataFactoryIntegrationRuntimeAzureSsis",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutCatalogInfo(value *DataFactoryIntegrationRuntimeAzureSsisCatalogInfo) {
	if err := d.validatePutCatalogInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCatalogInfo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutCopyComputeScale(value *DataFactoryIntegrationRuntimeAzureSsisCopyComputeScale) {
	if err := d.validatePutCopyComputeScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCopyComputeScale",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutCustomSetupScript(value *DataFactoryIntegrationRuntimeAzureSsisCustomSetupScript) {
	if err := d.validatePutCustomSetupScriptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomSetupScript",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutExpressCustomSetup(value *DataFactoryIntegrationRuntimeAzureSsisExpressCustomSetup) {
	if err := d.validatePutExpressCustomSetupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExpressCustomSetup",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutExpressVnetIntegration(value *DataFactoryIntegrationRuntimeAzureSsisExpressVnetIntegration) {
	if err := d.validatePutExpressVnetIntegrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExpressVnetIntegration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutPackageStore(value interface{}) {
	if err := d.validatePutPackageStoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPackageStore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutPipelineExternalComputeScale(value *DataFactoryIntegrationRuntimeAzureSsisPipelineExternalComputeScale) {
	if err := d.validatePutPipelineExternalComputeScaleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPipelineExternalComputeScale",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutProxy(value *DataFactoryIntegrationRuntimeAzureSsisProxy) {
	if err := d.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProxy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutTimeouts(value *DataFactoryIntegrationRuntimeAzureSsisTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) PutVnetIntegration(value *DataFactoryIntegrationRuntimeAzureSsisVnetIntegration) {
	if err := d.validatePutVnetIntegrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVnetIntegration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetCatalogInfo() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogInfo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetCopyComputeScale() {
	_jsii_.InvokeVoid(
		d,
		"resetCopyComputeScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetCredentialName() {
	_jsii_.InvokeVoid(
		d,
		"resetCredentialName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetCustomSetupScript() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSetupScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetEdition() {
	_jsii_.InvokeVoid(
		d,
		"resetEdition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetExpressCustomSetup() {
	_jsii_.InvokeVoid(
		d,
		"resetExpressCustomSetup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetExpressVnetIntegration() {
	_jsii_.InvokeVoid(
		d,
		"resetExpressVnetIntegration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetLicenseType() {
	_jsii_.InvokeVoid(
		d,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetMaxParallelExecutionsPerNode() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxParallelExecutionsPerNode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetNumberOfNodes() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfNodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetPackageStore() {
	_jsii_.InvokeVoid(
		d,
		"resetPackageStore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetPipelineExternalComputeScale() {
	_jsii_.InvokeVoid(
		d,
		"resetPipelineExternalComputeScale",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetProxy() {
	_jsii_.InvokeVoid(
		d,
		"resetProxy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ResetVnetIntegration() {
	_jsii_.InvokeVoid(
		d,
		"resetVnetIntegration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeAzureSsis) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

