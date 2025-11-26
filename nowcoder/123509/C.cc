/*
matched brackets
reverse bracket
*/
#include <cassert>
#include <iostream>
#include <string>

using namespace std;

int minimum_times_to_match(string str, int n) {
    if (n % 2 != 0) {
        return -1;
    }
    int in_stack_left_bracket = 0; // ( in stack
    int len_limit = n/2;
    int change_times = 0;
    int left_count = 0;
    bool break_loop = false;
    for (auto iter = str.begin(); iter != str.end(); iter++) {
        switch (*iter) {
        case '(': {
            if (left_count == len_limit) { // ( has been used out
                // all ( afterwards should be changed to )
                for (; iter != str.end(); iter++) {
                    if (*iter == '(') {
                        *iter = ')';
                        change_times++;
                    }
                }
                break_loop = true;
            } else { // not used out, left here
                left_count++;
                in_stack_left_bracket++;
            }
            break;
        }
        case ')': {
            if (in_stack_left_bracket == 0) {
                *iter = '(';
                left_count++;
                in_stack_left_bracket++;
                change_times++;
            } else {
                in_stack_left_bracket--;
            }
            break;
        }
        }
        if (break_loop) {
            break;
        }
    }
    return change_times;
}

void test() {
    cout << "test\n";
    assert((minimum_times_to_match("()))", 4) == 1) && "1");
    assert((minimum_times_to_match("(", 1) == -1) && "2");
    assert((minimum_times_to_match("(((())", 6) == 1) && "3");
}

int main(void)
{
#ifdef TEST
    test();
    return 0;
#endif
    int n;
    string str;
    cin >> n >> str;
    cout << minimum_times_to_match(str, n) << endl;
    /*
    )(
    ()
))(())
(((())
   (

aaaa
aaba
aaab
    */
    return 0;
}
