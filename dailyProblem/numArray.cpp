#include<vector>
using namespace std;

class NumArray {
private:
    vector<int> array;
public:
    NumArray(vector<int>& nums) {
        array = nums;
    }
    
    int sumRange(int i, int j) {
        int sum = 0;
        auto begin = array.begin() + i;
        auto end = array.begin() + j +1;
        while (begin != end) {
            sum += *begin;
            begin++;
        }
        return sum;
    }
};