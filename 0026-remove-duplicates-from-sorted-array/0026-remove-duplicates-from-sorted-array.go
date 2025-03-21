func removeDuplicates(nums []int) int {

    reference := nums[0]
    unique := 1
    for i := 1; i < len(nums); {
        if nums[i] == reference {
            i++
        } else {
            nums[unique] = nums[i]
            reference = nums[i]
            unique++
        }
    }
    return unique
}