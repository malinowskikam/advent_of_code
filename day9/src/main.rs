use util::open_input;

fn main() {
    let f = "day9/day9.txt";
    println!("day9");
    calc_1(f);
    calc_2(f);
}

pub fn calc_1(f: &'static str) {
    let mut sum = 0;
    for line in open_input(f).unwrap() {
        let line = line.unwrap();

        let mut line_report: Vec<Vec<i32>> = vec![];
        line_report.push(line.split(" ").map(|v| v.parse().unwrap()).collect());

        //generate derivatives
        while !line_report.last().unwrap().iter().all(|x| *x == 0) {
            let mut next_line = vec![];
            let last_line = line_report.last().unwrap();
            for i in 1..last_line.len() {
                next_line.push(last_line[i] - last_line[i - 1])
            }

            line_report.push(next_line);
        }

        //extrapolate
        let mut extrapolated_value = 0;
        for i in (0..=(line_report.len()-2)).rev() {
            extrapolated_value = line_report[i].last().unwrap() + extrapolated_value
        }
             
        sum += extrapolated_value;
    }
    println!("sum 1: {}", sum);
}

pub fn calc_2(f: &'static str) {
    let mut sum = 0;
    for line in open_input(f).unwrap() {
        let line = line.unwrap();

        let mut line_report: Vec<Vec<i32>> = vec![];
        line_report.push(line.split(" ").map(|v| v.parse().unwrap()).collect());

        //generate derivatives
        while !line_report.last().unwrap().iter().all(|x| *x == 0) {
            let mut next_line = vec![];
            let last_line = line_report.last().unwrap();
            for i in 1..last_line.len() {
                next_line.push(last_line[i] - last_line[i - 1])
            }

            line_report.push(next_line);
        }

        //extrapolate
        let mut extrapolated_value = 0;
        for i in (0..=(line_report.len()-2)).rev() {
            extrapolated_value = line_report[i].first().unwrap() - extrapolated_value
        }
             
        sum += extrapolated_value;
    }
    println!("sum 2: {}", sum);
}