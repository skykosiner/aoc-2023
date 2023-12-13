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
        if line.trim().is_empty() {
            continue;
        }

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

fn main() {
    // Part one
    part_one(fs::read_to_string("./src/bin/day1.test").expect("File not found").split('\n'));
}
