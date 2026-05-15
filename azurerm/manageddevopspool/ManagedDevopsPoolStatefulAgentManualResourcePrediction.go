// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package manageddevopspool


type ManagedDevopsPoolStatefulAgentManualResourcePrediction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#all_week_schedule ManagedDevopsPool#all_week_schedule}.
	AllWeekSchedule *float64 `field:"optional" json:"allWeekSchedule" yaml:"allWeekSchedule"`
	// friday_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#friday_schedule ManagedDevopsPool#friday_schedule}
	FridaySchedule interface{} `field:"optional" json:"fridaySchedule" yaml:"fridaySchedule"`
	// monday_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#monday_schedule ManagedDevopsPool#monday_schedule}
	MondaySchedule interface{} `field:"optional" json:"mondaySchedule" yaml:"mondaySchedule"`
	// saturday_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#saturday_schedule ManagedDevopsPool#saturday_schedule}
	SaturdaySchedule interface{} `field:"optional" json:"saturdaySchedule" yaml:"saturdaySchedule"`
	// sunday_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#sunday_schedule ManagedDevopsPool#sunday_schedule}
	SundaySchedule interface{} `field:"optional" json:"sundaySchedule" yaml:"sundaySchedule"`
	// thursday_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#thursday_schedule ManagedDevopsPool#thursday_schedule}
	ThursdaySchedule interface{} `field:"optional" json:"thursdaySchedule" yaml:"thursdaySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#time_zone_name ManagedDevopsPool#time_zone_name}.
	TimeZoneName *string `field:"optional" json:"timeZoneName" yaml:"timeZoneName"`
	// tuesday_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#tuesday_schedule ManagedDevopsPool#tuesday_schedule}
	TuesdaySchedule interface{} `field:"optional" json:"tuesdaySchedule" yaml:"tuesdaySchedule"`
	// wednesday_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.73.0/docs/resources/managed_devops_pool#wednesday_schedule ManagedDevopsPool#wednesday_schedule}
	WednesdaySchedule interface{} `field:"optional" json:"wednesdaySchedule" yaml:"wednesdaySchedule"`
}

