func lenghtOfLastWord(s string) int{
	dlina := 0
	probel := false
	for i := len(s)-1;i >= 0; i--{
		if s[i] != ' '{
			dlina++
			probel = true
		}else if probel{
			break
		}
	}
	return dlina
}
