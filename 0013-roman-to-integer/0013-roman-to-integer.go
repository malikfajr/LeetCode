func romanToInt(s string) int {
    var keys = map[byte]uint16{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    var ans uint16 = 0
    
    for i := 0; i < len(s) - 1; i++ {
        if keys[s[i]] < keys[s[i + 1]] {
            ans = ans - keys[s[i]]
        } else {
            ans = ans + keys[s[i]]
        }
    }
    
    ans = ans + keys[s[len(s) - 1]]
    return int(ans)
}