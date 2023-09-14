func isValid(s string) bool {
    var keys = map[uint8]uint8{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack = make([]uint8, 0)

	for i := 0; i < len(s); i++ {
		char := s[i]
		if val, valid := keys[char]; valid {
			stack = append([]uint8{val}, stack...)
		} else {
            if len(stack) == 0 || stack[0] != char {
				return false
			}
			stack = stack[1:]

		}

	}
    
	return (len(stack) == 0)
}