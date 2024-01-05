class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        int size = nums.size();
        
        
        int arrLeft[size], arrRight[size];
        arrLeft[0] = 1, arrRight[size - 1] = 1;

        vector<int> ans;
        
        for(int i = 1, j = size - 2; i < size; i++, j--) {
            arrLeft[i] = arrLeft[i - 1] * nums[i - 1];
            arrRight[j] = arrRight[j + 1] * nums[j + 1];

        }
        
        
        for(int i = 0; i < size; i++) ans.push_back(arrLeft[i] * arrRight[i]);

        return ans;
        
    }
};