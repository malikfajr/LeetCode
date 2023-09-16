func romanToInt(s string) int {
    var keys = map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    var ans = 0
    var prev = 0;
    
    for i := len(s) - 1; i >= 0; i-- {
        num := keys[s[i]]    
        if num < prev {
            ans = ans - num
        } else {
            ans = ans + num
        }
        prev = num
    }
    
    return (ans)
}