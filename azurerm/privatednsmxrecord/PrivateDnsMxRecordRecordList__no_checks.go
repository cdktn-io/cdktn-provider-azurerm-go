// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package privatednsmxrecord

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateDnsMxRecordRecordList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsMxRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivateDnsMxRecordRecordList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivateDnsMxRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivateDnsMxRecordRecordListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

