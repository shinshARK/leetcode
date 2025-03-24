func getConcatenation(nums []int) []int {
    newLen := len(nums) * 2;
    ans := make([]int, newLen);
    for i, j := 0, newLen / 2; i < len(nums); i, j = i+1, j+1 {
        ans[i] = nums[i]
        ans[j] = nums[i]
    }
    return ans

}