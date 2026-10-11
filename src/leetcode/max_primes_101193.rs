use crate::leetcode::Solution;

impl Solution {
    fn is_prime(n: i32) -> bool {
        for i in 2..=n/2 {
            if i * i > n {
                return true;
            }
            if n % i == 0 {
                return false;
            }
        }
        true
    }

    pub fn max_primes(n: i32, s: i32) -> Vec<i32> {
        if s == 1 {
            return vec![];
        }
        let mut sum = 0;
        let mut result = vec![];
        for i in 2..=n {
            if Solution::is_prime(i) {
                sum += i;
                if sum > s {
                    return result;
                }
                result.push(i);
            }
        }
        result
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test1() {
        assert_eq!(Solution::max_primes(7, 20), vec![2,3,5,7])
    }

    #[test]
    fn test2() {
        assert_eq!(Solution::max_primes(15, 15), vec![2,3,5])
    }

    #[test]
    fn test3() {
        assert_eq!(Solution::max_primes(5, 1), vec![])
    }
}