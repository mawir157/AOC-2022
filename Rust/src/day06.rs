use crate::advent_helper::advent_helper::read_strs;
use crate::advent_helper::advent_helper::print_soln;
use itertools::Itertools;

fn get_marker (s: &String, n: usize) -> usize
{
	let t = s.as_bytes().windows(n).position(|v| v.to_vec().iter().unique().collect::<Vec<&u8>>().len() == n).unwrap();
	return t + n
}

pub fn run()
{
	let input = read_strs("../input/input06.txt").unwrap();

	print_soln(6, get_marker(&input[0], 4), get_marker(&input[0], 14));
	return;
}
