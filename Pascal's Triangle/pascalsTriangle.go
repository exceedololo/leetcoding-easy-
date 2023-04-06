package Pascal_s_Triangle

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := range triangle {
		for j := range triangle[i] {
			if j == 0 || j == 1 {
				triangle[i][j] = 1
			} else {
				triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
			}
		}
	}
	return triangle
}
