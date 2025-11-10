func removeElement(nums []int, val int) int {
    write := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != val{
            nums[write] = nums[i]
            write++
        }
    }
    return write
}