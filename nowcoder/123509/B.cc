#include <iostream>

using namespace std;

int main() {
    int n;
    cin >> n;
    cout << 1;
    int limit = (n % 2 == 0)?(n-1):n;
    for (int i = 2; i <= limit; i++) {
        if (i % 2 == 0) {
            cout << ' ' << i+1;
        } else {
            cout << ' ' << i-1;
        }
    }
    if (n % 2 == 0) {
        cout << ' ' << n;
    }
    return 0;
}