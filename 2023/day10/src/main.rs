use std::{collections::VecDeque, fmt::Display};

use util::open_input;

#[derive(Debug, PartialEq)]
struct Tile {
    x: usize,
    y: usize,
    ttype: TileType,
}

#[derive(Debug, PartialEq)]
enum TileType {
    VerticalPipe,
    HorizontalPipe,
    NEBend,
    NWBend,
    SWBend,
    SEBend,
    Ground,
    StartingPosition,
}

impl Tile {
    fn from(c: char, x: usize, y: usize) -> Self {
        let ttype = match c {
            '|' => TileType::VerticalPipe,
            '-' => TileType::HorizontalPipe,
            'L' => TileType::NEBend,
            'J' => TileType::NWBend,
            '7' => TileType::SWBend,
            'F' => TileType::SEBend,
            '.' => TileType::Ground,
            'S' => TileType::StartingPosition,
            _ => panic!(),
        };
        Self { x, y, ttype }
    }

    fn get_n<'a>(&self, map: &'a [Vec<Tile>]) -> Option<&'a Tile> {
        if self.x > 0 {
            Some(&map[self.x - 1][self.y])
        } else {
            None
        }
    }

    fn get_s<'a>(&self, map: &'a [Vec<Tile>]) -> Option<&'a Tile> {
        if self.x < map.len() {
            Some(&map[self.x + 1][self.y])
        } else {
            None
        }
    }

    fn get_w<'a>(&self, map: &'a [Vec<Tile>]) -> Option<&'a Tile> {
        if self.y > 0 {
            Some(&map[self.x][self.y - 1])
        } else {
            None
        }
    }

    fn get_e<'a>(&self, map: &'a [Vec<Tile>]) -> Option<&'a Tile> {
        if self.y < map[0].len() {
            Some(&map[self.x][self.y + 1])
        } else {
            None
        }
    }

    fn get_connected<'a>(&self, map: &'a Vec<Vec<Tile>>) -> Vec<&'a Tile> {
        match self.ttype {
            TileType::VerticalPipe => vec![self.get_n(map), self.get_s(map)]
                .into_iter()
                .flatten()
                .collect(),
            TileType::HorizontalPipe => vec![self.get_e(map), self.get_w(map)]
                .into_iter()
                .flatten()
                .collect(),
            TileType::NEBend => vec![self.get_n(map), self.get_e(map)]
                .into_iter()
                .flatten()
                .collect(),
            TileType::NWBend => vec![self.get_n(map), self.get_w(map)]
                .into_iter()
                .flatten()
                .collect(),
            TileType::SWBend => vec![self.get_s(map), self.get_w(map)]
                .into_iter()
                .flatten()
                .collect(),
            TileType::SEBend => vec![self.get_s(map), self.get_e(map)]
                .into_iter()
                .flatten()
                .collect(),
            TileType::Ground => vec![],
            TileType::StartingPosition => [
                self.get_n(map),
                self.get_s(map),
                self.get_e(map),
                self.get_w(map),
            ]
            .iter()
            .flatten()
            .filter(|t| t.get_connected(map).contains(&self))
            .copied()
            .collect(),
        }
    }
}

impl Display for Tile {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "{}",
            match self.ttype {
                TileType::VerticalPipe => '|',
                TileType::HorizontalPipe => '-',
                TileType::NEBend => 'L',
                TileType::NWBend => 'J',
                TileType::SWBend => '7',
                TileType::SEBend => 'F',
                TileType::Ground => '.',
                TileType::StartingPosition => 'S',
            }
        )
    }
}

fn main() {
    let f = "day10/day10.txt";
    println!("day10");
    calc_1(f);
    calc_2(f);
}

pub fn calc_1(f: &'static str) {
    let mut map = vec![];

    let mut start_x = None;
    let mut start_y = None;

    for (x, line) in open_input(f).unwrap().enumerate() {
        map.push(vec![]);
        for (y, c) in line.unwrap().char_indices() {
            let tile = Tile::from(c, x, y);
            if let TileType::StartingPosition = tile.ttype {
                (start_x, start_y) = (Some(x), Some(y));
            }
            map[x].push(tile)
        }
    }

    let mut distance_map = vec![vec![None; map[0].len()]; map.len()];
    let mut queue = VecDeque::from([(&map[start_x.unwrap()][start_y.unwrap()], 0)]);

    while !queue.is_empty() {
        let (tile, current_d) = queue.pop_front().unwrap();
        if let Some(previous_d) = distance_map[tile.x][tile.y] {
            if current_d < previous_d {
                distance_map[tile.x][tile.y] = Some(current_d);
            } else {
                continue;
            }
        } else {
            distance_map[tile.x][tile.y] = Some(current_d);
        }

        for t in tile.get_connected(&map) {
            queue.push_back((t, current_d + 1));
        }
    }

    let sum = *distance_map
        .iter()
        .flat_map(|r| r.iter())
        .flatten()
        .max()
        .unwrap();
    println!("sum 1: {}", sum);
}

pub fn calc_2(f: &'static str) {
    let mut sum = 0;
    for _line in open_input(f).unwrap() {
        sum += 1;
    }
    println!("sum 2: {}", sum);
}
