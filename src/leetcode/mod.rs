pub struct Solution();

mod bintree_max_depth;
mod find_kth_largest;
mod inorder_traversal;
mod largest_integer;
mod search_range;
mod check_divisibility;
pub mod find_kth_smallest;
mod result_array_3069;
mod max_number_of_families_1386;

/// 解析 LeetCode 层序数组字符串（如 `"[3,9,20,null,null,15,7]"`）为 `Vec<Option<i32>>`。
/// `[]` 或 `""` 返回空 vec；非法 token 直接 panic。
pub fn parse_tree_array(s: &str) -> Vec<Option<i32>> {
    let trimmed = s.trim().trim_matches(['[', ']']).trim();
    if trimmed.is_empty() {
        return vec![];
    }
    trimmed
        .split(',')
        .map(|tok| match tok.trim() {
            "null" => None,
            t => Some(t.parse().expect("invalid tree array token")),
        })
        .collect()
}

pub fn parse_int_array(s: &str) -> Vec<i32> {
    let trimmed = s.trim().trim_matches(['[', ']']).trim();
    if trimmed.is_empty() {
        return vec![];
    }
    trimmed
        .split(',')
        .map(|tok| tok.trim().parse().expect("invalid token"))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse_tree_array() {
        assert_eq!(
            parse_tree_array("[3,9,20,null,null,15,7]"),
            vec![Some(3), Some(9), Some(20), None, None, Some(15), Some(7)]
        );
        assert_eq!(parse_tree_array("[]"), vec![]);
        assert_eq!(parse_tree_array("[null]"), vec![None]);
    }

    #[test]
    fn test_parse_int_array() {
        assert_eq!(parse_int_array("[]"), Vec::<i32>::new());
        assert_eq!(parse_int_array("[5,7,7,8,8,10]"), vec![5, 7, 7, 8, 8, 10])
    }
}
