#include <algorithm>
#include <cassert>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

enum Direction {
    Left,
    Up,
    Down,
    Right,
    DirectNum,
};

/*
在一个无限的二维平面上，有一个机器人初始位于原点 
(0,0)。你收到了一串由字符 𝐿 / 𝑅 / 𝑈 / 𝐷 组成的指令，分别表示向左/右/上/下移动一格。现在你可以将这串指令的顺序任意重排，然后让机器人按新的顺序执行全部指令。请问，最多能让机器人访问到多少个不同的位置（包含起点）

LRUD
*/

typedef struct OpInfo {
    Direction op;
    int count;
}OpInfo;

bool conflict(Direction a, Direction b) {
    return a + b == DirectNum-1;
}

int nodes_of_path(vector<int> p) {
    if (p[1] == 0) {
        return 1+p[0];
    }
    if (p[2] == 0) {
        return 1+p[0]+p[1];
    }
    if (p[3] == 0) {
        return 1+p[0]+p[1]+p[2];
    }
    if (p[3] == p[0]) {
        return p[0]+p[1]+p[2]+p[3];
    }
    return 1+p[0]+p[1]+p[2]+p[3];
}

int possible_pos_count(string str) {
    vector<OpInfo> opinfos(DirectNum);
    for (auto op = 0; op < DirectNum; op++) {
        opinfos[op] = OpInfo{
            .op = Direction(op),
            .count = 0,
        };
    }
    for (auto ch: str) {
        switch (ch) {
        case 'L':
            opinfos[Left].count++;
            break;
        case 'R':
            opinfos[Right].count++;
            break;
        case 'U':
            opinfos[Up].count++;
            break;
        case 'D':
            opinfos[Down].count++;
            break;
        }
    }
    sort(opinfos.begin(), opinfos.end(), [](const OpInfo&a, const OpInfo& b) -> bool {
        if (a.count > b.count) {
            return true;
        }
        return false;
    });
    if (conflict(opinfos[0].op, opinfos[1].op)) {
        // 0 2 1 3
        return nodes_of_path({opinfos[0].count, opinfos[2].count, opinfos[1].count, opinfos[3].count});
    } else {
        // 0 1 ? ?
        if (conflict(opinfos[1].op, opinfos[2].op)) {
            // 0 1 3 2
            return nodes_of_path({opinfos[0].count, opinfos[1].count, opinfos[3].count, opinfos[2].count});
        } else {
            // 0 1 2 3
            return nodes_of_path({opinfos[0].count, opinfos[1].count, opinfos[2].count, opinfos[3].count});
        }
    }
}

void group_handle() {
    int n;
    string str;
    cin >> n >> str;
    cout << possible_pos_count(str) << endl;
}

void test() {
    assert(possible_pos_count("LR") == 2);
    assert(possible_pos_count("UUUD") == 4);
    assert(possible_pos_count("URDL") == 4);
    assert(possible_pos_count("UUDDRRLL") == 8);
    assert(possible_pos_count("UUUUURRRRDDDDDLLL") == 18);
}

int main(void)
{
#ifdef TEST
    test();
    return 0;
#endif
    int T;
    cin >> T;
    for (int grp = 0; grp < T; grp++) {
        group_handle();
    }
    return 0;
}
