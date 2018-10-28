package raindrops

import "strconv"

// Convert converts a number to a string, the contents of which depend on the number's factors.
func Convert(num int) string {
	var result string
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			switch i {
			case 3:
				result = result + "Pling"
			case 5:
				result = result + "Plang"
			case 7:
				result = result + "Plong"
			}
		}
	}
	if result == "" {
		return strconv.Itoa(num)
	}
	return result
}
