class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {

        std::unordered_map<int, int> freq;

        for(const auto& num : nums) {
            freq[num]++;
            // auto res = freq.emplace(num, 1);
            // if(!res.second) res.first->second++;
            if(freq[num] > 1) return true;
        } 

        // for(const auto& pair : freq) {
        //     if(pair.second > 1) return true;
        // }

        return false;
        
    }
};