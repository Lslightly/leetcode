use super::Solution;

use std::collections::{HashMap, HashSet};

struct AllowedGroups {
    mask: i8
}

impl AllowedGroups {
    fn new() -> Self {
        Self {
            mask: 0b000,
        }
    }

    fn reserve_seat(&mut self, seat: i32) {
        match seat {
            2|3 => self.mask |= 0b100,
            4|5 => self.mask |= 0b110,
            6|7 => self.mask |= 0b011,
            8|9 => self.mask |= 0b001,
            _ => {}
        }
    }
    fn max_group_num(&self) -> i32 {
        let valid1 = self.mask & 0b100 == 0;
        let valid2 = self.mask & 0b010 == 0;
        let valid3 = self.mask & 0b001 == 0;
        if valid1 && valid3 {
            2
        } else if valid1 || valid2 || valid3 {
            1
        } else {
            0
        }
    }
}



impl Solution {
    pub fn max_number_of_families(n: i32, reserved_seats: Vec<Vec<i32>>) -> i32 {
        let mut row_reserved = HashMap::<i32, HashSet<i32>>::new();
        for seat in reserved_seats {
            let (row, col) = (seat[0], seat[1]);
            if let Some(reserved) = row_reserved.get_mut(&row) {
                reserved.insert(col);
            } else {
                row_reserved.insert(row, HashSet::from([col]));
            }
        }
        let mut result = 0;
        for (_, seats) in &row_reserved {
            let mut allowed = AllowedGroups::new();
            for seat in seats {
                allowed.reserve_seat(*seat);
            }
            result += allowed.max_group_num();
        }
        result += (n-row_reserved.len() as i32)*2;
        result
    }
}

#[cfg(test)]
mod test {
    use crate::leetcode::Solution;

    #[test]
    fn test1() {
        assert_eq!(Solution::max_number_of_families(3, vec![vec![1,2],vec![1,3],vec![1,8],vec![2,6],vec![3,1],vec![3,10]]), 4);
    }
}