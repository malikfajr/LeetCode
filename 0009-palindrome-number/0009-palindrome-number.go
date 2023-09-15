func isPalindrome(x int) bool {
    var reverse int = 0
    var temp int = x
    
    for temp > 0 {
        sisa := temp % 10
        reverse = reverse * 10 + sisa
        temp = temp / 10
    }
    
    return x == reverse
}