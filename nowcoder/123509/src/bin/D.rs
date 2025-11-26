use std::{io::{self}, str::FromStr};

/**
 * 000111000
 * 
 * 0 3 6 9
 * start[i]             start[i+k-1]
 * 
 * (start[i+1]-start[i])*(start[i+k]-start[i+k-1])
 * 
 */

fn gen_start_idxs(s: &String) -> Vec<usize> {
    let mut vec = Vec::<usize>::new();
    let mut last_one = s.as_bytes()[0];
    vec.push(0);
    for (i, ch) in s.as_bytes().iter().enumerate() {
        if *ch != last_one {
            last_one = *ch;
            vec.push(i);
        }
    }
    vec.push(s.len());
    vec
}

fn k_diff_sub_strs(s: &String, k: usize) -> usize {
    let start_idxs = gen_start_idxs(s);
    if k > start_idxs.len()-1 {
        return 0;
    }
    if k == 1 {
        let mut chances = 0;
        for i in 0..start_idxs.len()-1 {
            let unit_len = start_idxs[i+1]-start_idxs[i];
            chances += unit_len*(unit_len+1)/2;
        }
        return chances;
    }
    let mut chances = 0;
    for i in 0..=start_idxs.len()-1-k {
        chances += (start_idxs[i+1]-start_idxs[i])*(start_idxs[i+k]-start_idxs[i+k-1])
    }
    chances
}

fn main() {
    let mut line = String::new();
    let stdin = io::stdin();
    stdin.read_line(&mut line).expect("读取失败");
    let parts: Vec<&str> = line.trim().split_whitespace().collect();
    let k = if parts.len() >= 2 {
        match usize::from_str(parts[1]) {
            Ok(val) => val,
            Err(_) => {
                println!("解析失败");
                0
            }
        }
    } else {
        println!("输入格式错误");
        0
    };
    line.clear();
    stdin.read_line(&mut line).expect("读取失败");
    println!("{}", k_diff_sub_strs(&line.trim_end().to_string(), k));
}

#[cfg(test)]
mod test {
    use crate::k_diff_sub_strs;

    #[test]
    fn test_simple() {
        assert_eq!(k_diff_sub_strs(&"001100".to_string(), 2), 8);
        assert_eq!(k_diff_sub_strs(&"000".to_string(), 2), 0);
        assert_eq!(k_diff_sub_strs(&"000".to_string(), 1), 6);
        assert_eq!(k_diff_sub_strs(&"010".to_string(), 4), 0);
        assert_eq!(k_diff_sub_strs(&"010".to_string(), 2), 2);
        assert_eq!(k_diff_sub_strs(&"010".to_string(), 1), 3);
        assert_eq!(k_diff_sub_strs(&"0011000".to_string(), 3), 6);
        assert_eq!(k_diff_sub_strs(&"00100111".to_string(), 3), 7);
        assert_eq!(k_diff_sub_strs(&"000111".to_string(), 1), 6+3+3);
    }
}
