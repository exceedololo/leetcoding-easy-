package Excel_Sheet_Column_Number

func titleToNumber(columnTitle string) int {
	number := 0

	for _, char := range columnTitle {
		number *= 26
		number += int(char-'A') + 1
	}
	return number
}
