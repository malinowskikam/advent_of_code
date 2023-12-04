use util::open_input;
use anyhow::Result;

fn main() -> Result<()> {
    let f = "day4/day4.txt";
    println!("day4");
    calc_1(f)?;
    calc_2(f)?;
    Ok(())
}

pub fn calc_1(f: &'static str) -> Result<()> {
    let mut sum = 0;

    for line in open_input(f)? {
        let line = line?;

        let sections: Vec<&str> = line.split_terminator(&[':', '|']).collect();
        assert!(sections.len() == 3);

        let winning_ns: Vec<&str> = sections[1].trim().split_terminator(" ").filter(|n| !n.is_empty()).collect();
        let count = sections[2].trim().split_terminator(" ").filter(|n| !n.is_empty() && winning_ns.contains(n)).count() as u32;

        if count > 0 {
            sum += 2u32.pow(count-1);
        }
    }

    println!("sum 1: {}", sum);
    Ok(())
}

pub fn calc_2(f: &'static str) -> Result<()> {
    let mut sum = vec![0];
    
    for (n, line) in open_input(f)?.enumerate() {
        let line = line?;

        // add main card
        if sum.len() < n + 1 {
            sum.push(1);
        } else {
            sum[n] += 1
        }

        let sections: Vec<&str> = line.split_terminator(&[':', '|']).collect();
        assert!(sections.len() == 3);

        let winning_ns: Vec<&str> = sections[1].trim().split_terminator(" ").filter(|n| !n.is_empty()).collect();

        let count = sections[2].trim().split_terminator(" ").filter(|n| !n.is_empty() && winning_ns.contains(n)).count();
        
        //add extra cards
        for i in 0..count {
            if sum.len() < n + 2 + i {
                sum.push(sum[n]);
            } else {
                sum[n + i + 1] += sum[n]
            }
        }
    }
    
    println!("day4 sum: {:?}", sum.iter().sum::<usize>());
    Ok(())
}
