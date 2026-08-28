// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-azurerm-go/azurerm/v18/linuxvirtualmachine/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/linux_virtual_machine azurerm_linux_virtual_machine}.
type LinuxVirtualMachine interface {
	cdktn.TerraformResource
	AdditionalCapabilities() LinuxVirtualMachineAdditionalCapabilitiesOutputReference
	AdditionalCapabilitiesInput() *LinuxVirtualMachineAdditionalCapabilities
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminSshKey() LinuxVirtualMachineAdminSshKeyList
	AdminSshKeyInput() interface{}
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	AllowExtensionOperations() interface{}
	SetAllowExtensionOperations(val interface{})
	AllowExtensionOperationsInput() interface{}
	AvailabilitySetId() *string
	SetAvailabilitySetId(val *string)
	AvailabilitySetIdInput() *string
	BootDiagnostics() LinuxVirtualMachineBootDiagnosticsOutputReference
	BootDiagnosticsInput() *LinuxVirtualMachineBootDiagnostics
	BypassPlatformSafetyChecksOnUserScheduleEnabled() interface{}
	SetBypassPlatformSafetyChecksOnUserScheduleEnabled(val interface{})
	BypassPlatformSafetyChecksOnUserScheduleEnabledInput() interface{}
	CapacityReservationGroupId() *string
	SetCapacityReservationGroupId(val *string)
	CapacityReservationGroupIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ComputerName() *string
	SetComputerName(val *string)
	ComputerNameInput() *string
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
	CustomData() *string
	SetCustomData(val *string)
	CustomDataInput() *string
	DedicatedHostGroupId() *string
	SetDedicatedHostGroupId(val *string)
	DedicatedHostGroupIdInput() *string
	DedicatedHostId() *string
	SetDedicatedHostId(val *string)
	DedicatedHostIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisablePasswordAuthentication() interface{}
	SetDisablePasswordAuthentication(val interface{})
	DisablePasswordAuthenticationInput() interface{}
	DiskControllerType() *string
	SetDiskControllerType(val *string)
	DiskControllerTypeInput() *string
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EncryptionAtHostEnabled() interface{}
	SetEncryptionAtHostEnabled(val interface{})
	EncryptionAtHostEnabledInput() interface{}
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
	ExtensionsTimeBudget() *string
	SetExtensionsTimeBudget(val *string)
	ExtensionsTimeBudgetInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GalleryApplication() LinuxVirtualMachineGalleryApplicationList
	GalleryApplicationInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() LinuxVirtualMachineIdentityOutputReference
	IdentityInput() *LinuxVirtualMachineIdentity
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
	MaxBidPrice() *float64
	SetMaxBidPrice(val *float64)
	MaxBidPriceInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterfaceIds() *[]*string
	SetNetworkInterfaceIds(val *[]*string)
	NetworkInterfaceIdsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	OsDisk() LinuxVirtualMachineOsDiskOutputReference
	OsDiskInput() *LinuxVirtualMachineOsDisk
	OsImageNotification() LinuxVirtualMachineOsImageNotificationOutputReference
	OsImageNotificationInput() *LinuxVirtualMachineOsImageNotification
	OsManagedDiskId() *string
	SetOsManagedDiskId(val *string)
	OsManagedDiskIdInput() *string
	PatchAssessmentMode() *string
	SetPatchAssessmentMode(val *string)
	PatchAssessmentModeInput() *string
	PatchMode() *string
	SetPatchMode(val *string)
	PatchModeInput() *string
	Plan() LinuxVirtualMachinePlanOutputReference
	PlanInput() *LinuxVirtualMachinePlan
	PlatformFaultDomain() *float64
	SetPlatformFaultDomain(val *float64)
	PlatformFaultDomainInput() *float64
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	PrivateIpAddress() *string
	PrivateIpAddresses() *[]*string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisionVmAgent() interface{}
	SetProvisionVmAgent(val interface{})
	ProvisionVmAgentInput() interface{}
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
	PublicIpAddress() *string
	PublicIpAddresses() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RebootSetting() *string
	SetRebootSetting(val *string)
	RebootSettingInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Secret() LinuxVirtualMachineSecretList
	SecretInput() interface{}
	SecureBootEnabled() interface{}
	SetSecureBootEnabled(val interface{})
	SecureBootEnabledInput() interface{}
	Size() *string
	SetSize(val *string)
	SizeInput() *string
	SourceImageId() *string
	SetSourceImageId(val *string)
	SourceImageIdInput() *string
	SourceImageReference() LinuxVirtualMachineSourceImageReferenceOutputReference
	SourceImageReferenceInput() *LinuxVirtualMachineSourceImageReference
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerminationNotification() LinuxVirtualMachineTerminationNotificationOutputReference
	TerminationNotificationInput() *LinuxVirtualMachineTerminationNotification
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LinuxVirtualMachineTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VirtualMachineId() *string
	VirtualMachineScaleSetId() *string
	SetVirtualMachineScaleSetId(val *string)
	VirtualMachineScaleSetIdInput() *string
	VmAgentPlatformUpdatesEnabled() cdktn.IResolvable
	VtpmEnabled() interface{}
	SetVtpmEnabled(val interface{})
	VtpmEnabledInput() interface{}
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutAdditionalCapabilities(value *LinuxVirtualMachineAdditionalCapabilities)
	PutAdminSshKey(value interface{})
	PutBootDiagnostics(value *LinuxVirtualMachineBootDiagnostics)
	PutGalleryApplication(value interface{})
	PutIdentity(value *LinuxVirtualMachineIdentity)
	PutOsDisk(value *LinuxVirtualMachineOsDisk)
	PutOsImageNotification(value *LinuxVirtualMachineOsImageNotification)
	PutPlan(value *LinuxVirtualMachinePlan)
	PutSecret(value interface{})
	PutSourceImageReference(value *LinuxVirtualMachineSourceImageReference)
	PutTerminationNotification(value *LinuxVirtualMachineTerminationNotification)
	PutTimeouts(value *LinuxVirtualMachineTimeouts)
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
	ResetAdditionalCapabilities()
	ResetAdminPassword()
	ResetAdminSshKey()
	ResetAdminUsername()
	ResetAllowExtensionOperations()
	ResetAvailabilitySetId()
	ResetBootDiagnostics()
	ResetBypassPlatformSafetyChecksOnUserScheduleEnabled()
	ResetCapacityReservationGroupId()
	ResetComputerName()
	ResetCustomData()
	ResetDedicatedHostGroupId()
	ResetDedicatedHostId()
	ResetDisablePasswordAuthentication()
	ResetDiskControllerType()
	ResetEdgeZone()
	ResetEncryptionAtHostEnabled()
	ResetEvictionPolicy()
	ResetExtensionsTimeBudget()
	ResetGalleryApplication()
	ResetId()
	ResetIdentity()
	ResetLicenseType()
	ResetMaxBidPrice()
	ResetOsImageNotification()
	ResetOsManagedDiskId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPatchAssessmentMode()
	ResetPatchMode()
	ResetPlan()
	ResetPlatformFaultDomain()
	ResetPriority()
	ResetProvisionVmAgent()
	ResetProximityPlacementGroupId()
	ResetRebootSetting()
	ResetSecret()
	ResetSecureBootEnabled()
	ResetSourceImageId()
	ResetSourceImageReference()
	ResetTags()
	ResetTerminationNotification()
	ResetTimeouts()
	ResetUserData()
	ResetVirtualMachineScaleSetId()
	ResetVtpmEnabled()
	ResetZone()
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

// The jsii proxy struct for LinuxVirtualMachine
type jsiiProxy_LinuxVirtualMachine struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_LinuxVirtualMachine) AdditionalCapabilities() LinuxVirtualMachineAdditionalCapabilitiesOutputReference {
	var returns LinuxVirtualMachineAdditionalCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"additionalCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AdditionalCapabilitiesInput() *LinuxVirtualMachineAdditionalCapabilities {
	var returns *LinuxVirtualMachineAdditionalCapabilities
	_jsii_.Get(
		j,
		"additionalCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AdminSshKey() LinuxVirtualMachineAdminSshKeyList {
	var returns LinuxVirtualMachineAdminSshKeyList
	_jsii_.Get(
		j,
		"adminSshKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AdminSshKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminSshKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AllowExtensionOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExtensionOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AllowExtensionOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExtensionOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AvailabilitySetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) AvailabilitySetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) BootDiagnostics() LinuxVirtualMachineBootDiagnosticsOutputReference {
	var returns LinuxVirtualMachineBootDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) BootDiagnosticsInput() *LinuxVirtualMachineBootDiagnostics {
	var returns *LinuxVirtualMachineBootDiagnostics
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) BypassPlatformSafetyChecksOnUserScheduleEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassPlatformSafetyChecksOnUserScheduleEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) BypassPlatformSafetyChecksOnUserScheduleEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassPlatformSafetyChecksOnUserScheduleEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) CapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) CapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ComputerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DedicatedHostGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DedicatedHostGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DedicatedHostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DedicatedHostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DisablePasswordAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DisablePasswordAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DiskControllerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskControllerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) DiskControllerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskControllerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) EncryptionAtHostEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) EncryptionAtHostEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ExtensionsTimeBudget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ExtensionsTimeBudgetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudgetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) GalleryApplication() LinuxVirtualMachineGalleryApplicationList {
	var returns LinuxVirtualMachineGalleryApplicationList
	_jsii_.Get(
		j,
		"galleryApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) GalleryApplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Identity() LinuxVirtualMachineIdentityOutputReference {
	var returns LinuxVirtualMachineIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) IdentityInput() *LinuxVirtualMachineIdentity {
	var returns *LinuxVirtualMachineIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) MaxBidPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) MaxBidPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) NetworkInterfaceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) OsDisk() LinuxVirtualMachineOsDiskOutputReference {
	var returns LinuxVirtualMachineOsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) OsDiskInput() *LinuxVirtualMachineOsDisk {
	var returns *LinuxVirtualMachineOsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) OsImageNotification() LinuxVirtualMachineOsImageNotificationOutputReference {
	var returns LinuxVirtualMachineOsImageNotificationOutputReference
	_jsii_.Get(
		j,
		"osImageNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) OsImageNotificationInput() *LinuxVirtualMachineOsImageNotification {
	var returns *LinuxVirtualMachineOsImageNotification
	_jsii_.Get(
		j,
		"osImageNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) OsManagedDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osManagedDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) OsManagedDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osManagedDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PatchAssessmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PatchAssessmentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PatchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PatchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Plan() LinuxVirtualMachinePlanOutputReference {
	var returns LinuxVirtualMachinePlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PlanInput() *LinuxVirtualMachinePlan {
	var returns *LinuxVirtualMachinePlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PlatformFaultDomain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PlatformFaultDomainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) RebootSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) RebootSettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Secret() LinuxVirtualMachineSecretList {
	var returns LinuxVirtualMachineSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SourceImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SourceImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SourceImageReference() LinuxVirtualMachineSourceImageReferenceOutputReference {
	var returns LinuxVirtualMachineSourceImageReferenceOutputReference
	_jsii_.Get(
		j,
		"sourceImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) SourceImageReferenceInput() *LinuxVirtualMachineSourceImageReference {
	var returns *LinuxVirtualMachineSourceImageReference
	_jsii_.Get(
		j,
		"sourceImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) TerminationNotification() LinuxVirtualMachineTerminationNotificationOutputReference {
	var returns LinuxVirtualMachineTerminationNotificationOutputReference
	_jsii_.Get(
		j,
		"terminationNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) TerminationNotificationInput() *LinuxVirtualMachineTerminationNotification {
	var returns *LinuxVirtualMachineTerminationNotification
	_jsii_.Get(
		j,
		"terminationNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Timeouts() LinuxVirtualMachineTimeoutsOutputReference {
	var returns LinuxVirtualMachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) VirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) VirtualMachineScaleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) VirtualMachineScaleSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) VmAgentPlatformUpdatesEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"vmAgentPlatformUpdatesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) VtpmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) VtpmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachine) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/linux_virtual_machine azurerm_linux_virtual_machine} Resource.
func NewLinuxVirtualMachine(scope constructs.Construct, id *string, config *LinuxVirtualMachineConfig) LinuxVirtualMachine {
	_init_.Initialize()

	if err := validateNewLinuxVirtualMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxVirtualMachine{}

	_jsii_.Create(
		"@cdktn/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/5.3.0/docs/resources/linux_virtual_machine azurerm_linux_virtual_machine} Resource.
func NewLinuxVirtualMachine_Override(l LinuxVirtualMachine, scope constructs.Construct, id *string, config *LinuxVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachine",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetAllowExtensionOperations(val interface{}) {
	if err := j.validateSetAllowExtensionOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowExtensionOperations",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetAvailabilitySetId(val *string) {
	if err := j.validateSetAvailabilitySetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilitySetId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetBypassPlatformSafetyChecksOnUserScheduleEnabled(val interface{}) {
	if err := j.validateSetBypassPlatformSafetyChecksOnUserScheduleEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassPlatformSafetyChecksOnUserScheduleEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetCapacityReservationGroupId(val *string) {
	if err := j.validateSetCapacityReservationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetComputerName(val *string) {
	if err := j.validateSetComputerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerName",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetDedicatedHostGroupId(val *string) {
	if err := j.validateSetDedicatedHostGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedHostGroupId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetDedicatedHostId(val *string) {
	if err := j.validateSetDedicatedHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedHostId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetDisablePasswordAuthentication(val interface{}) {
	if err := j.validateSetDisablePasswordAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePasswordAuthentication",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetDiskControllerType(val *string) {
	if err := j.validateSetDiskControllerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskControllerType",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetEncryptionAtHostEnabled(val interface{}) {
	if err := j.validateSetEncryptionAtHostEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtHostEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetExtensionsTimeBudget(val *string) {
	if err := j.validateSetExtensionsTimeBudgetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionsTimeBudget",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetMaxBidPrice(val *float64) {
	if err := j.validateSetMaxBidPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBidPrice",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetNetworkInterfaceIds(val *[]*string) {
	if err := j.validateSetNetworkInterfaceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceIds",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetOsManagedDiskId(val *string) {
	if err := j.validateSetOsManagedDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osManagedDiskId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetPatchAssessmentMode(val *string) {
	if err := j.validateSetPatchAssessmentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchAssessmentMode",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetPatchMode(val *string) {
	if err := j.validateSetPatchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchMode",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetPlatformFaultDomain(val *float64) {
	if err := j.validateSetPlatformFaultDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformFaultDomain",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetRebootSetting(val *string) {
	if err := j.validateSetRebootSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rebootSetting",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetSecureBootEnabled(val interface{}) {
	if err := j.validateSetSecureBootEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetSize(val *string) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetSourceImageId(val *string) {
	if err := j.validateSetSourceImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImageId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetVirtualMachineScaleSetId(val *string) {
	if err := j.validateSetVirtualMachineScaleSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineScaleSetId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetVtpmEnabled(val interface{}) {
	if err := j.validateSetVtpmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vtpmEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachine)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTN code for importing a LinuxVirtualMachine resource upon running "cdktn plan <stack-name>".
func LinuxVirtualMachine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateLinuxVirtualMachine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachine",
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
func LinuxVirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxVirtualMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxVirtualMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxVirtualMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxVirtualMachine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxVirtualMachine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LinuxVirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxVirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxVirtualMachine) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LinuxVirtualMachine) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := l.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutAdditionalCapabilities(value *LinuxVirtualMachineAdditionalCapabilities) {
	if err := l.validatePutAdditionalCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdditionalCapabilities",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutAdminSshKey(value interface{}) {
	if err := l.validatePutAdminSshKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdminSshKey",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutBootDiagnostics(value *LinuxVirtualMachineBootDiagnostics) {
	if err := l.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutGalleryApplication(value interface{}) {
	if err := l.validatePutGalleryApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGalleryApplication",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutIdentity(value *LinuxVirtualMachineIdentity) {
	if err := l.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutOsDisk(value *LinuxVirtualMachineOsDisk) {
	if err := l.validatePutOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutOsImageNotification(value *LinuxVirtualMachineOsImageNotification) {
	if err := l.validatePutOsImageNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOsImageNotification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutPlan(value *LinuxVirtualMachinePlan) {
	if err := l.validatePutPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPlan",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutSecret(value interface{}) {
	if err := l.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSecret",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutSourceImageReference(value *LinuxVirtualMachineSourceImageReference) {
	if err := l.validatePutSourceImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSourceImageReference",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutTerminationNotification(value *LinuxVirtualMachineTerminationNotification) {
	if err := l.validatePutTerminationNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTerminationNotification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) PutTimeouts(value *LinuxVirtualMachineTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := l.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetAdditionalCapabilities() {
	_jsii_.InvokeVoid(
		l,
		"resetAdditionalCapabilities",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetAdminSshKey() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminSshKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetAdminUsername() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminUsername",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetAllowExtensionOperations() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowExtensionOperations",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetAvailabilitySetId() {
	_jsii_.InvokeVoid(
		l,
		"resetAvailabilitySetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		l,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetBypassPlatformSafetyChecksOnUserScheduleEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetBypassPlatformSafetyChecksOnUserScheduleEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		l,
		"resetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetComputerName() {
	_jsii_.InvokeVoid(
		l,
		"resetComputerName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetCustomData() {
	_jsii_.InvokeVoid(
		l,
		"resetCustomData",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetDedicatedHostGroupId() {
	_jsii_.InvokeVoid(
		l,
		"resetDedicatedHostGroupId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetDedicatedHostId() {
	_jsii_.InvokeVoid(
		l,
		"resetDedicatedHostId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetDisablePasswordAuthentication() {
	_jsii_.InvokeVoid(
		l,
		"resetDisablePasswordAuthentication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetDiskControllerType() {
	_jsii_.InvokeVoid(
		l,
		"resetDiskControllerType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		l,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetEncryptionAtHostEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEncryptionAtHostEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetExtensionsTimeBudget() {
	_jsii_.InvokeVoid(
		l,
		"resetExtensionsTimeBudget",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetGalleryApplication() {
	_jsii_.InvokeVoid(
		l,
		"resetGalleryApplication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetLicenseType() {
	_jsii_.InvokeVoid(
		l,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetMaxBidPrice() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxBidPrice",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetOsImageNotification() {
	_jsii_.InvokeVoid(
		l,
		"resetOsImageNotification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetOsManagedDiskId() {
	_jsii_.InvokeVoid(
		l,
		"resetOsManagedDiskId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetPatchAssessmentMode() {
	_jsii_.InvokeVoid(
		l,
		"resetPatchAssessmentMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetPatchMode() {
	_jsii_.InvokeVoid(
		l,
		"resetPatchMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetPlan() {
	_jsii_.InvokeVoid(
		l,
		"resetPlan",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetPlatformFaultDomain() {
	_jsii_.InvokeVoid(
		l,
		"resetPlatformFaultDomain",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetPriority() {
	_jsii_.InvokeVoid(
		l,
		"resetPriority",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		l,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		l,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetRebootSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetRebootSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetSecret() {
	_jsii_.InvokeVoid(
		l,
		"resetSecret",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetSecureBootEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetSecureBootEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetSourceImageId() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceImageId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetSourceImageReference() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceImageReference",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetTerminationNotification() {
	_jsii_.InvokeVoid(
		l,
		"resetTerminationNotification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetUserData() {
	_jsii_.InvokeVoid(
		l,
		"resetUserData",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetVirtualMachineScaleSetId() {
	_jsii_.InvokeVoid(
		l,
		"resetVirtualMachineScaleSetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetVtpmEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetVtpmEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) ResetZone() {
	_jsii_.InvokeVoid(
		l,
		"resetZone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachine) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

