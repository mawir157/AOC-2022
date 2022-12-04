use crate::advent_helper::advent_helper::read_strs;
use crate::advent_helper::advent_helper::print_soln;

fn to_range(s: &str) -> (i64, i64)
{
	let v: Vec<i64> = s.split("-").map(|x| x.parse::<i64>().unwrap()).collect();
	return (v[0], v[1]);
}

fn check(s: &String) -> (bool, bool)
{
	let bounds: Vec<(i64, i64)> = s.split(",").map(|x| to_range(x)).collect();

	let p1 = ((bounds[0].0 <= bounds[1].0) && (bounds[1].1 <= bounds[0].1)) ||
	         ((bounds[1].0 <= bounds[0].0) && (bounds[0].1 <= bounds[1].1));
	let p2 = (bounds[0].1 >= bounds[1].0) && (bounds[1].1 >= bounds[0].0);

	return (p1, p2);
}

pub fn run()
{
	let space = read_strs("../input/input04.txt").unwrap();
	let part1: Vec<&String> = space.iter().filter(|x| check(x).0).collect();
	let part2: Vec<&String> = space.iter().filter(|x| check(x).1).collect();

	print_soln(4, part1.len(), part2.len());
	return;
}
