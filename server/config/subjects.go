package config

import "hexcore/types"

var Subjects = []types.Subject{
	{Name: "ADA", MaxClasses: 59},
	{Name: "IT", MaxClasses: 59},
	{Name: "SE", MaxClasses: 59},
	{Name: "IC", MaxClasses: 32},
	{Name: "LANG", MaxClasses: 64},
	{Name: "ENG", MaxClasses: 64},
	{Name: "OE", MaxClasses: 32},
	{Name: "ADA Lab", MaxClasses: 16},
	{Name: "IT Lab", MaxClasses: 16},
}

var Schedules = []types.Schedule{
	{SubjectName: "ADA", Day: "Wednesday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "ADA", Day: "Thursday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "ADA", Day: "Friday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "ADA", Day: "Saturday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "IT", Day: "Monday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "IT", Day: "Wednesday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
	{SubjectName: "IT", Day: "Friday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "IT", Day: "Saturday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "SE", Day: "Monday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "SE", Day: "Tuesday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "SE", Day: "Wednesday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "SE", Day: "Saturday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "IC", Day: "Wednesday", StartTime: "9:30 AM", EndTime: "10:30 AM"},
	{SubjectName: "IC", Day: "Thursday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "LANG", Day: "Monday", StartTime: "10:45 AM", EndTime: "11:45 AM"},
	{SubjectName: "LANG", Day: "Wednesday", StartTime: "10:45 AM", EndTime: "11:45 PM"},
	{SubjectName: "LANG", Day: "Thursday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
	{SubjectName: "LANG", Day: "Friday", StartTime: "11:45 AM", EndTime: "12:45 PM"},
	{SubjectName: "ENG", Day: "Monday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "ENG", Day: "Tuesday", StartTime: "2:30 PM", EndTime: "3:30 PM"},
	{SubjectName: "ENG", Day: "Wednesday", StartTime: "2:30 PM", EndTime: "3:30 PM"},
	{SubjectName: "ENG", Day: "Friday", StartTime: "8:30 AM", EndTime: "9:30 AM"},
	{SubjectName: "OE", Day: "Tuesday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "OE", Day: "Thursday", StartTime: "1:30 PM", EndTime: "2:30 PM"},
	{SubjectName: "ADA Lab", Day: "Tuesday", StartTime: "9:30 AM", EndTime: "11:45 AM"},
	{SubjectName: "IT Lab", Day: "Friday", StartTime: "1:30 PM", EndTime: "4:30 PM"},
}
