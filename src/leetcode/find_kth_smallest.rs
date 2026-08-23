use crate::leetcode::Solution;

use std::{cmp::max, collections::{HashMap, BTreeSet}};

/// leetcode 3116

struct PrimePartition {
    prime_cnt: HashMap<i32, u32>
} 

impl PrimePartition {
    /// 1 <= coin <= 25
    fn from(mut coin: i32) -> Self {
        const PRIMES: &[i32] = &[2, 3, 5, 7, 11, 13, 17, 19, 23, 29];
        let mut result = HashMap::<i32, u32>::new();
        for p in PRIMES {
            let mut times = 0;
            while coin % p == 0 {
                times += 1;
                coin /= p;
            }
            result.insert(*p, times);
            if coin == 1 {
                break;
            }
        }
        Self {
            prime_cnt: result
        }
    }

    fn lcm(&mut self, other: &PrimePartition) {
        for (p, other_cnt) in other.prime_cnt.iter() {
            if let Some(cnt) = self.prime_cnt.get_mut(p) {
                *cnt = max(*cnt, *other_cnt);
            } else {
                self.prime_cnt.insert(*p, *other_cnt);
            }
        }
    }

    fn value(&self) -> i64 {
        let mut result: i64 = 1;
        for (p, cnt) in self.prime_cnt.iter() {
            result *= (*p as i64).pow(*cnt as u32);
        }
        result
    }
}

impl Solution {
    fn max_coin(coin: &[i32]) -> i32 {
        *coin.iter().max().expect("iter.max should not be None")
    }
    /// iter_one returns
    /// Option<i32> the number found or none
    /// Vec<i32> next generation of modulo
    /// i32 the count of number in this generation
    fn iter_one(max_num: i32, m: Vec<i32>, coins: &[i32], k: i32) -> (Option<i32>, Vec<i32>, i32) {
        if k == 1 {
            return (Some(0), Vec::<i32>::new(), 1)
        }
        let mut nums = BTreeSet::<i32>::new();
        m.iter().enumerate()
        .for_each(|(i, mi)| {
                if *mi == 0 {
                    nums.insert(0);
                }
                for t in 1..=(mi+max_num-1)/coins[i] {
                    nums.insert(t * coins[i] - mi);
                }
        });
        let count = nums.len();
        if count < (k as usize) {
            return (
                None,
                m.iter().enumerate()
                .map(|(i, mi)| (max_num+mi)%coins[i])
                .collect(),
                count as i32
            )
        }
        let mut result = 0;
        let mut iter = nums.iter();
        for _ in 0..k as usize {
            result = *iter.next().expect("shouldn't be none");
        }
        (Some(result), Vec::<i32>::new(), count as i32)
    }

    fn num_in_lcm(coins: &[i32]) -> (i64, BTreeSet<i64>) {
        let pps = coins.iter()
                       .map(|x| PrimePartition::from(*x))
                       .collect::<Vec<PrimePartition>>();
        let mut lcm = PrimePartition::from(1);
        for pp in pps {
            lcm.lcm(&pp);
        }
        let result_lcm = lcm.value();
        let mut result_nums = BTreeSet::new();
        for c in coins {
            for t in (0..result_lcm).step_by(*c as usize) {
                result_nums.insert(t);
            }
        }
        (result_lcm, result_nums)
    }

    pub fn find_kth_smallest(coins: Vec<i32>, mut k: i32) -> i64 {
        if coins.contains(&1) {
            return k as i64;
        }
        #[cfg(feature = "slow_impl")]
        {
            let max_num = Solution::max_coin(&coins);
            let mut iter_num = 0;
            let mut m = Vec::<i32>::new();
            m.resize(coins.len(), 0);
            k = k + 1; // 0 also as valid coin
            while k > 0 {
                match Solution::iter_one(max_num, m, &coins, k) {
                    (Some(modulo), _, _) => return iter_num*(max_num as i64) + (modulo as i64),
                    (None, new_m, count) => {
                        k -= count;
                        m = new_m;
                    }
                }
                iter_num += 1;
            }
            0
        }
        let (lcm, nums) = Solution::num_in_lcm(&coins);
        let lcm_cnt = (k as i64)/(nums.len() as i64);
        let num_idx = (k as i64)%(nums.len() as i64);
        let nums_in_range: Vec<i64> = nums.iter().copied().collect();
        lcm_cnt*lcm + nums_in_range[num_idx as usize]
    }
}

#[cfg(test)]
mod test {
    use super::super::Solution;
    #[test]
    fn test_iter_one() {
        let (res, m, count) = Solution::iter_one(6, vec![0, 0], &[2, 3], 5);
        assert_eq!(res, None);
        assert_eq!(m, vec![0, 0]);
        assert_eq!(count, 4);
        let (res, _, _) = Solution::iter_one(6, vec![0, 0], &[2, 3], 4);
        assert_eq!(res, Some(4));
    }

    #[test]
    fn test_iter_one_for_5_2() {
        let coins = &[5, 2];
        let (res, m, count) = Solution::iter_one(5, vec![0, 0], coins, 8);
        assert_eq!(res, None);
        assert_eq!(m, vec![0, 1]);
        assert_eq!(count, 3);
        let (res, m, count) = Solution::iter_one(5, m, coins, 5);
        assert_eq!(res, None);
        assert_eq!(m, vec![0, 0]);
        assert_eq!(count, 3);
    }

    #[test]
    fn test_find_kth_smallest() {
        assert_eq!(Solution::find_kth_smallest(vec![3,6,9], 3), 9);
        assert_eq!(Solution::find_kth_smallest(vec![5,2], 7), 12);
    }

    #[test]
    fn test_387() {
        let coins = &[3, 8, 7];
        let m = vec![0,0,0];
        let (res, m, count) = Solution::iter_one(8, m, coins, 7);
        assert_eq!(res, None);
        assert_eq!(m, vec![2, 0, 1]);
        assert_eq!(count, 4);
        let (res, m, count) = Solution::iter_one(8, m, coins, 3);
        assert_eq!(res, Some(4));
    }

    #[test]
    fn test_10_6() {
        assert_eq!(Solution::find_kth_smallest(vec![10,6], 8), 36);
    }

    #[test]
    fn test_5() {
        assert_eq!(Solution::find_kth_smallest(vec![5], 3), 15);
    }

    #[test]
    fn test() {
        // assert_eq!(Solution::find_kth_smallest(vec![8,9,12,11,15,6,25,23], 1123986064), 1);
        assert_eq!(Solution::find_kth_smallest(vec![5,25,23,16,7,8,10,6,11,15], 946326769), 1);
    }

}
