use std::{
    cmp::{max, min},
    collections::HashSet,
};

use util::open_input;
fn main() {
    let f = "day11/day11.txt";
    println!("day11");
    calc_1(f);
    calc_2(f);
}

struct Point {
    x: usize,
    y: usize,
}

fn d(p1: &Point, p2: &Point, er: &HashSet<usize>, ec: &HashSet<usize>) -> usize {
    let x_min = min(p1.x, p2.x);
    let x_max = max(p1.x, p2.x);
    let y_min = min(p1.y, p2.y);
    let y_max = max(p1.y, p2.y);

    x_max - x_min + y_max - y_min
        + er.iter().filter(|x| **x > x_min && **x < x_max).count()
        + ec.iter().filter(|y| **y > y_min && **y < y_max).count()
}

fn d2(p1: &Point, p2: &Point, er: &HashSet<usize>, ec: &HashSet<usize>) -> usize {
    let x_min = min(p1.x, p2.x);
    let x_max = max(p1.x, p2.x);
    let y_min = min(p1.y, p2.y);
    let y_max = max(p1.y, p2.y);

    x_max - x_min + y_max - y_min
        + (er.iter().filter(|x| **x > x_min && **x < x_max).count()) * 999999
        + (ec.iter().filter(|y| **y > y_min && **y < y_max).count()) * 999999
}

pub fn calc_1(f: &'static str) {
    let mut sum = 0;
    let mut points: Vec<Point> = vec![];

    let mut empty_rows: HashSet<usize> = HashSet::new();
    let mut empty_columns: HashSet<usize> = HashSet::new();

    for (x, line) in open_input(f).unwrap().enumerate() {
        empty_rows.insert(x);
        for (y, c) in line.unwrap().char_indices() {
            empty_columns.insert(y);
            if c == '#' {
                points.push(Point { x, y });
            }
        }
    }

    for p in &points {
        empty_rows.remove(&p.x);
        empty_columns.remove(&p.y);
    }

    let p_len = points.len();
    for i in 0..(p_len - 1) {
        for j in i + 1..p_len {
            let d = d(&points[i], &points[j], &empty_rows, &empty_columns);
            sum += d
        }
    }

    println!("sum 1: {}", sum);
}

pub fn calc_2(f: &'static str) {
    let mut sum = 0;
    let mut points: Vec<Point> = vec![];

    let mut empty_rows: HashSet<usize> = HashSet::new();
    let mut empty_columns: HashSet<usize> = HashSet::new();

    for (x, line) in open_input(f).unwrap().enumerate() {
        empty_rows.insert(x);
        for (y, c) in line.unwrap().char_indices() {
            empty_columns.insert(y);
            if c == '#' {
                points.push(Point { x, y });
            }
        }
    }

    for p in &points {
        empty_rows.remove(&p.x);
        empty_columns.remove(&p.y);
    }

    let p_len = points.len();
    for i in 0..(p_len - 1) {
        for j in i + 1..p_len {
            let d = d2(&points[i], &points[j], &empty_rows, &empty_columns);
            sum += d
        }
    }

    println!("sum 1: {}", sum);
}
