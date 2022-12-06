use crate::advent_helper::advent_helper::read_strs;
use crate::advent_helper::advent_helper::print_soln;
use core::iter::zip;

fn parse_chunk(c: char) -> String
{
	if c != ' ' {
		return c.to_string();
	}
	return "".to_string();
}

fn parse_input(ss: Vec<String>) -> Vec<String>
{
	let mut base: Vec<String>  = vec!["","","","","","","","",""].iter().map(|x| x.to_string()).collect::<Vec<String>>();
	for s in ss {
		let tt = s.chars().skip(1).step_by(4);
		base = zip(base, tt).map(|(b, t)| b + &parse_chunk(t)).collect();
	}

	return base;
}

struct Move {
    n: usize,
    f: usize,
    t: usize,
}

fn parse_moves(ss: Vec<String>) -> Vec<Move>
{
	let ns = ss
		.iter()
		.map(|s| s.split(" ")
							.nth(1)
							.unwrap()
							.parse::<usize>()
							.unwrap())
		.collect::<Vec<usize>>();

	let fs = ss
		.iter()
		.map(|s| s.split(" ")
							.nth(3)
							.unwrap()
							.parse::<usize>()
							.unwrap())
		.collect::<Vec<usize>>();

	let ts = ss
		.iter()
		.map(|s| s.split(" ")
							.nth(5)
							.unwrap()
							.parse::<usize>()
							.unwrap())
		.collect::<Vec<usize>>();

  let r = zip(ns, zip(fs, ts))
  	.map(|(n, (f, t))| Move{n:n, f:(f-1), t:(t-1)})
  	.collect::<Vec<Move>>();

  return r
}

fn apply_move(mut blocks: Vec<String>, mv: &Move, b: bool) -> Vec<String> {

	let from_new = blocks[mv.f][mv.n .. ].to_string();
	let mut mvbl = blocks[mv.f][0 .. mv.n].to_string();

	if b {
		mvbl = mvbl.chars().rev().collect::<String>();
	}

	let to_new = mvbl + &blocks[mv.t];

	blocks[mv.f] = from_new;
	blocks[mv.t] = to_new;

	return blocks;
}

pub fn run()
{
	let input = read_strs("../input/input05.txt").unwrap();
	let mut p1 = parse_input(input[0 .. 8].to_vec());
	let mut p2 = parse_input(input[0 .. 8].to_vec());
	let moves = parse_moves(input[10 .. ].to_vec());

	for m in moves {
		p1 = apply_move(p1, &m, true);
		p2 = apply_move(p2, &m, false);
	}

	let part1 = p1.iter().map(|s| s.chars().next().unwrap()).collect::<String>();
	let part2 = p2.iter().map(|s| s.chars().next().unwrap()).collect::<String>();

	print_soln(5, part1, part2);
	return;
}
