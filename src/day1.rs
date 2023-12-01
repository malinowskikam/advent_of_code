use crate::util::open_input;
use anyhow::Result;

pub fn calc_day_1() -> Result<()> {
    let mut sum = 0;
    let mut count = 0;

    for line in open_input("day1.txt")? {
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
        count += 1
    }
    println!("day12 sum: {}, count {}", sum, count);
    Ok(())
}

pub fn calc_day_1_2() -> Result<()> {
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
    let mut count = 0;

    for line in open_input("day1.txt")? {
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
        count += 1
    }
    println!("day1_2 sum: {}, count {}", sum, count);
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

#[cfg(test)]
mod tests {
    use super::{calc_day_1, calc_day_1_2};

    #[test]
    fn test_day_1() {
        assert!(calc_day_1().is_ok());
    }

    #[test]
    fn test_day_1_2() {
        assert!(calc_day_1_2().is_ok());
    }
}
