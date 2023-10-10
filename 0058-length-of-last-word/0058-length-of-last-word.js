/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    var lastWord = s.trim().split(" ").pop()
    return lastWord.length
};