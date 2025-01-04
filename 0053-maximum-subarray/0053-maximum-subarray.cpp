class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int best = INT_MIN, sum = 0;
        for(auto x : nums) {
            sum = max(x, sum+x);
            best = max(best, sum);
        }
        return best;
        
    }
};