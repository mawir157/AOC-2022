use crate::advent_helper::advent_helper::read_strs;
use crate::advent_helper::advent_helper::print_soln;

fn score_char(c: char) -> i64
{
	let mut val:i64 = c as i64;

	if val <= 90 {
		val -= 38;
	} else {
		val -= 96;
	}

	return val;
}

fn score(s: &String) -> i64
{
	let half = s.chars().count() / 2;
	let lhs = &s[..half];
	let rhs = &s[half..];
	let i:String = lhs.chars().filter(|c| rhs.contains(&c.to_string())).collect();

	let c = i.chars().next().unwrap();

	return score_char(c);
}

fn assign_badge(ss: &[String]) -> i64
{
	match &ss[..] {
		[e1, e2, e3] => {
			let i1:String = e1.chars().filter(|c| e2.contains(&c.to_string())).collect();
			let i2:String = i1.chars().filter(|c| e3.contains(&c.to_string())).collect();
			return score_char(i2.chars().next().unwrap());
		},
		_ => { return -1; },
		};
}

pub fn run()
{
	let bags = read_strs("../input/input03.txt").unwrap();
	let scores: Vec<i64> = bags.iter().map(|x| score(x)).collect();
	let part2: Vec<i64> = bags.chunks(3).map(|x| assign_badge(x)).collect();

	print_soln(3, scores.iter().sum::<i64>(), part2.iter().sum::<i64>());
	return;
}
