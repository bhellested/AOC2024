use std::cmp::Reverse;
use std::collections::BinaryHeap;
use std::fs;
fn main() {
    let data = fs::read_to_string("../inputs/day1.txt").expect("Unable to read file");
    let lines = data.split("\n");
    //create two heaps
    let mut min_heap_1: BinaryHeap<i32> = BinaryHeap::new();
    let mut min_heap_2: BinaryHeap<i32> = BinaryHeap::new();
    for pair in lines {
        let numbers: Vec<i32> = pair
            .split("   ")
            .map(|s: &str| s.parse::<i32>().unwrap())
            .collect::<Vec<i32>>();

        min_heap_1.push(numbers[0]);
        min_heap_2.push(numbers[1]);
    }
    let mut diff = 0;
    while min_heap_1.len() > 0 && min_heap_2.len() > 0 {
        let top_1 = min_heap_1.pop().unwrap();
        let top_2 = min_heap_2.pop().unwrap();
        diff += (top_1 - top_2).abs();
    }
    println!("{}", diff);
}
