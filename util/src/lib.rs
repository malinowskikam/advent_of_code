use anyhow::Result;
use std::fs::File;
use std::io::{prelude::*, BufReader, Lines};

pub fn open_input(filename: &'static str) -> Result<Lines<impl BufRead>> {
    let file = File::open(filename)?;
    let reader = BufReader::new(file);
    Ok(reader.lines())
}
