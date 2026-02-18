// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package storageshare

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageShareAclAccessPolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_StorageShareAclAccessPolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageShareAclAccessPolicyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageShareAclAccessPolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageShareAclAccessPolicyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

