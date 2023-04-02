package Valid_Parentheses

func isValid(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, "(")
		} else if s[i] == '[' {
			stack = append(stack, "[")
		} else if s[i] == '{' {
			stack = append(stack, "{")
		} else if s[i] == '}' && len(stack) > 0 && stack[len(stack)-1] == "{" {
			stack = stack[:len(stack)-1]
		} else if s[i] == ')' && len(stack) > 0 && stack[len(stack)-1] == "(" {
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && len(stack) > 0 && stack[len(stack)-1] == "[" {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
