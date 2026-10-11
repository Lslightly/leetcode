use std::{collections::HashMap, str::Chars};

struct Row([i32;4]);

impl Row {
    fn new() -> Self { Self([0;4]) }
    fn toggle(&mut self, pos: usize) -> i32 {
        if self.0[pos] != 0 {
            let time = self.0[pos];
            self.0[pos] = 0;
            return -time;
        }
        let time = match pos {
            0 => {
                if self.0[1] != 0 {
                    3
                } else {
                    1
                }
            },
            3 => {
                if self.0[2] != 0 {
                    3
                } else {
                    1
                }
            },
            _ => 1
        };
        self.0[pos] = time;
        time
    }
}

struct InOrder {
    n: i32,
    time: i32,
    rows: HashMap<i32, Row>
}

impl InOrder {
    fn new(n: i32) -> Self {
        Self {
            n: n,
            time: 0,
            rows: HashMap::new(),
        }
    }
    fn toggle(&mut self, req: Request) {
        self.time += self.rows.entry(req.row).or_insert(Row::new()).toggle(req.pos)
    }

    fn get_time(&self) -> i32 {
        self.time
    }
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Hash, Clone, Copy)]
struct Request {
    row: i32,
    pos: usize,
}

impl<'a> From<Chars<'a>> for Request {
    fn from(mut value: Chars) -> Self {
        let first = value.next().unwrap();
        let pos = u32::try_from(first).unwrap() - u32::try_from('A').unwrap();
        let row: i32 = value.as_str().parse().unwrap();
        Self {
            row,
            pos: pos as usize,
        }
    }
}

struct OptOrder {
    reqs_cnt: HashMap<Request, i32>
}

impl OptOrder {
    fn new() -> Self {
        Self {
            reqs_cnt: HashMap::new()
        }
    }

    fn toggle(&mut self, req: Request) {
        *self.reqs_cnt.entry(req).or_insert(0) += 1;
    }
    
    fn get_time(&self) -> i32 {
        let mut time: i32 = 0;
        for (_, &cnt) in self.reqs_cnt.iter() {
            time += cnt%2;
        }
        time
    }
}

struct BusBooking {
    in_order: InOrder,
    opt_order: OptOrder,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl BusBooking {

    fn new(n: i32) -> Self {
        Self {
            in_order: InOrder::new(n),
            opt_order: OptOrder::new()
        }
    }
    
    fn toggle(&mut self, seat: String) {
        let req: Request = seat.chars().into();
        self.in_order.toggle(req);
        self.opt_order.toggle(req);
    }
    
    fn get_total_time(&self) -> i32 {
        self.in_order.get_time()
    }
    
    fn get_min_time(&self) -> i32 {
        self.opt_order.get_time()
    }
}

/**
 * Your BusBooking object will be instantiated and called as such:
 * let obj = BusBooking::new(n);
 * obj.toggle(seat);
 * let ret_2: i32 = obj.get_total_time();
 * let ret_3: i32 = obj.get_min_time();
 */
#[cfg(test)]
mod tests {
    use super::BusBooking;
    #[test]
    fn test1() {
        let mut obj = BusBooking::new(20);
        obj.toggle("B3".into());
        obj.toggle("A3".into());
        assert_eq!(obj.get_total_time(), 4);
        assert_eq!(obj.get_min_time(), 2);
        obj.toggle("C3".into());
        obj.toggle("D3".into());
        assert_eq!(obj.get_total_time(), 8);
        assert_eq!(obj.get_min_time(), 4);
    }
}
