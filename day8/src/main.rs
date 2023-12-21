use std::{cell::RefCell, collections::HashMap, rc::Rc};

use util::open_input;

fn main() {
    let f = "day8/day8.txt";
    println!("day8");
    //calc_1(f);
    calc_2(f);
}

#[derive(Debug)]
struct _Node {
    // key: String,
    left: Option<Node>,
    right: Option<Node>,
    is_finish: bool,
}

type Node = Rc<RefCell<_Node>>;

fn create_node(key: &str, left: Option<Node>, right: Option<Node>) -> Node {
    Node::new(RefCell::new(_Node {
        // key: key.into(),
        left,
        right,
        is_finish: key.ends_with('Z'),
    }))
}

pub fn calc_1(f: &'static str) {
    let mut sum = 0;

    let mut lines_iter = open_input(f).unwrap();
    let header = lines_iter.next().unwrap().unwrap();
    let header_bytes = header.as_bytes();
    let header_len = header.len();
    let mut nodes_map: HashMap<String, Node> = HashMap::new();

    for line in lines_iter.skip(1) {
        let line = line.unwrap();
        let line_split = line.split(" = ").collect::<Vec<&str>>();
        let key = line_split[0];
        let edges = line_split[1][1..line_split[1].len() - 1]
            .split(", ")
            .collect::<Vec<&str>>();

        for edge in &edges {
            if !nodes_map.contains_key(*edge) {
                let zzz_node = create_node(edge, None, None);
                nodes_map.insert((*edge).into(), zzz_node.clone());
            }
        }

        let node = if nodes_map.contains_key(key) {
            nodes_map.get(key).unwrap().clone()
        } else {
            let new_node: Rc<RefCell<_Node>> = create_node(key, None, None);
            nodes_map.insert(key.into(), new_node.clone());
            new_node
        };

        node.borrow_mut().left = Some(nodes_map.get(edges[0]).expect("no starting node").clone());
        node.borrow_mut().right = Some(nodes_map.get(edges[1]).expect("no starting node").clone());
    }

    let mut current_node = nodes_map.get("AAA").expect("no starting node").clone();

    while !current_node.borrow().is_finish {
        let new_current_node = match header_bytes[sum % header_len] as char {
            'R' => current_node.borrow().right.clone().unwrap(),
            'L' => current_node.borrow().left.clone().unwrap(),
            _ => panic!("asda"),
        };

        current_node = new_current_node;
        sum += 1;
    }
    println!("sum 1: {}", sum);
}

pub fn calc_2(f: &'static str) {
    let mut sum = 0;

    let mut lines_iter = open_input(f).unwrap();
    let header = lines_iter.next().unwrap().unwrap();
    let header_bytes = header.as_bytes();
    let header_len = header.len();
    let mut nodes_map: HashMap<String, Node> = HashMap::new();
    let mut starting_nodes = vec![];

    for line in lines_iter.skip(1) {
        let line = line.unwrap();
        let line_split = line.split(" = ").collect::<Vec<&str>>();
        let key = line_split[0];
        let edges = line_split[1][1..line_split[1].len() - 1]
            .split(", ")
            .collect::<Vec<&str>>();

        for edge in &edges {
            if !nodes_map.contains_key(*edge) {
                let zzz_node = create_node(edge, None, None);
                nodes_map.insert((*edge).into(), zzz_node.clone());
            }
        }

        let node = if nodes_map.contains_key(key) {
            nodes_map.get(key).unwrap().clone()
        } else {
            let new_node: Rc<RefCell<_Node>> = create_node(key, None, None);
            nodes_map.insert(key.into(), new_node.clone());
            new_node
        };

        node.borrow_mut().left = Some(nodes_map.get(edges[0]).unwrap().clone());
        node.borrow_mut().right = Some(nodes_map.get(edges[1]).unwrap().clone());

        if key.ends_with('A') {
            starting_nodes.push(key.to_owned())
        };
    }

    let mut current_nodes: Vec<Node> = starting_nodes
        .iter()
        .map(|key| nodes_map.get(key).unwrap().clone())
        .collect();
    let current_nodes_len = current_nodes.len();
    let mut current_finish_nodes = 0;
    let mut loop_lens = vec![0; current_nodes_len];
    while current_nodes_len > current_finish_nodes {
        let dir = header_bytes[sum % header_len] as char;
        for i in 0..current_nodes_len {
            if !current_nodes[i].borrow().is_finish {
                let current_node = &current_nodes[i];
                let new_node = match dir {
                    'R' => current_node.borrow().right.clone().unwrap(),
                    'L' => current_node.borrow().left.clone().unwrap(),
                    _ => panic!("Wrong directions"),
                };
                if new_node.borrow().is_finish {
                    loop_lens[i] = sum + 1;
                    current_finish_nodes += 1;
                };
                current_nodes[i] = new_node;
            }
        }
        sum += 1;
    }
    println!("loops: {:?}", loop_lens);
    println!("sum 2: {}", sum);
}
