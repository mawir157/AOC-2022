#include "AH.h"

namespace Day01
{

	int calorieCount(const std::string cals, const char sep) {
		int total = 0;
		const auto cs = AH::Split(cals, sep);
		for (auto & c : cs) {
			total += std::stoi(c);
		}

		return total;
	}

	int Run(const std::string& filename)
	{
		const auto inputLines = AH::ReadTextFile(filename);
		const auto elfs = AH::ParseLineGroups(inputLines, '|');

		std::vector<int> cals;
		cals.reserve(elfs.size());

		for (auto e : elfs) {
			cals.push_back(calorieCount(e, '|'));
		}

		std::sort(cals.begin(), cals.end(), std::greater{});
		
		AH::PrintSoln(1, cals[0], cals[0] + cals[1] + cals[2]);

		return 0;
	}

}
