#include<iostream>
#include<vector>
using namespace std;

bool isMonotonic(vector<int>& A) {
    if (A.size() <= 2) {
        return true;
    }
    int prev = 0;
    for (int i = 1; i < A.size(); i++) {
        int cur = A[i] - A[i-1];
        if (cur * prev < 0) {
            return false;
        }
        if (cur != 0) {
            prev = cur;
        }
    }
    
    return true;
}
int main() {
    vector<int> A = {11,11,9,4,3,3,3,1,-1,-1,3,3,3,5,5,5};
    auto resu = isMonotonic(A);
    cout << resu << endl;

}