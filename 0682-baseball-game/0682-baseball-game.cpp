template <typename T>
void printStack(std::stack<T> s) {
  std::cout << "Stack (top to bottom):" << std::endl;
  if (s.empty()) {
    std::cout << "(empty)" << std::endl;
    return;
  }
  while (!s.empty()) {
    std::cout << s.top() << " ";
    s.pop();
  }
  std::cout << "\n---" << std::endl;
}

class Solution {
public:
    int calPoints(vector<string>& operations) {
        std::stack<int> records;
        int newScores[2];
        for(int i = 0; i < operations.size(); i++) {
            // cout << "--------------------------\n";
            // printStack(records);
            // cout << "current operations: " << operations[i] << endl;
            // cout << operations[i][0] << endl;
            // cout << isdigit(operations[i][0]) << endl;
            // cout << stoi(operations[i]) << endl;
            if(isdigit(operations[i][0]) || operations[i][0] == '-') {
                // cout << "is a digit, adding..." << endl;
                records.push(stoi(operations[i]));
            } else if(operations[i] == "+") {
                // cout << "adding..." << endl;
                
                newScores[0] = records.top();
                records.pop();
                newScores[1] = records.top();
                records.pop();
                
                records.push(newScores[1]);
                records.push(newScores[0]);

                records.push(newScores[0] + newScores[1]);
            } else if(operations[i] == "D") {
                // cout << "doubling..." << endl;
                newScores[0] = records.top() * 2;
                cout << newScores[0] << endl;
                records.push(newScores[0]);
            } else if(operations[i] == "C") {
                // cout << "clearing..." << endl;

                records.pop();
            }
            // cout << "after operations: ";
            // printStack(records);
            // cout << "--------------------------\n";
        }
        int total = 0;
        while(!records.empty()) {
            // cout << curr << " ";
            total += records.top();
            records.pop();
        }
        return total;
    }
};