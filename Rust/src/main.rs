mod advent_helper;
mod day01;

fn main() {
	let pattern = std::env::args().nth(1).expect("no pattern given");

	match pattern.as_str()
	{
		"01" => day01::run(), 
		_  =>
		{
			day01::run();
		},
	}



}
