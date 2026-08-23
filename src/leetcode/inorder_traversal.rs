/*
 * @lc app=leetcode.cn id=94 lang=rust
 *
 * [94] 二叉树的中序遍历
 */
// Definition for a binary tree node.

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

    pub fn new_some(val: i32) -> Option<Rc<RefCell<Self>>> {
        Some(Rc::new(RefCell::new(Self::new(val))))
    }
}

// @lc code=start

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut stack = vec![];
        let mut cur = root;
        let mut res = vec![];
        while cur.is_some() || !stack.is_empty() {
            while let Some(node) = cur {
                stack.push(node.clone());
                cur = node.borrow().left.clone();
            }
            let mid = stack.pop().expect("stack shouldn' contain None value");
            res.push(mid.borrow().val);
            if mid.borrow().right.is_some() {
                cur = mid.borrow().right.clone();
            }
        }
        res
    }
}

// @lc code=end

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_1() {
        let n1 = TreeNode::new_some(1);
        let n2 = TreeNode::new_some(2);
        let n3 = TreeNode::new_some(3);
        n1.clone().unwrap().borrow_mut().right = n2.clone();
        n2.unwrap().borrow_mut().left = n3.clone();
        assert_eq!(Solution::inorder_traversal(n1), vec![1, 3, 2]);
    }
}
