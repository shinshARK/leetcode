class Solution {
public:
    bool isValidSudoku(vector<vector<char>>& board) {
        std::unordered_set<char> row, col[9], box[3][3];

        int i = 0, j = 0;
        
        for(int i = 0; i < 9; i++) {
            for(int j = 0; j < 9; j++) {
                if(board[i][j] == '.') continue;

                auto it_row = row.find(board[i][j]);
                auto it_box = box[i/3][j/3].find(board[i][j]);
                auto it_col = col[j].find(board[i][j]);

                if(
                    it_row != row.end() || 
                    it_box != box[i/3][j/3].end() || 
                    it_col != col[j].end()) 
                    return false;

                row.insert(board[i][j]);
                box[i/3][j/3].insert(board[i][j]);
                col[j].insert(board[i][j]);
                
            }
            
            row.clear(); 
        }

        return true;
        
    }
};