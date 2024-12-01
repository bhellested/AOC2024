use std::fs;
use std::collections::HashMap;
use std::cmp::Reverse;
fn main() {
    let data = fs::read_to_string("../inputs/day1.txt").expect("Unable to read file");
    let lines = data.split("\n");
    let mut left_vec: Vec<i32> = Vec::new();
    let mut right_counts: HashMap<i32, i32> = HashMap::new();
    for pair in lines {
        let numbers = pair.split("   ").map(|s| s.parse::<i32>().unwrap()).collect::<Vec<i32>>();
        left_vec.push(numbers[0]);
        right_counts.insert(numbers[1], right_counts.get(&numbers[1]).unwrap_or(&0) + 1);
    }
    let mut diff =0;
    for i in 0..left_vec.len() {
        if right_counts.contains_key(&left_vec[i]) {
            diff += left_vec[i] * right_counts.get(&left_vec[i]).unwrap();

        }
    }
    println!("{}", diff);
}