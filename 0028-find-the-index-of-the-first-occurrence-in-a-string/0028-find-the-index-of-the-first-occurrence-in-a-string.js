/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function(haystack, needle) {
    const s = haystack.toLowerCase()
    const key = needle.toLowerCase()
    const isStringIncluded = s.includes(key)
    
    if (! isStringIncluded) {
        return -1;
    }
    
    return s.split(key)[0].length
    
};

