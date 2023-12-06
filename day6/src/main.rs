use anyhow::{anyhow, Result};
use util::open_input;

fn main() -> Result<()> {
    let f = "day6/day6.txt";
    println!("day6");
    calc_1(f)?;
    calc_2(f)?;
    Ok(())
}

struct Race {
    time: i64,
    distance: i64,
}

impl TryFrom<(&str, &str)> for Race {
    type Error = anyhow::Error;

    fn try_from((time_s, distance_s): (&str, &str)) -> std::prelude::v1::Result<Self, Self::Error> {
        Ok(Race {
            time: time_s.parse()?,
            distance: distance_s.parse()?,
        })
    }
}

pub fn calc_1(f: &'static str) -> Result<()> {
    let mut prod = 1;
    let mut lines_iter = open_input(f)?;

    let line1 = lines_iter.next().ok_or(anyhow!("no first line"))??;
    let line2 = lines_iter.next().ok_or(anyhow!("no second line"))??;

    let records: Vec<Race> = line1
        .split(" ")
        .filter(|s| *s != "")
        .skip(1)
        .zip(line2.split(" ").filter(|s| *s != "").skip(1))
        .map(|tup| tup.try_into().unwrap())
        .collect();
    for record in records {
        let mut ways = 0;
        for t_acc in 1..record.time {
            if record.distance < (record.time - t_acc) * t_acc {
                ways += 1;
            };
        };
        prod *= ways
    };

    println!("sum 1: {}", prod);
    Ok(())
}

pub fn calc_2(f: &'static str) -> Result<()> {
    let mut sum = 0;
    let mut lines_iter = open_input(f)?;

    let line1 = lines_iter.next().ok_or(anyhow!("no first line"))??;
    let line2 = lines_iter.next().ok_or(anyhow!("no second line"))??;

    let record = Race {
        time: line1.split(" ").filter(|s| *s != "").skip(1).collect::<Vec<&str>>().join("").parse()?,
        distance: line2.split(" ").filter(|s| *s != "").skip(1).collect::<Vec<&str>>().join("").parse()?
    };

    for t_acc in 1..record.time {
        if record.distance < (record.time - t_acc) * t_acc {
            sum += 1;
        }
    };

    println!("sum 1: {}", sum);
    Ok(())
}
