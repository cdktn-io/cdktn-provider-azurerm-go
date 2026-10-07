// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataazurermoracleautonomousdatabasecrossregiondisasterrecovery

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/dataazurermoracleautonomousdatabasecrossregiondisasterrecovery/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/data-sources/oracle_autonomous_database_cross_region_disaster_recovery azurerm_oracle_autonomous_database_cross_region_disaster_recovery}.
type DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery interface {
	cdktn.TerraformDataSource
	ActualUsedDataStorageSizeInTb() *float64
	AllocatedStorageSizeInTb() *float64
	AutoScalingEnabled() cdktn.IResolvable
	AutoScalingForStorageEnabled() cdktn.IResolvable
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodInDays() *float64
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CharacterSet() *string
	ComputeCount() *float64
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CpuCoreCount() *float64
	CustomerContacts() *[]*string
	DatabaseType() *string
	DatabaseVersion() *string
	DatabaseWorkload() *string
	DataStorageSizeInGb() *float64
	DataStorageSizeInTb() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	FailedDataRecoveryInSeconds() *float64
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
	InMemoryAreaInGb() *float64
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LifecycleDetails() *string
	LocalAdgAutoFailoverMaximumDataLossLimitInSeconds() *float64
	LocalDataGuardEnabled() cdktn.IResolvable
	Location() *string
	MemoryPerOracleComputeUnitInGb() *float64
	MtlsConnectionRequired() cdktn.IResolvable
	Name() *string
	SetName(val *string)
	NameInput() *string
	NationalCharacterSet() *string
	NextLongTermBackupTimestampInUtc() *string
	// The tree node.
	Node() constructs.Node
	Ocid() *string
	OciUrl() *string
	PeerDatabaseIds() *[]*string
	Preview() cdktn.IResolvable
	PreviewVersionWithServiceTermsAccepted() cdktn.IResolvable
	PrivateEndpointIp() *string
	PrivateEndpointLabel() *string
	PrivateEndpointUrl() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProvisionableCpus() *[]*float64
	// Experimental.
	RawOverrides() interface{}
	RemoteDataGuardEnabled() cdktn.IResolvable
	RemoteDisasterRecoveryType() *string
	ReplicateAutomaticBackupsEnabled() cdktn.IResolvable
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	ServiceConsoleUrl() *string
	SourceAutonomousDatabaseId() *string
	SourceLocation() *string
	SourceOcid() *string
	SourceType() *string
	SqlWebDeveloperUrl() *string
	SubnetId() *string
	Tags() cdktn.StringMap
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeCreatedInUtc() *string
	TimeDataGuardRoleChangedInUtc() *string
	TimeDeletionOfFreeAutonomousDatabaseInUtc() *string
	TimeLocalDataGuardEnabledInUtc() *string
	TimeMaintenanceBeginInUtc() *string
	TimeMaintenanceEndInUtc() *string
	TimeOfLastFailoverInUtc() *string
	TimeOfLastRefreshInUtc() *string
	TimeOfLastRefreshPointInUtc() *string
	TimeOfLastSwitchoverInUtc() *string
	Timeouts() DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeReclamationOfFreeAutonomousDatabaseInUtc() *string
	UsedDataStorageSizeInGb() *float64
	UsedDataStorageSizeInTb() *float64
	VirtualNetworkId() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryTimeouts)
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery
type jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ActualUsedDataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) AllocatedStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) AutoScalingEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"autoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) AutoScalingForStorageEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"autoScalingForStorageEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) BackupRetentionPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) CustomerContacts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) DatabaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) DatabaseWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) DataStorageSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) DataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) FailedDataRecoveryInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failedDataRecoveryInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) InMemoryAreaInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"inMemoryAreaInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) LocalAdgAutoFailoverMaximumDataLossLimitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaximumDataLossLimitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) LocalDataGuardEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"localDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) MemoryPerOracleComputeUnitInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) MtlsConnectionRequired() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) NationalCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nationalCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) NextLongTermBackupTimestampInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTimestampInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) PeerDatabaseIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDatabaseIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Preview() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) PreviewVersionWithServiceTermsAccepted() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"previewVersionWithServiceTermsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) PrivateEndpointUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ProvisionableCpus() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"provisionableCpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) RemoteDataGuardEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"remoteDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) RemoteDisasterRecoveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDisasterRecoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ReplicateAutomaticBackupsEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"replicateAutomaticBackupsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ServiceConsoleUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceConsoleUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SourceAutonomousDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAutonomousDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SourceLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SourceOcid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceOcid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Tags() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeCreatedInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreatedInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeDataGuardRoleChangedInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDataGuardRoleChangedInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeDeletionOfFreeAutonomousDatabaseInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeDeletionOfFreeAutonomousDatabaseInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeLocalDataGuardEnabledInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeLocalDataGuardEnabledInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeMaintenanceBeginInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceBeginInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeMaintenanceEndInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeMaintenanceEndInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeOfLastFailoverInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastFailoverInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeOfLastRefreshInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefreshInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeOfLastRefreshPointInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastRefreshPointInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeOfLastSwitchoverInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeOfLastSwitchoverInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) Timeouts() DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryTimeoutsOutputReference {
	var returns DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) TimeReclamationOfFreeAutonomousDatabaseInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeReclamationOfFreeAutonomousDatabaseInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) UsedDataStorageSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) UsedDataStorageSizeInTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeInTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/data-sources/oracle_autonomous_database_cross_region_disaster_recovery azurerm_oracle_autonomous_database_cross_region_disaster_recovery} Data Source.
func NewDataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery(scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryConfig) DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery {
	_init_.Initialize()

	if err := validateNewDataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery.DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.8.0/docs/data-sources/oracle_autonomous_database_cross_region_disaster_recovery azurerm_oracle_autonomous_database_cross_region_disaster_recovery} Data Source.
func NewDataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_Override(d DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery, scope constructs.Construct, id *string, config *DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.dataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery.DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTN code for importing a DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery resource upon running "cdktn plan <stack-name>".
func DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery.DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery",
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
func DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery.DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery.DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.dataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery.DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.dataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery.DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) PutTimeouts(value *DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecoveryTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermOracleAutonomousDatabaseCrossRegionDisasterRecovery) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

