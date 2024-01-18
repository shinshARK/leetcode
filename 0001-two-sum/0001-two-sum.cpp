class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {

        vector<pair<int, int>> sorted;
        for(int i = 0; i < nums.size(); i++) {
            sorted.push_back({nums[i], i});
        }
        
        sort(sorted.begin(), sorted.end());

        vector<int> index(2, 0);
        int current;

        for(int i = 0, j = sorted.size() - 1; i != j; ) {

            current = sorted[i].first + sorted[j].first;
            if(current == target) {
                index[0] = sorted[i].second, index[1] = sorted[j].second;
                return index;
            } 
            else if (current < target) i++;
            else j--;
        } 
        
        
        
        
        return index;

        
    }
};