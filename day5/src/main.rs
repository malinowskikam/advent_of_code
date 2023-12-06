use std::cmp::{max, min};

use anyhow::{anyhow, Result};
use util::open_input;

fn main() -> Result<()> {
    let f = "day5/day5.txt";
    println!("day5");
    calc_1(f)?;
    calc_2(f)?;
    Ok(())
}

pub fn calc_1(f: &'static str) -> Result<()> {
    let mut lines_iter = open_input(f)?.peekable();
    let header_line = lines_iter.next().unwrap()?;
    let mut seed_numbers: Vec<i64> = header_line.split(":").collect::<Vec<&str>>()[1]
        .trim()
        .split(" ")
        .map(|s| s.parse::<i64>().unwrap())
        .collect();

    //skip empty line
    lines_iter.next();

    while lines_iter.peek().is_some() {
        let mut map: Vec<(i64, i64, i64)> = vec![];

        //skip header
        lines_iter.next().unwrap()?;

        for line in &mut lines_iter {
            let line = line?;
            if line == "" {
                break;
            }
            let map_entry: Vec<i64> = line
                .trim()
                .split(" ")
                .map(|s| s.parse::<i64>().unwrap())
                .collect();
            assert!(map_entry.len() == 3);
            map.push((map_entry[0], map_entry[1], map_entry[2]));
        }

        seed_numbers.iter_mut().for_each(|n| {
            for map_entry in &map {
                if *n >= map_entry.1 && *n < map_entry.1 + map_entry.2 {
                    *n = map_entry.0 + *n - map_entry.1;
                    break;
                }
            }
        });
    }

    println!("sum 1: {}", seed_numbers.iter().min().unwrap());
    Ok(())
}

#[derive(Debug, Clone, Copy)]
struct Range {
    start: i64,
    end: i64,
}

impl Range {
    fn intersect(&self, mapping: &MappingRange) -> Range {
        Range::new(max(self.start, mapping.start), min(self.end, mapping.end))
    }

    fn count(&self) -> i64 {
        max(self.end - self.start + 1, 0)
    }

    fn new(start: i64, end: i64) -> Self {
        Range { start, end }
    }

    fn from_count(start: i64, count: i64) -> Self {
        Range {
            start: start,
            end: start + count - 1,
        }
    }

    fn apply_mapping(&self, mapping: &MappingRange) -> Range {
        Range {
            start: self.start + mapping.map,
            end: self.end + mapping.map
        }
    }
}

fn into_ranges(iter: impl Iterator<Item = i64>) -> Result<Vec<Range>> {
    let mut res = vec![];
    let mut v_iter = iter.peekable();
    while v_iter.peek().is_some() {
        res.push(Range::from_count(
            v_iter.next().ok_or(anyhow!("Missing start"))?,
            v_iter.next().ok_or(anyhow!("Missing count"))?,
        ));
    }
    Ok(res)
}

fn map_ranges(ranges: Vec<Range>, mappings: &Vec<MappingRange>) -> Vec<Range> {
    ranges.into_iter().fold(Vec::new(), |mut x, y| {
        x.append(&mut map_range(y, mappings));
        x
    })
}

fn map_range(range: Range, mappings: &Vec<MappingRange>) -> Vec<Range> {
    let mut new_ranges = Vec::new();

    let mut mapped = false;
    for mapping in mappings {
        let i_range = range.intersect(mapping);
        
        if i_range.count() > 0 {
            mapped = true;
            new_ranges.push(i_range.apply_mapping(mapping));
            
            
            let before_range = Range::new(range.start, i_range.start - 1);
            
            if before_range.count() > 0 {
                new_ranges.append(&mut map_range(before_range, mappings));
            }

            let after_range = Range::new(i_range.end + 1, range.end);
            
            if after_range.count() > 0 {
                new_ranges.append(&mut map_range(after_range, mappings))
            }
            break;
        }
    };

    if !mapped {
        new_ranges.push(range);
    }

    new_ranges
}

struct MappingRange {
    start: i64,
    end: i64,
    map: i64,
}

impl MappingRange {
    fn from_count(mapped_start: i64, orig_start: i64, count: i64) -> Self {
        MappingRange {
            start: orig_start,
            end: orig_start + count - 1,
            map: mapped_start - orig_start,
        }
    }
}

pub fn calc_2(f: &'static str) -> Result<()> {
    let mut lines_iter = open_input(f)?.peekable();
    let header_line = lines_iter.next().unwrap()?;
    let mut seed_ranges = into_ranges(
        header_line.split(":").collect::<Vec<&str>>()[1]
            .trim()
            .split(" ")
            .map(|s| s.parse::<i64>().unwrap()),
    )?;
    
    // skip empty line
    lines_iter.next();
    while lines_iter.peek().is_some() {
        let mut mapping_ranges: Vec<MappingRange> = vec![];

        //skip header
        lines_iter.next().unwrap()?;

        for line in &mut lines_iter {
            let line = line?;
            if line == "" {
                break;
            }
            let map_entry: Vec<i64> = line
                .trim()
                .split(" ")
                .map(|s| s.parse::<i64>().unwrap())
                .collect();
            assert!(map_entry.len() == 3);
            mapping_ranges.push(MappingRange::from_count(
                map_entry[0],
                map_entry[1],
                map_entry[2],
            ));
        }

        seed_ranges = map_ranges(seed_ranges, &mapping_ranges);
    }

    println!(
        "sum 2: {}",
        seed_ranges.iter().map(|r| r.start).min().unwrap()
    );
    Ok(())
}
