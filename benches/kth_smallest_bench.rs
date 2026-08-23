use leetcode_rs::leetcode::Solution;

use criterion::{criterion_group, criterion_main, Criterion};

fn bench(c: &mut Criterion) {
    c.bench_function(
        "kth_smallest_bench",
        |b| b.iter(|| {
            Solution::find_kth_smallest(vec![5,25,23,16,7,8,10,6,11,15], 946326769)
        })
    );
}

criterion_group!(benches, bench);
criterion_main!(benches);
