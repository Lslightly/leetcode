#include <vector>
using namespace std;
/*
 * @lc app=leetcode.cn id=74 lang=cpp
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
class Solution {
    enum CmpResult {
        Less,
        InRange,
        Greater,
        CmpNum,
    };
    CmpResult cmp(vector<vector<int>>& matrix, int rowNum, int target) {
        if (target >= matrix[rowNum].front() && target <= matrix[rowNum].back()) {
            return InRange;
        }
        if (target < matrix[rowNum].front()) {
            return Less;
        }
        if (target > matrix[rowNum].back()) {
            return Greater;
        }
        return CmpNum;
    }
    int searchRow(vector<vector<int>>& matrix, int target) {
        int low = 0, high = matrix.size()-1;
        while (low <= high) {
            auto mid = (low+high)/2;
            switch (cmp(matrix, mid, target)) {
            case Less:
                high = mid-1;
                break;
            case InRange:
                return mid;
            case Greater:
                low = mid+1;
                break;
            default:
                throw "unreachable";
            }
        }
        return -1;
    }
    bool searchElem(vector<int>& row, int target) {
        int low = 0, high = row.size()-1;
        while (low <= high) {
            auto mid = (low+high)/2;
            auto midElem = row[mid];
            if (midElem < target) {
                low = mid+1;
            } else if (midElem > target) {
                high = mid-1;
            } else {
                return true;
            }
        }
        return false;
    }
public:
    bool searchMatrix(vector<vector<int>>& matrix, int target) {
        int low = 0, high = matrix.size()-1;       
        int tgtRow = searchRow(matrix, target);
        if (tgtRow == -1) {
            return false;
        }
        return searchElem(matrix[tgtRow], target);
    }
};
// @lc code=end

