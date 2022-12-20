/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function (digits) {
    var res = new Array();
    var n = digits.length;
    if (n === 0) return res;
    //如果长度有1，则遍历第一个
    var m = {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }
    var digit1 = m[digits[0]];
    for (var a = 0; digit1 && a < digit1.length; a++) {
        var letters1 = digit1[a];
        //如果长度有2，则遍历第二个
        var digit2 = m[digits[1]];
        for (var b = 0; digit2 && b < digit2.length; b++) {
            var letters2 = digit2[b];
            //如果长度有3，则遍历第三个
            var digit3 = m[digits[2]];
            for (var c = 0; digit3 && c < digit3.length; c++) {
                var letters3 = digit3[c];
                //如果长度有4，则遍历第四个
                var digit4 = m[digits[3]];
                for (var d = 0; digit4 && d < digit4.length; d++) {
                    var letters4 = digit4[d];
                    //因为长度最多有4个，我们就不往下枚举了
                    //此时就得到了一个字母组合
                    res.push(letters1 + letters2 + letters3 + letters4);
                }
                if (digit4) return res;
                // 如果长度只有3个
                res.push(letters1 + letters2 + letters3);
            }
            if (digit3) return res;
            // 如果长度只有2个
            res.push(letters1 + letters2);
        }
        if (digit2) return res;
        // 如果长度只有1个
        res.push(letters1);
    }
    return res;
};