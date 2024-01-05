class Solution {
public:

    static int sortByValue(std::pair<int, int> a, std::pair<int, int> b) {
        return a.second > b.second;
    }

    vector<int> topKFrequent(vector<int>& nums, int k) {

        // vector<pair<int, int>> freq;
        std::unordered_map<int, int> freq;
        vector<int> ans;

        for(const auto& num : nums) {
            freq[num]++;
        }
        
        vector<pair<int, int>> freq_vector(freq.begin(), freq.end());

        sort(freq_vector.begin(), freq_vector.end(), sortByValue);

        for(int i = 0; i < k; i++) ans.push_back(freq_vector[i].first);

        return ans;

        
    }
};