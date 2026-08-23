use super::Solution;

use std::{cmp::max, collections::{BTreeMap, BTreeSet, HashSet}};

impl Solution {
    fn k1(nums: Vec<i32>) -> i32 {
        let mut num_cnt_map = BTreeMap::<i32, i32>::new();
        for val in nums {
            if let Some(cnt) = num_cnt_map.get_mut(&val) {
                *cnt += 1;
            } else {
                num_cnt_map.insert(val, 1);
            }
        }
        for (num, cnt) in num_cnt_map.iter().rev() {
            if *cnt == 1 {
                return *num
            }
        }
        -1
    }
    fn klen(nums: Vec<i32>) -> i32 {
        let mut num_set = BTreeSet::<i32>::new();
        for val in nums {
            num_set.insert(val);
        }
        *num_set.iter().rev().next().expect("should not be empty")
    }
    pub fn largest_integer(nums: Vec<i32>, k: i32) -> i32 {
        if k == 1 {
            return Solution::k1(nums);
        }
        if (k as usize) == nums.len() {
            return Solution::klen(nums);
        }
        if nums[0] == nums[nums.len()-1] {
            return -1;
        }
        let mut num_set = HashSet::<i32>::new();
        for val in &nums[1..nums.len()-1] {
            num_set.insert(*val);
        }
        let zero_elem = nums[0];
        let last_elem = nums[nums.len()-1];
        let zero_in = num_set.contains(&zero_elem);
        let last_in = num_set.contains(&last_elem);
        if zero_in && last_in {
            return -1;
        } else if zero_in {
            return last_elem;
        } else if last_in {
            return zero_elem;
        } else {
            if zero_elem == last_elem {
                return -1;
            }
            return max(zero_elem, last_elem);
        }
    }
}

#[cfg(test)]
mod test {
    use super::super::*;
    #[test]
    fn test_largest_integer() {
        let nums = vec![3, 9, 2, 1, 7];
        assert_eq!(Solution::largest_integer(nums, 3), 7);
        let nums = vec![3, 9, 7, 2, 1, 7];
        assert_eq!(Solution::largest_integer(nums, 4), 3);
        let nums = vec![0, 0];
        assert_eq!(Solution::largest_integer(nums, 1), -1);
        let nums = vec![0, 0];
        assert_eq!(Solution::largest_integer(nums, 2), 0);
    }
}
