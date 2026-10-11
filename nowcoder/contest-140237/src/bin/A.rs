use std::{ops::{Index, IndexMut}, slice::Iter};


struct Matrix<T>
where T: Default + Copy
{
    data: Vec<Vec<T>>,
    n: usize, // row num
    m: usize, // col num
}

impl<T> Matrix<T> 
where T: Default + Copy
{
    fn new(n: usize, m: usize) -> Self {
        let mut row = Vec::new();
        row.resize(m as usize, T::default());
        let mut data = Vec::new();
        data.resize(n as usize, row.iter().copied().collect());
        Self {
            data,
            n,
            m
        }
    }
}

impl<T> Index<usize> for Matrix<T>
where T: Default + Copy
{
    type Output = [T];
    fn index(&self, index: usize) -> &Self::Output {
        &self.data[index]
    }
}

impl<T> IndexMut<usize> for Matrix<T>
where T: Default + Copy
{
    fn index_mut(&mut self, index: usize) -> &mut Self::Output {
        &mut self.data[index]
    }
}

impl From<&Matrix<bool>> for Matrix<State> {
    fn from(raw: &Matrix<bool>) -> Self {
        let mut state = Self::new(raw.n, raw.m);
        for (row_idx, row) in raw.data.iter().enumerate() {
            for (col_idx, col) in row.iter().enumerate() {
                state.data[row_idx][col_idx] = (*col).into();
            }
        }
        state
    }
}

#[derive(Clone, Copy, PartialEq, Eq)]
enum State {
    AllOne,
    AllZero,
    Mixed,
}

impl From<bool> for State {
    fn from(value: bool) -> Self {
        if value {
            Self::AllOne
        } else {
            Self::AllZero
        }
    }
}

impl Default for State {
    fn default() -> Self {
        Self::AllOne
    }
}

struct CornerStateMatrix {
    matrix: Matrix<State>,
    col_bound_iter: fn(m: &Matrix<State>, col_bound: usize) -> Box<dyn Iterator<Item = usize>>,
    zero_row_match: fn(m: &Matrix<State>, col: usize) -> bool,
    zero_col: fn(m: &mut Matrix<State>, col: usize) -> &mut State,
    row_iter: fn(m: &Matrix<State>) -> Box<dyn Iterator<Item = usize>>,
    row_bound_match: fn(m: &Matrix<State>, row: usize) -> bool,
    two_match: fn(m: &Matrix<State>, row: usize, col: usize) -> bool,
    rest_col: fn(m: usize, col_bound: usize) -> Box<dyn Iterator<Item = usize>>,
    rest_row: fn(n: usize, row_bound: usize) -> Box<dyn Iterator<Item = usize>>,
}

impl CornerStateMatrix {
    fn draw_rest_as_mixed(&mut self, row_bound: usize, col_bound: usize) {
        for row in (self.rest_row)(self.matrix.n, row_bound) {
            for col in (self.rest_col)(self.matrix.m, col_bound) {
                self.matrix[row][col] = State::Mixed
            }
        }
    }

    fn fill(&mut self) {
        let mut col_bound = self.matrix.m;
        for col in (self.col_bound_iter)(&self.matrix, col_bound) {
            if (self.zero_row_match)(&self.matrix, col) {
                *(self.zero_col)(&mut self.matrix, col) = State::Mixed;
                self.draw_rest_as_mixed(0, col);
                col_bound = col;
                break;
            }
        }
        'outer: for row in (self.row_iter)(&self.matrix) {
            if (self.row_bound_match)(&self.matrix, row) {
                self.draw_rest_as_mixed(row, 0);
                break;
            }
            for col in (self.col_bound_iter)(&self.matrix, col_bound) {
                if (self.two_match)(&self.matrix, row, col) {
                    self.draw_rest_as_mixed(row, col);
                    col_bound = col;
                    continue 'outer;
                }
            }
        }
    }
}

fn left_col_bound(m: usize) -> usize {
    m
}

fn right_col_bound(m: usize) -> usize {
    0
}

fn left_col_iter(m: &Matrix<State>, col_bound: usize) -> Box<dyn Iterator<Item = usize>> {
    Box::new(1..col_bound)
}

fn right_col_iter(m: &Matrix<State>, col_bound: usize) -> Box<dyn Iterator<Item = usize>> {
    Box::new((col_bound..m.m-1).rev())
}

fn up_zero_col(m: &mut Matrix<State>, col: usize) -> &mut State {
    &mut m[0][col]
}

fn down_zero_col(m: &mut Matrix<State>, col: usize) -> &mut State {
    let end_row = m.n-1;
    &mut m[end_row][col]
}

