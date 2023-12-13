use std::fs;

fn combine_first_and_last(numbers: Vec<i32>) -> i32 {
    if let Some(first) = numbers.first() {
        if let Some(last) = numbers.last() {
            let combined_string = format!("{}{}", first, last);
            if let Ok(combined_number) = combined_string.parse::<i32>() {
                return combined_number;
            }
        }
    }

    // Return a default value if combining fails
    0
}

fn part_one(lines: std::str::Split<'_, char>) {
    let mut points_arr: Vec<i32> = vec![];
    let mut total = 0;

    for line in lines {
        let points: Vec<i32> = line
            .chars()
            .filter_map(|x| x.to_digit(10))
            .map(|x| x as i32)
            .collect();

        points_arr.push(combine_first_and_last(points));
    }

    for point in points_arr {
        total += point;
    }

    println!("Total for part one {}", total);
}

fn part_two(lines: std::str::Split<'_, char>) {
    // let i = 0;
    let mut total = 0;

    for line in lines {
        if line.is_empty() {
            continue;
        }

        let mut digits = Vec::new();

        for (i, c) in line.chars().enumerate() {
            if c.is_digit(10) {
                digits.push(c.to_digit(10).unwrap() as i32);
            }

            let prefixes = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"];

            for (d, val) in prefixes.iter().enumerate() {
                if line.get(i..).unwrap_or("").starts_with(val) {
                    digits.push((d + 1) as i32);
                }
            }
        }

        if !digits.is_empty() {
            total += combine_first_and_last(digits);
        }
    }

    println!("Total for day one part two {}", total);
}

fn main() {
    part_one(fs::read_to_string("./src/bin/day1.input").expect("File not found").split('\n'));
    part_two(fs::read_to_string("./src/bin/day1.input").expect("File not found").split('\n'));
}
