// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devtestschedule


type DevTestScheduleWeeklyRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/dev_test_schedule#time DevTestSchedule#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.75.0/docs/resources/dev_test_schedule#week_days DevTestSchedule#week_days}.
	WeekDays *[]*string `field:"optional" json:"weekDays" yaml:"weekDays"`
}

