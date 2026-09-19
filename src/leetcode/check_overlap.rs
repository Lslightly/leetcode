/// 1410
/// 给你一个以 (radius, xCenter, yCenter) 表示的圆和一个与坐标轴平行的矩形 (x1, y1, x2, y2) ，其中 (x1, y1) 是矩形左下角的坐标，而 (x2, y2) 是右上角的坐标。
/// 如果圆和矩形有重叠的部分，请你返回 true ，否则返回 false 。
/// 换句话说，请你检测是否 存在 点 (xi, yi) ，它既在圆上也在矩形上（两者都包括点落在边界上的情况）。
/// 示例 1 ：
/// 输入：radius = 1, xCenter = 0, yCenter = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
/// 输出：true
/// 解释：圆和矩形存在公共点 (1,0) 。
/// 示例 2 ：
/// 输入：radius = 1, xCenter = 1, yCenter = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
/// 输出：false
/// 示例 3 ：
/// 输入：radius = 1, xCenter = 0, yCenter = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
/// 输出：true
/// 
/// 
/// 
/// 
use super::Solution;

use std::cmp::{max, min};

#[derive(Clone, Copy)]
struct Coordinate {
    x: i32,
    y: i32,
}

#[derive(Clone, Copy)]
struct Circle {
    center: Coordinate,
    radius: i32,
}

impl Circle {
    fn cover(&self, coord: Coordinate) -> bool {
        let center = self.center;
        (center.x - coord.x).pow(2) + (center.y - coord.y).pow(2) <= self.radius.pow(2)
    }
}

impl Solution {
    pub fn check_overlap(radius: i32, x_center: i32, y_center: i32, x1: i32, y1: i32, x2: i32, y2: i32) -> bool {
        let circle = Circle {
            center: Coordinate { x: x_center, y: y_center },
            radius,
        };
        let left_bound = max(x_center - radius, x1);
        let right_bound = min(x_center + radius, x2);
        let up_bound = min(y_center + radius, y2);
        let down_bound = max(y_center - radius, y1);
        for x in left_bound..=right_bound {
            for y in down_bound..=up_bound {
                if circle.cover(Coordinate { x, y }) {
                    return true;
                }
            }
        }
        false
    }
}

#[cfg(test)]
mod test {
    use super::Solution;
    #[test]
    fn test1() {
        assert!(Solution::check_overlap(1, 0, 0, 1, -1, 3, 1));
        assert!(!Solution::check_overlap(1, 1, 1, 1, -3, 2, -1));
        assert!(Solution::check_overlap(1, 0, 0, -1, 0, 0, 1));
    }
}