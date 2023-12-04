use std::cmp::min;
use util::open_input;
use anyhow::Result;

fn main() -> Result<()> {
    let f = "day3/day3.txt";
    println!("day3");
    calc_1(f)?;
    calc_2(f)?;
    Ok(())
}

pub fn calc_1(f: &'static str) -> Result<()> {
    let mut sum = 0u32;

    let mut numbers: Vec<(String, usize, usize)> = vec![];
    let mut signs: Vec<(usize, usize)> = vec![];
    let mut num_buff: Vec<char> = vec![];

    for (row, line) in open_input(f)?.enumerate() {
        let line = line?;

        for (col, c) in (&line).chars().enumerate() {
            if c.is_ascii_digit() {
                num_buff.push(c);
            } else if c != '.' {
                signs.push((row, col));
                if !num_buff.is_empty() {
                    numbers.push((num_buff.iter().collect(), row, col - num_buff.len()));
                    num_buff.clear();
                }
            } else {
                if !num_buff.is_empty() {
                    numbers.push((num_buff.iter().collect(), row, col - num_buff.len()));
                    num_buff.clear();
                }
            }
        }

        if !num_buff.is_empty() {
            numbers.push((num_buff.iter().collect(), row, line.len() - num_buff.len()));
            num_buff.clear();
        }
    }

    'number: for number in &numbers {
        let num_len = number.0.len();
        for sign in &signs {
            if sign.0 >= number.1 - min(number.1, 1)
                && sign.0 <= number.1 + 1
                && sign.1 >= number.2 - min(number.2, 1)
                && sign.1 <= number.2 + num_len
            {
                // sing row >=/<= number row -/+1
                // sign col >=/<= number start col -1/end col +1
                sum += number.0.parse::<u32>()?;
                continue 'number;
            }
        }
    }

    println!("sum 1: {}", sum);
    Ok(())
}

pub fn calc_2(f: &'static str) -> Result<()> {
    let mut sum = 0u32;

    let mut numbers: Vec<(String, usize, usize)> = vec![];
    let mut signs: Vec<(usize, usize)> = vec![];
    let mut num_buff: Vec<char> = vec![];

    for (row, line) in open_input(f)?.enumerate() {
        let line = line?;

        for (col, c) in (&line).chars().enumerate() {
            if c.is_ascii_digit() {
                num_buff.push(c);
            } else if c == '*' {
                signs.push((row, col));
                if !num_buff.is_empty() {
                    numbers.push((num_buff.iter().collect(), row, col - num_buff.len()));
                    num_buff.clear();
                }
            } else {
                if !num_buff.is_empty() {
                    numbers.push((num_buff.iter().collect(), row, col - num_buff.len()));
                    num_buff.clear();
                }
            }
        }

        if !num_buff.is_empty() {
            numbers.push((num_buff.iter().collect(), row, line.len() - num_buff.len()));
            num_buff.clear();
        }
    }

    'sign: for sign in &signs {
        let mut gear_ratio = 1u32;
        let mut nums_found = 0u32;

        for number in &numbers {
            if sign.0 >= number.1 - min(number.1, 1)
                && sign.0 <= number.1 + 1
                && sign.1 >= number.2 - min(number.2, 1)
                && sign.1 <= number.2 + number.0.len()
            {

                // sing row >=/<= number row -/+1
                // sign col >=/<= number start col -1/end col +1
                gear_ratio *= number.0.parse::<u32>()?;
                nums_found += 1;
                if nums_found > 2 {
                    continue 'sign;
                }
            }
        }

        if nums_found == 2 {
            sum += gear_ratio;
        }
    }
    println!("sum 2: {}", sum);
    Ok(())
}
