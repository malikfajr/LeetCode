func isPalindrome(x int) bool {
    if x < 0 || x % 10 == 0 && x != 0 {
        return false
    }

    var reverse int = 0
    
    for x > reverse {
        sisa := x % 10
        reverse = reverse * 10 + sisa
        x = x / 10
    }
    
    return x == reverse || x == reverse/10
}