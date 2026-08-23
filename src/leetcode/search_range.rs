/*
 * @lc app=leetcode.cn id=34 lang=rust
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */
use crate::leetcode::Solution;

// @lc code=start
use std::{cmp::Ordering, ops::Sub};
impl Solution {
    /**
     * find one target in nums, return left, right bound. return None if not found
     */
    fn find_one_target(nums: &Vec<i32>, target: i32) -> Option<(usize, usize)> {
        if nums.len() == 0 {
            return None;
        }
        let mut left = 0;
        let mut right = nums.len() - 1;
        while left <= right {
            let mid = (left + right) / 2;
            match nums[mid].cmp(&target) {
                Ordering::Equal => return Some((left, right)),
                Ordering::Less => match mid.checked_add(1) {
                    Some(res) => left = res,
                    None => return None,
                },
                Ordering::Greater => match mid.checked_sub(1) {
                    Some(res) => right = res,
                    None => return None,
                },
            }
        }
        None
    }

    fn find_begin_target(nums: &Vec<i32>, mut left: usize, mut right: usize, target: i32) -> usize {
        while left <= right {
            let mid = (left + right) / 2;
            match nums[mid].cmp(&target) {
                Ordering::Equal => match mid.checked_sub(1) {
                    Some(res) => right = res,
                    None => return left,
                },
                Ordering::Less => left = mid + 1,
                _ => panic!(
                    "impossible for find_begin_target left: {}, right: {}, nums[mid: {}]={}",
                    left, right, mid, nums[mid]
                ),
            }
        }
        left
    }

    fn find_end_target(nums: &Vec<i32>, mut left: usize, mut right: usize, target: i32) -> usize {
        while left <= right {
            let mid = (left + right) / 2;
            match nums[mid].cmp(&target) {
                Ordering::Equal => match mid.checked_add(1) {
                    Some(res) => left = res,
                    None => return right,
                },
                Ordering::Greater => right = mid - 1,
                _ => panic!(
                    "impossible for find_begin_target left: {}, right: {}, nums[mid: {}]={}",
                    left, right, mid, nums[mid]
                ),
            }
        }
        right
    }

    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let Some((left, right)) = Solution::find_one_target(&nums, target) else {
            return vec![-1, -1];
        };
        vec![
            Solution::find_begin_target(&nums, left, right, target) as i32,
            Solution::find_end_target(&nums, left, right, target) as i32,
        ]
    }
}
// @lc code=end

#[cfg(test)]
mod test {
    use super::super::*;
    #[test]
    fn test1() {
        let nums = parse_int_array("[5,7,7,8,8,10]");
        assert_eq!(Solution::search_range(nums, 8), vec![3, 4],);
        let nums = parse_int_array("[5,7,7,8,8,10]");
        assert_eq!(Solution::search_range(nums, 6), vec![-1, -1],);
    }
}
