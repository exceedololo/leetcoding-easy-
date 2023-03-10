package Longest_Common_Prefix

func longestCommonPrefix(strs []string) string {
	/*if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for k := 0; k < len(prefix); k++ {
			if k > len(strs[i]) || prefix[k] != strs[i][k] {
				return prefix[:k]
			}
		}
	}
	return prefix*/
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for k := 0; k < len(prefix); k++ {
			if k >= len(strs[i]) || prefix[k] != strs[i][k] {
				prefix = prefix[:k]
				break
			}
		}
	}
	return prefix
}
