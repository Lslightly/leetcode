/*
 * @lc app=leetcode.cn id=104 lang=rust
 *
 * [104] 二叉树的最大深度
 */

use crate::leetcode::Solution;
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }

    pub fn new_with_some(val: i32) -> Option<Rc<RefCell<Self>>> {
        Some(Rc::new(RefCell::new(TreeNode::new(val))))
    }
}

// @lc code=start
// Definition for a binary tree node.

struct NodeDepthPair {
    node_ref: Rc<RefCell<TreeNode>>,
    depth: i32,
}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut result = 0;
        let mut stack = vec![];
        let Some(cur) = root else {
            return 0;
        };
        stack.push(NodeDepthPair {
            node_ref: cur,
            depth: 1,
        });
        while let Some(cur_pair) = stack.pop() {
            if cur_pair.depth > result {
                result = cur_pair.depth
            }
            for node_option in [
                cur_pair.node_ref.borrow().left.clone(),
                cur_pair.node_ref.borrow().right.clone(),
            ] {
                if let Some(node_ref) = node_option {
                    stack.push(NodeDepthPair {
                        node_ref,
                        depth: cur_pair.depth + 1,
                    });
                }
            }
        }
        result
    }
}
// @lc code=end

#[cfg(test)]
mod test {
    use super::super::*;
    use super::*;
    fn create_tree_nodes(vals: Vec<Option<i32>>) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        vals.iter()
            .map(|v| {
                if let Some(v) = v {
                    TreeNode::new_with_some(*v)
                } else {
                    None
                }
            })
            .collect()
    }
    #[test]
    fn test_1() {
        let mut nodes = create_tree_nodes(parse_tree_array("[3,9,20,null,null,15,7]"));
        nodes[0].clone().unwrap().borrow_mut().left = nodes[1].clone();
        nodes[0].clone().unwrap().borrow_mut().right = nodes[2].clone();
        nodes[2].clone().unwrap().borrow_mut().left = nodes[5].clone();
        nodes[2].clone().unwrap().borrow_mut().right = nodes[6].clone();
        assert_eq!(Solution::max_depth(nodes[0].clone()), 3);
    }
}
