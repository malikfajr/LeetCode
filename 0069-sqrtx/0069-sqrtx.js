/**
 * @param {number} x
 * @return {number}
 */
var mySqrt = function(x) {
    if (x == 0) return 0;
    if (x < 4) return 1;
    
    let i = 2
    while (i*i <= x) {
        i++
    }
    
    return i-1
};