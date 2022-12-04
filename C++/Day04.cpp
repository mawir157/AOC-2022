#include "AH.h"

namespace Day04
{
	
	std::pair<bool, bool> check (const std::string s)
	{
		const auto bounds = AH::Split(s, ',');
		const auto lhs = AH::Split(bounds[0], '-');
		const auto rhs = AH::Split(bounds[1], '-');

		const auto l1 = stoi(lhs[0]);
		const auto l2 = stoi(lhs[1]);
		const auto r1 = stoi(rhs[0]);
		const auto r2 = stoi(rhs[1]);

		const auto part1 = ((l1 <= r1) && (r2 <= l2)) || ((r1 <= l1) && (l2 <= r2));
		const auto part2 = ((l2 >= r1) && (r2 >= l1));

		return std::make_pair (part1, part2);
	}

	int Run(const std::string& filename)
	{
		const auto space = AH::ReadTextFile(filename);

		int part1 = 0;
		int part2 = 0;

		for (auto & s : space)
		{
			const auto p = check(s);
			part1 += p.first ? 1 : 0;
			part2 += p.second ? 1 : 0;
		}
		
		AH::PrintSoln(4, part1, part2);

		return 0;
	}

}
