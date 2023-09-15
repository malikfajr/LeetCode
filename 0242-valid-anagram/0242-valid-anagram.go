func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    
    var mapS = make(map[byte]int)
    var mapT = make(map[byte]int)
    
    for i := 0; i < len(s); i++ {
        mapS[s[i]]++
        mapT[t[i]]++
    }
    
    return reflect.DeepEqual(mapS, mapT) 
}