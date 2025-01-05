/**
 * @param {string} s
 * @param {number[][]} shifts
 * @return {string}
 */
var shiftingLetters = function(s, shifts) {

    // console.log(s)

    let prefixes = new Array(s.length + 1).fill(0)
    for(let i = 0; i < shifts.length; i++) {
        if(shifts[i][2]) {
            prefixes[shifts[i][0]]--
            prefixes[shifts[i][1]+1]++
        } else {
            prefixes[shifts[i][0]]++
            prefixes[shifts[i][1]+1]--
        }
    }
    let diff = 0;
    let newstring = [];

    // console.log(prefixes)

    for(let i = prefixes.length - 1; i > 0; i--) {
        
        diff += prefixes[i]
        // console.log(`shifting ${s[i-1]} by ${diff}`)
        newstring[i-1] = wrapChar(s[i-1], diff)
    }

    // console.log(newstring.join(''))
    return newstring.join('')
    
};


function wrapChar(char, offset) {
    return String.fromCharCode(((char.charCodeAt(0) - 97 + offset) % 26 + 26) % 26 + 97);
}
