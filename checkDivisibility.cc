#include <vector>
using namespace std;
class Solution {
public:
    vector<int> digitsOf(int n) {
        vector<int> results{};
        while (n != 0) {
            results.push_back(n % 10);
            n /= 10;
        }
        return results;
    }
    bool checkDivisibility(int n) {
        auto digits = digitsOf(n);
        int sum = 0, mul = 1;
        for (auto digit: digits) {
            sum += digit;
            mul *= digit;
        }
        return n % (sum + mul) == 0;
    }
};