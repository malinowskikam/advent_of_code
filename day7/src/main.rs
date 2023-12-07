use std::{cmp::Ordering, collections::HashMap};

use anyhow::{anyhow, Result};
use util::open_input;

fn main() -> Result<()> {
    let f = "day7/day7.txt";
    println!("day7");
    calc_1(f)?;
    calc_2(f)?;
    Ok(())
}

#[derive(Debug)]
struct Hand {
    cards: String,
    bet: i64,
    cards_map: HashMap<char, i32>,
}

fn get_card_value(card: &char) -> i32 {
    match *card {
        card if card.is_ascii_digit() => card as i32 - '0' as i32,
        'T' => 10,
        'J' => 11,
        'Q' => 12,
        'K' => 13,
        'A' => 14,
        _ => panic!("wtf"),
    }
}

fn compare_figures(hand1: &Hand, hand2: &Hand) -> Ordering {
    let hand1_max = hand1.cards_map.values().max().unwrap();
    let hand2_max = hand2.cards_map.values().max().unwrap();

    match hand1_max.cmp(hand2_max) {
        Ordering::Equal => {
            let hand1_count = &hand1.cards_map.values().count();
            let hand2_count = &hand2.cards_map.values().count();

            match hand1_count.cmp(hand2_count) {
                Ordering::Less => Ordering::Greater,
                Ordering::Equal => Ordering::Equal,
                Ordering::Greater => Ordering::Less,
            }
        }
        ord => ord,
    }
}

fn compare_cards(hand1: &Hand, hand2: &Hand) -> Ordering {
    hand1
        .cards
        .chars()
        .zip(hand2.cards.chars())
        .filter(|(c1, c2)| c1 != c2)
        .map(|(c1, c2)| get_card_value(&c1).cmp(&get_card_value(&c2)))
        .next()
        .expect("figures failed")
}

fn compare_hands(hand1: &Hand, hand2: &Hand) -> Ordering {
    match compare_figures(hand1, hand2) {
        Ordering::Equal => compare_cards(hand1, hand2),
        ord => ord,
    }
}

impl Hand {
    fn from_vec(v: Vec<&str>) -> Result<Self> {
        if v.len() != 2 {
            return Err(anyhow!("Wrong Vec"));
        }
        let bet = v[1].parse()?;

        let mut cards_map = HashMap::new();
        v[0].chars().for_each(|c| {
            *cards_map.entry(c).or_default() += 1;
        });

        Ok(Hand {
            cards: v[0].into(),
            bet,
            cards_map,
        })
    }
}

pub fn calc_1(f: &'static str) -> Result<()> {
    let sum: i64;
    let mut hands = Vec::new();
    for line in open_input(f)? {
        let line = line?;
        hands.push(Hand::from_vec(line.split(" ").collect())?)
    }

    hands.sort_by(compare_hands);

    sum = hands
        .iter()
        .enumerate()
        .map(|(i, hand)| (i as i64 + 1) * hand.bet)
        .sum();

    println!("sum 1: {}", sum);
    Ok(())
}

#[derive(Debug)]
struct Hand2 {
    cards: String,
    bet: i64,
    cards_map: HashMap<char, i32>,
    j_amount: i32,
}

fn get_card_value2(card: &char) -> i32 {
    match *card {
        card if card.is_ascii_digit() => card as i32 - '0' as i32,
        'T' => 10,
        'J' => 1,
        'Q' => 12,
        'K' => 13,
        'A' => 14,
        _ => panic!("wtf"),
    }
}

fn compare_figures2(hand1: &Hand2, hand2: &Hand2) -> Ordering {
    let hand1_max = hand1.cards_map.values().max().unwrap_or(&0) + hand1.j_amount;
    let hand2_max = hand2.cards_map.values().max().unwrap_or(&0) + hand2.j_amount;

    match hand1_max.cmp(&hand2_max) {
        Ordering::Equal => {
            let hand1_count = hand1.cards_map.values().count() as i32 + hand1.j_amount / 5;
            let hand2_count = hand2.cards_map.values().count() as i32 + hand2.j_amount / 5;

            match hand1_count.cmp(&hand2_count) {
                Ordering::Less => Ordering::Greater,
                Ordering::Equal => Ordering::Equal,
                Ordering::Greater => Ordering::Less,
            }
        }
        ord => ord,
    }
}

fn compare_cards2(hand1: &Hand2, hand2: &Hand2) -> Ordering {
    hand1
        .cards
        .chars()
        .zip(hand2.cards.chars())
        .filter(|(c1, c2)| c1 != c2)
        .map(|(c1, c2)| get_card_value2(&c1).cmp(&get_card_value2(&c2)))
        .next()
        .expect("figures failed")
}

fn compare_hands2(hand1: &Hand2, hand2: &Hand2) -> Ordering {
    match compare_figures2(hand1, hand2) {
        Ordering::Equal => compare_cards2(hand1, hand2),
        ord => ord,
    }
}

impl Hand2 {
    fn from_vec(v: Vec<&str>) -> Result<Self> {
        if v.len() != 2 {
            return Err(anyhow!("Wrong Vec"));
        }
        let bet = v[1].parse()?;

        let mut cards_map = HashMap::new();
        let mut j_amount = 0;
        v[0].chars().for_each(|c| {
            if c != 'J' {
                *cards_map.entry(c).or_default() += 1;
            } else {
                j_amount += 1;
            }
        });

        Ok(Hand2 {
            cards: v[0].into(),
            bet,
            cards_map,
            j_amount,
        })
    }
}

pub fn calc_2(f: &'static str) -> Result<()> {
    let sum: i64;

    let mut hands = Vec::new();
    for line in open_input(f)? {
        let line = line?;
        hands.push(Hand2::from_vec(line.split(" ").collect())?)
    }

    hands.sort_by(compare_hands2);
    
    sum = hands
        .iter()
        .enumerate()
        .map(|(i, hand)| (i as i64 + 1) * hand.bet)
        .sum();

    println!("sum 2: {}", sum);
    Ok(())
}
