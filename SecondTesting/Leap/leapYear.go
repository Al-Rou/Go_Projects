package Leap

func IsLeap(year int) bool {
	if (year%100 != 0) && (year%4 == 0) {
		return true
	} else if year%400 == 0 {
		return true
	}
	return false
}
