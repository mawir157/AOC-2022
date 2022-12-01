use crate::advent_helper::advent_helper::parse_line_groups;

pub fn calorie_count(elf: &String, sep: &str) -> i64
{
	let mut total: i64 = 0;

	let cals = elf.split(sep);

	for cal in cals {
		total += cal.parse::<i64>().unwrap();
	}

	return total;
}

pub fn run()
{
	let elfs = parse_line_groups("../input/input01.txt", "|");
	let mut cals: Vec<_> = elfs.iter().map(|x| calorie_count(x, "|")).collect();
	cals.sort_by(|a, b| b.cmp(a));

	println!("Day 1");
	println!("  Part 1: {}", cals[0]);
	println!("  Part 2: {}", cals[0] + cals[1] + cals[2]);
	return;
}
