// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchaccount

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchAccountEncryptionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BatchAccountEncryptionList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchAccountEncryptionList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchAccountEncryptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchAccountEncryptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchAccountEncryptionList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchAccountEncryptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchAccountEncryptionListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

