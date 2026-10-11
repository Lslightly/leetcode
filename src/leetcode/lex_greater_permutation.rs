use std::{array::repeat, collections::BTreeMap, iter, str::Chars};

use super::Solution;

impl Solution {
    fn concat_chars(mut result_str: String, char_cnt_map: &BTreeMap<char, usize>) -> String {
        for (&k, &v) in char_cnt_map.iter() {
            result_str.extend(iter::repeat_n(k, v));
        }
        result_str
    }
    fn find_min_permutation<'a>(result_str: String, cur_tch: char, k: char, cur_s: Option<(&char, &'a mut usize)>, char_cnt_map: &'a BTreeMap<char, usize>, mut target_iter: Chars) -> String {

    }
    pub fn lex_greater_permutation(s: String, target: String) -> String {
        let empty_string = String::from("");
        let mut char_cnt_map: BTreeMap<char, usize> = BTreeMap::new();
        for ch in s.chars() {
            if let Some(cnt) = char_cnt_map.get_mut(&ch) {
                *cnt += 1;
            } else {
                char_cnt_map.insert(ch, 1);
            }
        }
        let mut target_iter = target.chars();
        let mut char_cnt_map_iter = char_cnt_map.iter_mut().rev();
        let mut cur_s = None;
        let mut result_chars = String::new();
        while let Some(tch) = target_iter.next() {
            if cur_s.is_none() {
                cur_s = char_cnt_map_iter.next();
            }
            if let Some((k, v)) = cur_s {
                if *k > tch {
                    return Solution::find_min_permutation(result_chars, tch, *k, cur_s, &char_cnt_map, target_iter);
                } else if *k < tch {
                    return empty_string;
                }
                *v -= 1;
                result_chars.push(*k);
                if *v == 0 {
                    cur_s = None;
                }
            }
        }
        Solution::concat_chars(result_chars, &char_cnt_map)
    }
}

#[cfg(test)]
mod test {
    use crate::leetcode::Solution;

    #[test]
    fn test1() {
        assert_eq!(Solution::lex_greater_permutation("abc".into(), "bba".into()), "bca");
        assert_eq!(Solution::lex_greater_permutation("abb".into(), "bba".into()), "bba");
        assert_eq!(Solution::lex_greater_permutation("aab".into(), "bba".into()), "");
    }
}