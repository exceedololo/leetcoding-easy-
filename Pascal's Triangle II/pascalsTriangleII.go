package Pascal_s_Triangle_II

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		tmp := row[i-1]
		tmp = tmp * (rowIndex - i + 1) / i
		row[i] = tmp
	}
	return row
}
