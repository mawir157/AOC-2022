use crate::advent_helper::advent_helper::read_strs;
use crate::advent_helper::advent_helper::print_soln;

pub fn score(round: &String) -> (i64, i64)
{
	let players: Vec<String> = round.split(" ").map(|s| s.to_string()).collect();
	let player1:i64 = (players[0].chars().next().unwrap() as i64) - 65;
	let player2:i64 = (players[1].chars().next().unwrap() as i64) - 88;

	let result1 = (player2 - player1 + 4) % 3;
	let play2 = (player1 + player2 + 2) % 3;

	return (1 + player2 + (3 * result1), 1 + play2 + (3 * player2));
}

pub fn run()
{
	let rounds = read_strs("../input/input02.txt").unwrap();
	let scores: Vec<(i64, i64)> = rounds.iter().map(|x| score(x)).collect();
	let result: (i64, i64) = scores.iter().fold((0,0), |mut base, &x| {base.0 += x.0; base.1 += x.1; base});

	print_soln(2, result.0, result.1);
	return;
}
