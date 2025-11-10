func shuffle(nums []int, n int) []int {
    // using a new array
    shuffled := make([]int, n * 2)
    fi, se := 0, n 
    for i := 0; i < (n * 2) - 1; i+=2 {
        shuffled[i] = nums[fi]
        fi++
        shuffled[i+1] = nums[se]
        se++      
    }

    return shuffled
    // inplace
}