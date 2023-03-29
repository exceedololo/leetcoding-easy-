package Add_Binary

import "strings"

func addBinary(a string, b string) string {
	result := ""
	if len(a) > len(b) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	} else {
		a = strings.Repeat("0", len(b)-len(a)) + a
	}
	carry := "0"
	for i := len(a) - 1; i >= 0; i-- {
		if a[i:i+1] == "0" && b[i:i+1] == "0" {
			if carry == "1" {
				result = "1" + result
				carry = "0"
			} else {
				result = "0" + result
			}
		} else if a[i:i+1] == "1" && b[i:i+1] == "1" {
			if carry == "0" {
				result = "0" + result
				carry = "1"
			} else {
				result = "1" + result
			}
		} else {
			if carry == "0" {
				result = "1" + result
			} else {
				result = "0" + result
			}
		}
	}
	if carry == "1" {
		result = "1" + result
	}
	return result
}
