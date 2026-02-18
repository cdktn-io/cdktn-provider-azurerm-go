// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dnscaarecord

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsCaaRecordRecordList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DnsCaaRecordRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsCaaRecordRecordList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsCaaRecordRecordList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsCaaRecordRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsCaaRecordRecordList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsCaaRecordRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsCaaRecordRecordListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

