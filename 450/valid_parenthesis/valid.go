package main

func isValid(s string) bool {
	var stack []rune
	if len(s)%2 != 0 {
		return false
	}
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else {
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				if (ch == ')' && top == '(') || (ch == ']' && top == '[' || (ch == '}' && top == '{')) {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
