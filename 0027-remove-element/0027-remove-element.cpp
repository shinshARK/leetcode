

class Solution {
public:
    int removeElement(vector<int>& nums, int val) {
        int count = 0;
        int length = nums.size();
        int leftIndex = length - 1;
        for(int i = 0; i < length; i++) {
            if(nums[i] == val) {
                while(leftIndex > i && leftIndex > 0) {
                    if(nums[leftIndex] != val) {
                        break;
                    } 
                    leftIndex--;
                }
                nums[i] = nums[leftIndex];
                if(leftIndex != 0) {
                    leftIndex--;

                }
            } else {
                count++;
            }
        }

        return count;
    }
};