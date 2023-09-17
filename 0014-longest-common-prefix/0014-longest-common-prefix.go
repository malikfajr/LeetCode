func longestCommonPrefix(strs []string) string {
    sort.Strings(strs)
    first := strs[0]
    last := strs[len(strs) - 1]

    var ans bytes.Buffer
    for i := 0; i < len(first); i++ {
        if first[i] != last[i] {
            break;
        }
        ans.WriteByte(first[i])
    }
    
    
    return ans.String()
}