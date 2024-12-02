use anyhow::Result;
use util::open_input;

fn main() -> Result<()> {
    let f = "day1/day1.txt";
    println!("day1");
    calc_1(f)?;
    calc_2(f)?;
    Ok(())
}

pub fn calc_1(f: &'static str) -> Result<()> {
    let mut sum = 0;

    for line in open_input(f)? {
        let line = line?;
        let mut digits = Vec::new();

        line.chars().for_each(|c| {
            if c.is_ascii_digit() {
                digits.push(c.to_digit(10).unwrap())
            }
        });
        let first = digits.first().unwrap();
        let last = digits.last().unwrap();

        sum += first * 10 + last;
    }
    println!("sum 1: {}", sum);
    Ok(())
}

pub fn calc_2(f: &'static str) -> Result<()> {
    let words = [
        ("one", 1),
        ("two", 2),
        ("three", 3),
        ("four", 4),
        ("five", 5),
        ("six", 6),
        ("seven", 7),
        ("eight", 8),
        ("nine", 9),
    ];

    let mut sum = 0;

    for line in open_input(f)? {
        let line = line?;
        let mut digits = Vec::new();

        line.char_indices().for_each(|(i, c)| {
            if c.is_ascii_digit() {
                digits.push(c.to_digit(10).unwrap())
            } else {
                for word in words {
                    if check_substr_at(line.as_str(), word.0, i) {
                        digits.push(word.1)
                    }
                }
            }
        });
        let first = digits.first().unwrap();
        let last = digits.last().unwrap();

        sum += first * 10 + last;
    }
    println!("sum 2: {}", sum);
    Ok(())
}

fn check_substr_at(s: &str, subs: &str, pos: usize) -> bool {
    if pos + subs.len() > s.len() {
        return false;
    }
    s.chars()
        .skip(pos)
        .zip(subs.chars())
        .all(|(c1, c2)| c1 == c2)
}
