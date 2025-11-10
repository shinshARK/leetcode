func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    readIndex, writeIndex := 1, 0
    for ; readIndex < len(nums); {
        if nums[writeIndex] == nums[readIndex] {
            readIndex++
        } else {
            writeIndex++
            nums[writeIndex] = nums[readIndex]
            readIndex++
        }
    }
    return writeIndex + 1
    
}