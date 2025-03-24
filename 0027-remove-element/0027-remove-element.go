func removeElement(nums []int, val int) int {
    count := 0
    length := len(nums)
    rightIndex := length - 1;
    for i := 0; i < length; i++ {
        if nums[i] == val {
            for ; rightIndex > i && rightIndex > 0; {
                if nums[rightIndex] != val {
                    break;
                }
                rightIndex--;
            }
            nums[i] = nums[rightIndex];
            if rightIndex != 0 {
                rightIndex--;
            }
        } else {
            count++;
        }
    }
    return count
}



