// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package signalrservice

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SignalrServiceUpstreamEndpointList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SignalrServiceUpstreamEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SignalrServiceUpstreamEndpointList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SignalrServiceUpstreamEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSignalrServiceUpstreamEndpointListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

