use crate::advent_helper::advent_helper::parse_line_groups;

pub fn calorie_count(elf: &String, sep: &str) -> i64
{
	return elf.split(sep).map(|x| x.parse::<i64>().unwrap()).sum::<i64>();
}

pub fn run()
{
	let elfs = parse_line_groups("../input/input01.txt", "|");
	let mut cals: Vec<_> = elfs.iter().map(|x| calorie_count(x, "|")).collect();
	cals.sort_by(|a, b| b.cmp(a));

	println!("Day 1");
	println!("  Part 1: {}", cals[0]);
	println!("  Part 2: {}", cals.iter().take(3).sum::<i64>());
	return;
}
