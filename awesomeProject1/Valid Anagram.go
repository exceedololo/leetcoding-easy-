package awesomeProject1

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	frequency := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		frequency[s[i]]++
		frequency[t[i]]--
	}
	for _, count := range frequency {
		if count != 0 {
			return false
		}
	}
	return true
}
