func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack) - len(needle)+1;i++{
		found := true
		for k:=0; k< len(needle);k++{
			if haystack[i+k] != needle[k]{
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}