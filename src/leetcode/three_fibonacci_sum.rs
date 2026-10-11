use super::Solution;

impl Solution {
    pub fn three_fibonacci_sum(n: i32) -> bool {
        if n % 2 != 0 {
            return false;
        }
        let to_check = n/2;
        let mut m = 0;
        let mut n = 1;
        while m + n < to_check {
            (m, n) = (n, m+n)
        }
        m + n == to_check
    }
}