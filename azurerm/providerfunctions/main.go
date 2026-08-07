// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package providerfunctions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/provider-azurerm.providerFunctions.AzurermProviderFunctions",
		reflect.TypeOf((*AzurermProviderFunctions)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "normaliseResourceId", GoMethod: "NormaliseResourceId"},
			_jsii_.MemberMethod{JsiiMethod: "parseResourceId", GoMethod: "ParseResourceId"},
		},
		func() interface{} {
			return &jsiiProxy_AzurermProviderFunctions{}
		},
	)
}
