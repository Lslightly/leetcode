/*
 * @lc app=leetcode.cn id=215 lang=rust
 *
 * [215] 数组中的第K个最大元素
 */
use super::Solution;

// @lc code=start
use rand::Rng;
use std::{cmp::Ordering, collections::HashMap};

impl Solution {
    fn partition(nums: &mut [i32], begin: usize, end: usize, k: usize) -> usize {
        let mut rng = rand::thread_rng();
        let pivot_index = rng.gen_range(begin..=end);
        nums.swap(pivot_index, end);
        let pivot = nums[end];
        let mut to_fill_index = begin;
        let mut greater_num_cnt_map: HashMap<i32, usize> = HashMap::new();
        let mut fewer_num_cnt_map: HashMap<i32, usize> = HashMap::new();
        for j in begin..end {
            if nums[j] >= pivot {
                if let Some(vref) = greater_num_cnt_map.get_mut(&nums[j]) {
                    *vref += 1;
                } else {
                    greater_num_cnt_map.insert(nums[j], 1);
                }
                nums.swap(to_fill_index, j);
                to_fill_index += 1;
            } else {
                if let Some(vref) = fewer_num_cnt_map.get_mut(&nums[j]) {
                    *vref += 1;
                } else {
                    fewer_num_cnt_map.insert(nums[j], 1);
                }
            }
        }
        nums.swap(to_fill_index, end);
        if greater_num_cnt_map.len() == 0 {
            return to_fill_index;
        }
        if greater_num_cnt_map.len() == 1 && to_fill_index > k - 1 {
            return k - 1;
        }
        if fewer_num_cnt_map.len() <= 1 && to_fill_index < k - 1 {
            return k - 1;
        }
        to_fill_index
    }
    pub fn find_kth_largest_for_slice(nums: &mut [i32], k: usize, begin: usize, end: usize) -> i32 {
        let partition_index = Solution::partition(nums, begin, end, k);
        match partition_index.cmp(&(k - 1)) {
            Ordering::Equal => nums[partition_index],
            Ordering::Less => {
                Solution::find_kth_largest_for_slice(nums, k, partition_index + 1, end)
            }
            Ordering::Greater => {
                Solution::find_kth_largest_for_slice(nums, k, begin, partition_index - 1)
            }
        }
    }
    pub fn find_kth_largest(mut nums: Vec<i32>, k: i32) -> i32 {
        let end = nums.len() - 1;
        Solution::find_kth_largest_for_slice(&mut nums, k as usize, 0, end)
    }
}
// @lc code=end

#[cfg(test)]
mod test {
    use super::super::*;
    #[test]
    fn test_kth_large() {
        let nums = vec![3, 2, 1, 5, 6, 4];
        assert_eq!(Solution::find_kth_largest(nums, 2), 5)
    }

    #[test]
    fn test_double() {
        let nums = vec![-1, -1];
        assert_eq!(Solution::find_kth_largest(nums, 2), -1);
    }
}
