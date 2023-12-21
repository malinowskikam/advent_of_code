use anyhow::{anyhow, Result};
use util::open_input;

const RED_LIMIT: u32 = 12;
const GREEN_LIMIT: u32 = 13;
const BLUE_LIMIT: u32 = 14;

fn main() -> Result<()> {
    let f = "day2/day2.txt";
    println!("day2");
    calc_1(f)?;
    calc_2(f)?;
    Ok(())
}

pub fn calc_1(f: &'static str) -> Result<()> {
    let mut sum = 0;

    for line in open_input(f)? {
        let line = line?;
        let sections: Vec<&str> = line.split(':').collect();

        let game_id_section: Vec<&str> = sections[0].split(' ').collect();
        let game_id: u32 = game_id_section[1].parse::<u32>()?;

        let games: Vec<&str> = sections[1].split(';').collect();
        let mut valid = true;
        for game in games {
            let game: Vec<&str> = game[1..].split(", ").collect();
            for pick in game {
                let pick_elements: Vec<&str> = pick.split(' ').collect();
                let c: u32 = pick_elements[0].parse()?;
                match pick_elements[1] {
                    "red" => {
                        if c > RED_LIMIT {
                            valid = false;
                            break;
                        }
                    }
                    "green" => {
                        if c > GREEN_LIMIT {
                            valid = false;
                            break;
                        }
                    }
                    "blue" => {
                        if c > BLUE_LIMIT {
                            valid = false;
                            break;
                        }
                    }
                    _ => return Err(anyhow!("Invalid color")),
                }
            }

            if !valid {
                break;
            }
        }

        if valid {
            sum += game_id;
        }
    }

    println!("sum 1: {}", sum);
    Ok(())
}

pub fn calc_2(f: &'static str) -> Result<()> {
    let mut sum = 0;

    for line in open_input(f)? {
        let line = line?;
        let sections: Vec<&str> = line.split(':').collect();

        let mut min_red = 0u32;
        let mut min_green = 0u32;
        let mut min_blue = 0u32;

        let games: Vec<&str> = sections[1].split(';').collect();
        for game in games {
            let game: Vec<&str> = game[1..].split(", ").collect();
            for pick in game {
                let pick_elements: Vec<&str> = pick.split(' ').collect();
                let c: u32 = pick_elements[0].parse()?;
                match pick_elements[1] {
                    "red" => {
                        if c > min_red {
                            min_red = c
                        }
                    }
                    "green" => {
                        if c > min_green {
                            min_green = c
                        }
                    }
                    "blue" => {
                        if c > min_blue {
                            min_blue = c
                        }
                    }
                    _ => return Err(anyhow!("Invalid color")),
                }
            }
        }

        sum += min_red * min_green * min_blue;
    }

    println!("sum 2: {}", sum);
    Ok(())
}
