// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kubernetesautomaticcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v16/kubernetesautomaticcluster/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/kubernetes_automatic_cluster azurerm_kubernetes_automatic_cluster}.
type KubernetesAutomaticCluster interface {
	cdktn.TerraformResource
	ApiServerAccess() KubernetesAutomaticClusterApiServerAccessOutputReference
	ApiServerAccessInput() *KubernetesAutomaticClusterApiServerAccess
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
	CurrentKubernetesVersion() *string
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
	FullyQualifiedDomainName() *string
	HostedSystem() KubernetesAutomaticClusterHostedSystemOutputReference
	HostedSystemInput() *KubernetesAutomaticClusterHostedSystem
	Id() *string
	SetId(val *string)
	Identity() KubernetesAutomaticClusterIdentityOutputReference
	IdentityInput() *KubernetesAutomaticClusterIdentity
	IdInput() *string
	KubeConfig() KubernetesAutomaticClusterKubeConfigList
	KubeConfigRaw() *string
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
	NodeResourceGroupId() *string
	OidcIssuerUrl() *string
	PortalFullyQualifiedDomainName() *string
	PrivateCluster() KubernetesAutomaticClusterPrivateClusterOutputReference
	PrivateClusterInput() *KubernetesAutomaticClusterPrivateCluster
	PrivateFullyQualifiedDomainName() *string
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
	ServiceMesh() KubernetesAutomaticClusterServiceMeshOutputReference
	ServiceMeshInput() *KubernetesAutomaticClusterServiceMesh
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() KubernetesAutomaticClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	WebAppRoutingIngress() KubernetesAutomaticClusterWebAppRoutingIngressOutputReference
	WebAppRoutingIngressInput() *KubernetesAutomaticClusterWebAppRoutingIngress
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
	PutApiServerAccess(value *KubernetesAutomaticClusterApiServerAccess)
	PutHostedSystem(value *KubernetesAutomaticClusterHostedSystem)
	PutIdentity(value *KubernetesAutomaticClusterIdentity)
	PutPrivateCluster(value *KubernetesAutomaticClusterPrivateCluster)
	PutServiceMesh(value *KubernetesAutomaticClusterServiceMesh)
	PutTimeouts(value *KubernetesAutomaticClusterTimeouts)
	PutWebAppRoutingIngress(value *KubernetesAutomaticClusterWebAppRoutingIngress)
	ResetApiServerAccess()
	ResetHostedSystem()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateCluster()
	ResetServiceMesh()
	ResetTags()
	ResetTimeouts()
	ResetWebAppRoutingIngress()
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

// The jsii proxy struct for KubernetesAutomaticCluster
type jsiiProxy_KubernetesAutomaticCluster struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ApiServerAccess() KubernetesAutomaticClusterApiServerAccessOutputReference {
	var returns KubernetesAutomaticClusterApiServerAccessOutputReference
	_jsii_.Get(
		j,
		"apiServerAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ApiServerAccessInput() *KubernetesAutomaticClusterApiServerAccess {
	var returns *KubernetesAutomaticClusterApiServerAccess
	_jsii_.Get(
		j,
		"apiServerAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) CurrentKubernetesVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentKubernetesVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) FullyQualifiedDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullyQualifiedDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) HostedSystem() KubernetesAutomaticClusterHostedSystemOutputReference {
	var returns KubernetesAutomaticClusterHostedSystemOutputReference
	_jsii_.Get(
		j,
		"hostedSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) HostedSystemInput() *KubernetesAutomaticClusterHostedSystem {
	var returns *KubernetesAutomaticClusterHostedSystem
	_jsii_.Get(
		j,
		"hostedSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Identity() KubernetesAutomaticClusterIdentityOutputReference {
	var returns KubernetesAutomaticClusterIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) IdentityInput() *KubernetesAutomaticClusterIdentity {
	var returns *KubernetesAutomaticClusterIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) KubeConfig() KubernetesAutomaticClusterKubeConfigList {
	var returns KubernetesAutomaticClusterKubeConfigList
	_jsii_.Get(
		j,
		"kubeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) KubeConfigRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubeConfigRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) NodeResourceGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeResourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) OidcIssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcIssuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) PortalFullyQualifiedDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portalFullyQualifiedDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) PrivateCluster() KubernetesAutomaticClusterPrivateClusterOutputReference {
	var returns KubernetesAutomaticClusterPrivateClusterOutputReference
	_jsii_.Get(
		j,
		"privateCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) PrivateClusterInput() *KubernetesAutomaticClusterPrivateCluster {
	var returns *KubernetesAutomaticClusterPrivateCluster
	_jsii_.Get(
		j,
		"privateClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) PrivateFullyQualifiedDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateFullyQualifiedDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ServiceMesh() KubernetesAutomaticClusterServiceMeshOutputReference {
	var returns KubernetesAutomaticClusterServiceMeshOutputReference
	_jsii_.Get(
		j,
		"serviceMesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) ServiceMeshInput() *KubernetesAutomaticClusterServiceMesh {
	var returns *KubernetesAutomaticClusterServiceMesh
	_jsii_.Get(
		j,
		"serviceMeshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) Timeouts() KubernetesAutomaticClusterTimeoutsOutputReference {
	var returns KubernetesAutomaticClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) WebAppRoutingIngress() KubernetesAutomaticClusterWebAppRoutingIngressOutputReference {
	var returns KubernetesAutomaticClusterWebAppRoutingIngressOutputReference
	_jsii_.Get(
		j,
		"webAppRoutingIngress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesAutomaticCluster) WebAppRoutingIngressInput() *KubernetesAutomaticClusterWebAppRoutingIngress {
	var returns *KubernetesAutomaticClusterWebAppRoutingIngress
	_jsii_.Get(
		j,
		"webAppRoutingIngressInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/kubernetes_automatic_cluster azurerm_kubernetes_automatic_cluster} Resource.
func NewKubernetesAutomaticCluster(scope constructs.Construct, id *string, config *KubernetesAutomaticClusterConfig) KubernetesAutomaticCluster {
	_init_.Initialize()

	if err := validateNewKubernetesAutomaticClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesAutomaticCluster{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.81.0/docs/resources/kubernetes_automatic_cluster azurerm_kubernetes_automatic_cluster} Resource.
func NewKubernetesAutomaticCluster_Override(k KubernetesAutomaticCluster, scope constructs.Construct, id *string, config *KubernetesAutomaticClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticCluster",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_KubernetesAutomaticCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a KubernetesAutomaticCluster resource upon running "cdktn plan <stack-name>".
func KubernetesAutomaticCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKubernetesAutomaticCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticCluster",
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
func KubernetesAutomaticCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesAutomaticCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesAutomaticCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesAutomaticCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesAutomaticCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesAutomaticCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesAutomaticCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.kubernetesAutomaticCluster.KubernetesAutomaticCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (k *jsiiProxy_KubernetesAutomaticCluster) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) PutApiServerAccess(value *KubernetesAutomaticClusterApiServerAccess) {
	if err := k.validatePutApiServerAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putApiServerAccess",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) PutHostedSystem(value *KubernetesAutomaticClusterHostedSystem) {
	if err := k.validatePutHostedSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putHostedSystem",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) PutIdentity(value *KubernetesAutomaticClusterIdentity) {
	if err := k.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIdentity",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) PutPrivateCluster(value *KubernetesAutomaticClusterPrivateCluster) {
	if err := k.validatePutPrivateClusterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putPrivateCluster",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) PutServiceMesh(value *KubernetesAutomaticClusterServiceMesh) {
	if err := k.validatePutServiceMeshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putServiceMesh",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) PutTimeouts(value *KubernetesAutomaticClusterTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) PutWebAppRoutingIngress(value *KubernetesAutomaticClusterWebAppRoutingIngress) {
	if err := k.validatePutWebAppRoutingIngressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putWebAppRoutingIngress",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetApiServerAccess() {
	_jsii_.InvokeVoid(
		k,
		"resetApiServerAccess",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetHostedSystem() {
	_jsii_.InvokeVoid(
		k,
		"resetHostedSystem",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetPrivateCluster() {
	_jsii_.InvokeVoid(
		k,
		"resetPrivateCluster",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetServiceMesh() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceMesh",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ResetWebAppRoutingIngress() {
	_jsii_.InvokeVoid(
		k,
		"resetWebAppRoutingIngress",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesAutomaticCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesAutomaticCluster) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

