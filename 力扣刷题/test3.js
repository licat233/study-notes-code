/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
    var sl = s.length
    if (sl === 0) {
        return 0
    }
    var max = 0;
    var dl = 0;
    var z = 0;
    var dm = new Map();
    var sv = "";
    for (var k = 0; k < sl; k++) {
        sv = s[k]
        if (dm.has(sv)) {
            z = k - dm.get(sv);
            if (z <= dl) {
                dl = z;
            }else{
                dl++;
            }
        } else {
            dl++;
        }
        dm.set(sv, k);
        if (dl > max) {
            max = dl;
            continue;
        }
        if (sl - k - 1 + dl <= max) {
            return max;
        }
    }
    return max;
}

var data = "tmmzuxt";
var res = 0;
res = lengthOfLongestSubstring(data);
console.log(res); // 5
// for (var i = 0; i < 100000; i++) {
    // res = lengthOfLongestSubstring(data);
// }
// console.log(res); // 6