use crate::leetcode::Solution;

impl Solution {
    /// get count of overlap of img1[0..=n-1-y][0..=n-1-x] and img2[y..=n][x..=n]
    fn overlap_cnt_right_down(x: usize, y: usize, img1: &Vec<Vec<i32>>, img2: &Vec<Vec<i32>>) -> i32 {
        let n = img1.len();
        let mut count = 0;
        for ky in 0..=n-1-y {
            for kx in 0..=n-1-x {
                if img1[ky][kx] == 1 && img2[y+ky][x+kx] == 1 {
                    count += 1;
                }
            }
        }
        count
    }
    fn overlap_cnt_right_up(x: usize, y: usize, img1: &Vec<Vec<i32>>, img2: &Vec<Vec<i32>>) -> i32 {
        let n = img1.len();
        let mut count = 0;
        for ky in 0..=y {
            for kx in 0..=n-1-x {
                if img1[ky][x+kx] == 1 && img2[n-1-y+ky][kx] == 1 {
                    count += 1;
                }
            }
        }
        count
    }
    pub fn largest_overlap(img1: Vec<Vec<i32>>, img2: Vec<Vec<i32>>) -> i32 {
        let mut largest_cnt = 0;
        let n = img1.len();
        let mut candiates = vec![];
        candiates.reserve(5);
        for y in 0..n {
            for x in 0..n {
                candiates.clear();
                candiates.push(largest_cnt as i32);
                if (n-y)*(n-x) > largest_cnt as usize {
                    candiates.push(Solution::overlap_cnt_right_down(x, y, &img1, &img2));
                    candiates.push(Solution::overlap_cnt_right_down(x, y, &img2, &img1));
                }
                if (y+1)*(n-x) > largest_cnt as usize {
                    candiates.push(Solution::overlap_cnt_right_up(x, y, &img1, &img2));
                    candiates.push(Solution::overlap_cnt_right_up(x, y, &img2, &img1));
                }
                largest_cnt = *candiates.iter().max().expect("shouldn't panic");
            }
        }
        largest_cnt
    }
}

#[cfg(test)]
mod tests {
    use crate::leetcode::Solution;

    #[test]
    fn test1() {
        let img1 = vec![vec![1,1,0],vec![0,1,0],vec![0,1,0]];
        let img2 = vec![vec![0,0,0],vec![0,1,1],vec![0,0,1]];
        assert_eq!(Solution::largest_overlap(img1, img2), 3)
    }

    #[test]
    fn test_move1() {
        let img1 = vec![vec![1,1,0],vec![0,1,0],vec![0,1,0]];
        let img2 = vec![vec![0,0,0],vec![0,1,1],vec![0,0,1]];
        assert_eq!(Solution::overlap_cnt_right_down(1, 0, &img1, &img2), 2);
    }
}
