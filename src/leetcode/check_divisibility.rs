use crate::leetcode::Solution;

impl Solution {
    fn digits_of_n(mut n: i32) -> Vec<i32> {
        let mut result = vec![];
        while n != 0 {
            result.push(n % 10);
            n /= 10;
        }
        result
    }
    pub fn check_divisibility(n: i32) -> bool {
        let digits = Solution::digits_of_n(n);
        n % (digits.iter().fold(0, |acc, x| acc + x)
            + digits.iter().fold(1, |acc, x| acc * x))
            == 0
    }
}

#[cfg(test)]
mod test {
    use crate::leetcode::{Solution};

    #[test]
    fn test1() {
        assert!(Solution::check_divisibility(99));
        assert!(!Solution::check_divisibility(23));
    }
}