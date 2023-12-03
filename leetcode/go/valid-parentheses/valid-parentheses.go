func isValid(s string) bool {
	var stack []byte

	for i := range s {
		b := s[i]
		if b == '{' || b == '[' || b == '(' {
			stack = append(stack, b)
		} else {
            if len(stack) == 0 {
                return false
            }
            
			latest := stack[len(stack)-1]

			if isPair[latest] != b {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

var isPair = map[byte]byte{
	'{': '}',
	'[': ']',
	'(': ')',
}