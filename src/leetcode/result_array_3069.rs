use super::Solution;

impl Solution {
    pub fn result_array(nums: Vec<i32>) -> Vec<i32> {
        let (mut arr1, mut arr2) = (vec![nums[0]], vec![nums[1]]);
        for n in &nums[2..] {
            if arr1.last().expect("shouldn't be none") > arr2.last().expect("shouldn't be none") {
                arr1.push(*n);
            } else {
                arr2.push(*n);
            }
        }
        arr1.append(&mut arr2);
        arr1
    }
}