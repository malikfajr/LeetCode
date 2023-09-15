func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    var mapS = make(map[byte]uint16)
    var mapT = make(map[byte]uint16)
    
    for i := 0; i < len(s); i++ {
        mapS[s[i]]++
        mapT[t[i]]++
    }
    
    for char, count := range mapS {
        if value, valid := mapT[char]; !valid || count != value {
            return false
        }
    }
    
    return true
}