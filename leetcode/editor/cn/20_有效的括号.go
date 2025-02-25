package main

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s) %2 !=0 {
		return false
	}
	stack := make([]byte, 0)
	for _, c := range s {
		if c == '[' || c == '(' || c == '{' {
			stack = append(stack, byte(c))
		} else {
			switch c {
			case ']':
				if len(stack) <=0 || stack[len(stack)-1] != '[' {
					return false
				}
			case ')':
				if  len(stack) <=0  || stack[len(stack)-1] != '(' {
					return false
				}
			case '}':
				if  len(stack) <=0 || stack[len(stack)-1] != '{' {
					return false
				}
			default:
				return false

			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true

}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	isValid("()")
}
