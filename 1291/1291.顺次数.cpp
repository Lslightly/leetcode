#define BACKWARD_HAS_DW 1
#include "../lib/backward.hpp"
#define BACKWARD 1
/*
 * @lc app=leetcode.cn id=1291 lang=cpp
 *
 * [1291] 顺次数
 */

// @lc code=start
#include <algorithm>
#include <cassert>
#include <cstdio>
#include <iostream>
#include <vector>

using namespace std;
class Solution {
public:
    int getDigitNum(int input) {
        int num = 0;
        auto tmp = input;
        while (tmp != 0) {
            num++;
            tmp /= 10;
        }
        return num;
    }

    int getHighestDigit(int input) {
        int last = 0;
        auto tmp = input;
        while (tmp != 0) {
            last = tmp;
            tmp /= 10;
        }
        return last;
    }

    int genInOrderNum(int digitNum, int first) {
        if (digitNum + first > 10) {
#ifdef BACKWARD
            using namespace backward;
            StackTrace st;
            st.load_here();
            Printer p;
            p.object = true;
            p.color_mode = ColorMode::always;
            p.address = true;
            p.print(st, stderr);
#else
            printf("digitNum %d + first %d > 10\n", digitNum, first);
#endif
        }
        int result = 0;
        for (int i = first; i < first+digitNum; i++) {
            result = (result*10)+i;
        }
        return result;
    }

    // gen 11..1
    int genInOrderIncStep(int digitNum) {
        int result = 0;
        for (int i = 0; i < digitNum; i++) {
            result = (result*10)+1;
        }
        return result;
    }

    vector<int> sequentialDigits(int low, int high) {
        auto lowDigitNum = getDigitNum(low), highDigitNum = getDigitNum(high);
        auto lowHighest = getHighestDigit(low), highHighest = getHighestDigit(high);
        vector<int> result = {};
        if (lowDigitNum == highDigitNum) {
            // same digit nums
            if (lowHighest+lowDigitNum > 10) { // low 90
                return result; // empty
            }
            auto tmp = genInOrderNum(lowDigitNum, lowHighest);
            auto step = genInOrderIncStep(lowDigitNum);
            for (int first = lowHighest; first <= highHighest; first++, tmp += step) {
                if (low <= tmp && tmp <= high) {
                    result.push_back(tmp);
                } else {
                    break;
                }
            }
            return result;
        }

        // different digit nums
        int digitNum = lowDigitNum, tmp = 0, step = 0;
        if (lowHighest+digitNum <= 10) {
            /// low digit num operation
            tmp = genInOrderNum(digitNum, lowHighest);
            step = genInOrderIncStep(digitNum);
            if (low <= tmp) {
                result.push_back(tmp); // push num closest to low
            }
            tmp += step;
            // push left num
            for (int first = lowHighest+1; first <= 10-digitNum; first++, tmp += step) {
                result.push_back(tmp);
            }
        }
        /// middle digit num operation
        digitNum++;
        for (; digitNum < highDigitNum; digitNum++) {
            tmp = genInOrderNum(digitNum, 1);
            step = genInOrderIncStep(digitNum);
            for (int first = 1; first <= 10-digitNum; first++, tmp += step) {
                result.push_back(tmp);
            }
        }

        /// high digit num operation
        if (digitNum + 1 <= 10) {
            tmp = genInOrderNum(digitNum, 1);
            step = genInOrderIncStep(digitNum);
            for (int first = 1; first <= min(10-digitNum, highHighest); first++, tmp += step) {
                if (tmp <= high) {
                    result.push_back(tmp);
                }
            }
        }
        return std::move(result);
    }
};
// @lc code=end

int main() {
    Solution sol;
    auto res = sol.sequentialDigits(10, 1000000000);
    vector<int> exp{1234,2345,3456,4567,5678,6789,12345};
    assert(res.size() == exp.size());
    for (auto i = 0; i < res.size(); i++) {
        assert(res[i] == exp[i]);
    }
    return 0;
}
