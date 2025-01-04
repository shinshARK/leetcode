/**
 * @param {string} s
 * @return {number}
 */
 var countPalindromicSubsequence = function(s) {

    const indices = new Map();
    
    for (let i = 0; i < 26; i++) {
        const key = String.fromCharCode(97 + i); // 'a' to 'z'
        indices.set(key, []); 
    }

    
    for(let i = 0; i < s.length; i++) {
        indices.get(s[i]).push(i)
    }

    let count = 0

    for(let i = 0; i < 26; i++) {
        let pivotIndices = indices.get(String.fromCharCode(i+97))

        if(pivotIndices.length < 2) continue
        
        for(let j = 0; j < 26; j++) {
            let arr = indices.get(String.fromCharCode(j+97))

            if(arr.length === 0) continue

            if( binarySearch(pivotIndices[0], pivotIndices[pivotIndices.length - 1], indices.get(String.fromCharCode(j+97) ) ) )  count++
        
        }
    }

    return count
    
    
};

function binarySearch(min, max, arr) {

    let left = 0
    let right = arr.length - 1
    while(left <= right) {
        const mid = Math.floor((left + right)/2)
        if(arr[mid] > min && arr[mid] < max) {
            return true
        } else if(arr[mid] <= min) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}