package Date

import (
	l "SecondTesting/Leap"
	s "strconv"
)

func Date(inputYear int, inputNumber int) string {
	if (inputNumber > 366) || (inputNumber <= 0) {
		return "The input number cannot be bigger than 366 or less than one!"
	}
	var day string
	var month string
	var year string = s.Itoa(inputYear)
	if (!l.IsLeap(inputYear)) && (inputNumber == 366) {
		return "The entered number cannot be 366 when the year is not a leap year!"
	} else if (!l.IsLeap(inputYear)) && (inputNumber != 366) {
		if inputNumber <= 31 {
			month = "January"
			day = s.Itoa(inputNumber)
		} else if inputNumber <= 59 {
			month = "February"
			day = s.Itoa(inputNumber - 31)
		} else if inputNumber <= 90 {
			month = "March"
			day = s.Itoa(inputNumber - 59)
		} else if inputNumber <= 120 {
			month = "April"
			day = s.Itoa(inputNumber - 90)
		} else if inputNumber <= 151 {
			month = "May"
			day = s.Itoa(inputNumber - 120)
		} else if inputNumber <= 181 {
			month = "June"
			day = s.Itoa(inputNumber - 151)
		} else if inputNumber <= 212 {
			month = "July"
			day = s.Itoa(inputNumber - 181)
		} else if inputNumber <= 243 {
			month = "August"
			day = s.Itoa(inputNumber - 212)
		} else if inputNumber <= 273 {
			month = "September"
			day = s.Itoa(inputNumber - 243)
		} else if inputNumber <= 304 {
			month = "October"
			day = s.Itoa(inputNumber - 273)
		} else if inputNumber <= 334 {
			month = "November"
			day = s.Itoa(inputNumber - 304)
		} else {
			month = "December"
			day = s.Itoa(inputNumber - 334)
		}
		var result string = day + "/" + month + "/" + year
		return result
	} else {
		if inputNumber <= 31 {
			month = "January"
			day = s.Itoa(inputNumber)
		} else if inputNumber <= 60 {
			month = "February"
			day = s.Itoa(inputNumber - 31)
		} else if inputNumber <= 91 {
			month = "March"
			day = s.Itoa(inputNumber - 60)
		} else if inputNumber <= 121 {
			month = "April"
			day = s.Itoa(inputNumber - 91)
		} else if inputNumber <= 152 {
			month = "May"
			day = s.Itoa(inputNumber - 121)
		} else if inputNumber <= 182 {
			month = "June"
			day = s.Itoa(inputNumber - 152)
		} else if inputNumber <= 213 {
			month = "July"
			day = s.Itoa(inputNumber - 182)
		} else if inputNumber <= 244 {
			month = "August"
			day = s.Itoa(inputNumber - 213)
		} else if inputNumber <= 274 {
			month = "September"
			day = s.Itoa(inputNumber - 244)
		} else if inputNumber <= 305 {
			month = "October"
			day = s.Itoa(inputNumber - 274)
		} else if inputNumber <= 335 {
			month = "November"
			day = s.Itoa(inputNumber - 305)
		} else {
			month = "December"
			day = s.Itoa(inputNumber - 335)
		}
		var result string = day + "/" + month + "/" + year
		return result
	}
}
