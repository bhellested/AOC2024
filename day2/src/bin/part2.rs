use std::fs;
fn main() {
    let data = fs::read_to_string("../inputs/day2.txt").expect("Unable to read file");
    let lines = data.split("\n");
    let mut good_reports=0;
    for line in lines {
        let mut numbers = line.split(" ").map(|s| s.parse::<i32>().unwrap()).collect::<Vec<i32>>();
        if is_good_report(&numbers, true) {
            good_reports+=1;
        }
    }
    println!("{}", good_reports);
}

fn is_good_report(numbers: &Vec<i32>, canRecurse: bool) -> bool {
    let mut increasing = false;
    if numbers[0] < numbers[1] {
        increasing = true;
    }
    for i in 0..numbers.len()-1 {
        let diff = numbers[i+1] - numbers[i];
        if increasing {
            if diff < 1 || diff > 3 {
                if canRecurse {
                    //lets do every possible removal here
                    for j in 0..numbers.len() {
                        let mut new_line = numbers.clone();
                        new_line.remove(j);
                        if is_good_report(&new_line, false) {
                            return true;
                        }
                    }
                    return false;
                }
                return false;
            }
        } else {
            if diff > -1 || diff < -3 {
                if canRecurse {
                    for j in 0..numbers.len() {
                        let mut new_line = numbers.clone();
                        new_line.remove(j);
                        if is_good_report(&new_line, false) {
                            return true;
                        }
                    }
                    return false;
                }
                return false;
            }
        }
    }
    return true;
}