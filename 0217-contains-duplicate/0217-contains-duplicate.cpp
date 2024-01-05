class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {

        std::unordered_map<int, int> freq;

        for(const auto& num : nums) {
            freq[num]++;

            if(freq[num] > 1) return true;
        } 


        return false;
        
    }
};