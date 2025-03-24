func shuffle(nums []int, n int) []int {
    numsCopy := make([]int, n * 2)
    _ = copy(numsCopy, nums)
    xIndex := 0;
    yIndex := n;
    for i := 0; i < n * 2 - 1; i++ {
        if i % 2 == 0 {
            nums[i] = numsCopy[xIndex];
            xIndex++;
        } else {
            nums[i] = numsCopy[yIndex];
            yIndex++;
        }

    }
    return nums
}