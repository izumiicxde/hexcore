package utils

import "hexcore/types"

var Subjects = CalculateMaxClasses(85)

func CalculateMaxClasses(workingDays int) []types.Subject {
	weeks := workingDays / 6

	subjects := []types.Subject{
		{Name: "ADA", MaxClasses: (weeks * 4) - 3},
		{Name: "SE", MaxClasses: (weeks * 4) - 3},
		{Name: "IT", MaxClasses: (weeks * 4) - 3},
		{Name: "LANG", MaxClasses: weeks * 4},
		{Name: "ENG", MaxClasses: weeks * 4},
		{Name: "OE", MaxClasses: weeks * 2},
		{Name: "IC", MaxClasses: weeks * 2},
		{Name: "ADA Lab", MaxClasses: weeks * 1},
		{Name: "IT Lab", MaxClasses: weeks * 1},
	}

	return subjects
}
