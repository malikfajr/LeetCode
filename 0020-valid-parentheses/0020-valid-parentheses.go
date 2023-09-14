func isValid(s string) bool {
    if len(s)%2 == 1 {
		return false
	}
    var keys = map[uint8]uint8{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack = make([]uint8, len(s))

	for i := 0; i < len(s); i++ {
		char := s[i]
		if val, valid := keys[char]; valid {
			stack = append([]uint8{val}, stack...)
		} else {
			top := stack[0]
			if top != char {
				return false
			}
			stack = stack[1:]

		}

	}
    
    if stack[0] != 0 {
		return false
	}

	return true
}