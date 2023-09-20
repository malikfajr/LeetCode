import (
    "strings"
    "regexp"
)

func isPalindrome(s string) bool {
    nonAlphabetRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
    
    wordLower := strings.ToLower(s)
    word := nonAlphabetRegex.ReplaceAllString(wordLower, "")
    
    var reversedWord  bytes.Buffer
    for i := len(word) - 1; i >= 0; i-- {
        reversedWord.WriteByte(word[i])
    }

    if word != reversedWord.String() {
        return false
    }
    return true

}