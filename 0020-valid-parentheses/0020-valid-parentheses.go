func isValid(s string) bool {
    	var keys = map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	var stack = make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {

		if _, valid := keys[s[i]]; valid {
			stack = append(stack, keys[s[i]])
			fmt.Println(stack, keys[s[i]])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]

		}

	}
	return (len(stack) == 0)
}