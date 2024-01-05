class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int x = 1, size = nums.size();
        vector<int> ans(size);

        for(int i = 0; i < size; i++) {
            ans[i] = x;
            x *= nums[i];
        }

        x = 1;
        for(int i = size - 1; i >= 0; i--) {
            ans[i] *= x;
            x *= nums[i];
        }


        return ans;
        
    }
};