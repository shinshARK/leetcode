class Solution {
public:
    bool isAnagram(string s, string t) {
        if(s.length() != t.length()) return false;
        int s_freq[26] = {0}, t_freq[26] = {0};

        for(const auto& character : s) {
            s_freq[character - 'a']++;
        }

        for(const auto& character : t) {
            t_freq[character - 'a']++;
        }

        for(int i = 0; i < 26; i++) {
            if(s_freq[i] != t_freq[i]) return false;
        }
        
        return true;

        // if inputs contain unicode, use unordered_map ?

        

    }
};