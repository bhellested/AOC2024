use std::fs;

fn main() {
    let data = fs::read_to_string("../inputs/day2.txt").expect("Unable to read file");
    let lines = data.split("\n");
    let mut good_reports=0;
    for line in lines {
        if is_good_report(line) {
            good_reports+=1;
        }
    }
    println!("{}", good_reports);
}

fn is_good_report(line: &str) -> bool {
    let numbers = line.split(" ").map(|s| s.parse::<i32>().unwrap()).collect::<Vec<i32>>();
    let mut increasing = false;
    //check first against second
    if numbers[0] < numbers[1] {
        increasing = true;
    }
    for i in 0..numbers.len()-1 {
        let diff = numbers[i+1] - numbers[i];
        if increasing {
            if diff < 1 || diff > 3 {
                return false;
            }
        } else {
            if diff > -1 || diff < -3 {
                return false;
            }
        }
    }
    return true;
}