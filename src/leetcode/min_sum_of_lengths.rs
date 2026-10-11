// 给你一个整数数组 arr 和一个整数值 target 。

// 请你在 arr 中找 两个互不重叠的子数组 且它们的和都等于 target 。可能会有多种方案，请你返回满足要求的两个子数组长度和的 最小值 。

// 请返回满足要求的最小长度和，如果无法找到这样的两个子数组，请返回 -1 。

use super::Solution;

use std::{cmp::min, collections::HashMap, i32};

impl Solution {
    pub fn min_sum_of_lengths(mut arr: Vec<i32>, target: i32) -> i32 {
        let mut acc = 0;
        let mut prefix_sum2index: HashMap<i32, i32> = HashMap::new();
        for (i, elem) in arr.iter_mut().enumerate() {
            acc += *elem;
            *elem = acc;
            prefix_sum2index.insert(acc, i as i32);
        }
        prefix_sum2index.insert(0, -1);
        let prefix_sum = arr;
        let n = prefix_sum.len();
        let mut end2start: Vec<Option<usize>> = Vec::new();
        end2start.resize(n, None);
        for end in 0..n {
            end2start[end] = prefix_sum2index.get(&(prefix_sum[end]-target)).map_or(None, |x| Some((*x+1) as usize));
        }
        let mut min_len = Vec::with_capacity(n);
        min_len.resize(n, None);
        min_len[0] = if prefix_sum[0] == target {
            Some(1)
        } else {
            None
        };
        for end in 1..n {
            match end2start[end] {
                Some(start) => {
                    min_len[end] = Some(min((end-start+1) as i32, min_len[end-1].unwrap_or(i32::MAX)));
                }
                None => {
                    min_len[end] = min_len[end-1];
                }
            }
        }
        let mut result = i32::MAX;
        for end in 1..n {
            let Some(start) = end2start[end] else {
                continue;
            };
            if start == 0 {
                continue;
            }
            match min_len[start-1] {
                Some(len) => result = min(
                    (end-start+1) as i32 + len,
                    result,
                ),
                None => continue
            }
        }
        if result == i32::MAX {
            -1
        } else {
            result
        }
    }
}

#[cfg(test)]
mod test {
    use super::Solution;
    #[test]
    fn test1() {
        assert_eq!(Solution::min_sum_of_lengths(vec![3,2,2,4,3], 3), 2);
        assert_eq!(Solution::min_sum_of_lengths(vec![7,3,4,7], 7), 2);
        assert_eq!(Solution::min_sum_of_lengths(vec![4,3,2,6,2,3,4], 6), -1);
    }

    #[test]
    fn test2() {
        assert_eq!(Solution::min_sum_of_lengths(vec![1,2,2,3,2,6,7,2,1,4,8], 5), 4)
    }

    #[test]
    fn test3() {
        assert_eq!(Solution::min_sum_of_lengths(vec![24,1,21,1,4,3,27,7,5,1,12,1,1,43,2,5,4,54,34], 54), 5);
    }
}