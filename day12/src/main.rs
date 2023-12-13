use std::collections::HashMap;

use util::open_input;

fn main() {
    let f = "day12/day12.txt";
    println!("day12");
    calc_1(f);
    calc_2(f);
}

pub fn count_hash(
    source: &[u8],
    index: usize,
    groups: &Vec<usize>,
    group_index: usize,
    current_group_len: usize,
    cache: &mut HashMap<(usize, usize, usize), usize>,
) -> usize {
    if group_index >= groups.len() {
        0
    } else {
        if current_group_len < groups[group_index] {
            count_possibilities(
                source,
                index + 1,
                groups,
                group_index,
                current_group_len + 1,
                cache,
            )
        } else {
            0
        }
    }
}

pub fn count_dot(
    source: &[u8],
    index: usize,
    groups: &Vec<usize>,
    group_index: usize,
    current_group_len: usize,
    cache: &mut HashMap<(usize, usize, usize), usize>,
) -> usize {
    if current_group_len == 0 {
        count_possibilities(
            source,
            index + 1,
            groups,
            group_index,
            current_group_len,
            cache,
        )
    } else {
        if groups[group_index] == current_group_len {
            count_possibilities(source, index + 1, groups, group_index + 1, 0, cache)
        } else {
            0
        }
    }
}

pub fn count_possibilities(
    source: &[u8],
    index: usize,
    groups: &Vec<usize>,
    group_index: usize,
    current_group_len: usize,
    cache: &mut HashMap<(usize, usize, usize), usize>,
) -> usize {
    let key = (index, group_index, current_group_len);
    if cache.contains_key(&key) {
        return *cache.get(&key).unwrap();
    }

    if group_index >= groups.len() && current_group_len > 0 {
        cache.insert(key, 0);
        return 0;
    }
    if group_index < groups.len() && groups[group_index] < current_group_len {
        cache.insert(key, 0);
        return 0;
    }

    if index >= source.len() {
        if group_index == groups.len() && current_group_len == 0
            || group_index + 1 == groups.len() && current_group_len == groups[group_index]
        {
            cache.insert(key, 1);
            return 1;
        } else {
            cache.insert(key, 0);
            return 0;
        }
    }

    let ret_value = match source[index] as char {
        '#' => count_hash(source, index, groups, group_index, current_group_len, cache),
        '.' => count_dot(source, index, groups, group_index, current_group_len, cache),
        '?' => {
            count_hash(source, index, groups, group_index, current_group_len, cache)
                + count_dot(source, index, groups, group_index, current_group_len, cache)
        }
        _ => panic!(),
    };
    cache.insert(key, ret_value);
    ret_value
}

pub fn calc_1(f: &'static str) {
    let mut sum = 0;

    for line in open_input(f).unwrap() {
        let line = line.unwrap();

        let parts = line.split(" ").collect::<Vec<_>>();
        let arrangement = parts[0];
        let groups = parts[1]
            .split(",")
            .map(|x| x.parse::<usize>().unwrap())
            .collect::<Vec<_>>();

        let mut cache: HashMap<(usize, usize, usize), usize> = HashMap::new();
        sum += count_possibilities(arrangement.as_bytes(), 0, &groups, 0, 0, &mut cache);
    }

    println!("sum 1: {}", sum);
}

pub fn calc_2(f: &'static str) {
    let mut sum = 0;

    for (i, line) in open_input(f).unwrap().enumerate() {
        let line = line.unwrap();

        let parts = line.split(" ").collect::<Vec<_>>();
        let arrangement = format!(
            "{}?{}?{}?{}?{}",
            parts[0], parts[0], parts[0], parts[0], parts[0]
        );

        let groups_line = parts[1]
            .split(",")
            .map(|x| x.parse::<usize>().unwrap())
            .collect::<Vec<_>>();
        let mut groups = vec![];
        for _ in 0..5 {
            groups.append(&mut groups_line.clone());
        }

        let mut cache: HashMap<(usize, usize, usize), usize> = HashMap::new();
        sum += count_possibilities(arrangement.as_bytes(), 0, &groups, 0, 0, &mut cache);
    }

    println!("sum 2 {}", sum);
}