fn up_row_iter(m: &Matrix<State>) -> Box<dyn Iterator<Item = usize>> {
    Box::new(1..m.n)
}

fn down_row_iter(m: &Matrix<State>) -> Box<dyn Iterator<Item = usize>> {
    Box::new((0..m.n-1).rev())
}

fn left_up_match(m: &Matrix<State>, col: usize) -> bool {
    m[0][col-1] != m[0][col]
}

fn left_down_match(m: &Matrix<State>, col: usize) -> bool {
    let last_row = m.n-1;
    m[last_row][col-1] != m[last_row][col]
}

fn right_up_match(m: &Matrix<State>, col: usize) -> bool {
    m[0][col+1] != m[0][col]
}

fn right_down_match(m: &Matrix<State>, col: usize) -> bool {
    let last_row = m.n-1;
    m[last_row][col+1] != m[last_row][col]
}

fn left_up_last_row_bound_match(m: &Matrix<State>, row: usize) -> bool {
    m[row-1][0] != m[row][0]
}

fn left_down_last_row_bound_match(m: &Matrix<State>, row: usize) -> bool {
    m[row+1][0] != m[row][0]
}

fn right_up_last_row_bound_match(m: &Matrix<State>, row: usize) -> bool {
    m[row-1][m.m-1] != m[row][m.m-1]
}

fn right_down_last_row_bound_match(m: &Matrix<State>, row: usize) -> bool {
    m[row+1][m.m-1] != m[row][m.m-1]
}

fn left_up_two_match(m: &Matrix<State>, row: usize, col: usize) -> bool {
    m[row-1][col] != m[row][col-1]
    || m[row-1][col] != m[row][col]
}

fn left_down_two_match(m: &Matrix<State>, row: usize, col: usize) -> bool {
    m[row+1][col] != m[row][col-1]
    || m[row+1][col] != m[row][col]
}

fn right_up_two_match(m: &Matrix<State>, row: usize, col: usize) -> bool {
    m[row-1][col] != m[row][col+1]
    || m[row-1][col] != m[row][col]
}

fn right_down_two_match(m: &Matrix<State>, row: usize, col: usize) -> bool {
    m[row+1][col] != m[row][col+1]
    || m[row+1][col] != m[row][col]
}

fn left_rest_col(m: usize, col_bound: usize) -> Box<dyn Iterator<Item = usize>> {
    Box::new(col_bound..m)
}

fn right_rest_col(m: usize, col_bound: usize) -> Box<dyn Iterator<Item = usize>> {
    Box::new((0..=col_bound).rev())
}

fn up_rest_row(n: usize, row_bound: usize) -> Box<dyn Iterator<Item = usize>> {
    Box::new(row_bound..n)
}

fn down_rest_row(n: usize, row_bound: usize) -> Box<dyn Iterator<Item = usize>> {
    Box::new((0..=row_bound).rev())
}

struct StateMatrix {
    left_up: CornerStateMatrix,
    left_down: CornerStateMatrix,
    right_up: CornerStateMatrix,
    right_down: CornerStateMatrix,
}

impl StateMatrix {
    fn new(matrix: &Matrix<bool>) -> Self {
        Self {
            left_up: CornerStateMatrix { matrix: matrix.into(), col_bound_iter: left_col_iter, zero_row_match: left_up_match, zero_col: up_zero_col, row_iter: up_row_iter, row_bound_match: left_up_last_row_bound_match, two_match: left_up_two_match, rest_col: left_rest_col, rest_row: up_rest_row },
            left_down: CornerStateMatrix { matrix: matrix.into(), col_bound_iter: left_col_iter, zero_row_match: left_down_match, zero_col: down_zero_col, row_iter: down_row_iter, row_bound_match: left_down_last_row_bound_match, two_match: left_down_two_match, rest_col: left_rest_col, rest_row: down_rest_row },
            right_up: CornerStateMatrix { matrix: matrix.into(), col_bound_iter: right_col_iter, zero_row_match: right_up_match, zero_col: up_zero_col, row_iter: up_row_iter, row_bound_match: right_up_last_row_bound_match, two_match: right_up_two_match, rest_col: right_rest_col, rest_row: up_rest_row },
            right_down: CornerStateMatrix { matrix: matrix.into(), col_bound_iter: right_col_iter, zero_row_match: right_down_match, zero_col: down_zero_col, row_iter: down_row_iter, row_bound_match: right_down_last_row_bound_match, two_match: right_down_two_match, rest_col: right_rest_col, rest_row: down_rest_row },
        }
    }
}

fn number_of_valid_tuples(m: Matrix<bool>) -> i32;

fn main() {

}