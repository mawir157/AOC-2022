#include "AH.h"

namespace Day03
{
	
	std::string string_intersection(std::string s1, std::string s2) {
		std::sort(s1.begin(), s1.end());
		std::sort(s2.begin(), s2.end());

		std::string intersection;
		std::set_intersection(s1.begin(), s1.end(),
		                      s2.begin(), s2.end(),
		                      std::back_inserter(intersection));

		return intersection;
	}

	int score(const std::vector<std::string> bags)
	{
		int total = 0;

		for (auto & bag : bags)
		{
			const int half = bag.size() / 2;
			std::string lhs = bag.substr(0, half);
			std::string rhs = bag.substr(half, half);
			std::string i = string_intersection(lhs, rhs);

			int score = int(i.at(0));
			score -= (score <= 90) ? 38 : 96;

			total += score;
		}

		return total;
	}

	int assignBadges(const std::vector<std::string> ss)
	{
		int total = 0;
		for (size_t i = 0; i < ss.size(); i += 3)
		{
			std::string i1 = string_intersection(ss[i], ss[i+1]);
			std::string i2 = string_intersection(ss[i+2], i1);

			auto score = int(i2.at(0));
			score -= (score <= 90) ? 38 : 96;
			total += score;
		}

		return total;
	}

	int Run(const std::string& filename)
	{
		const auto rounds = AH::ReadTextFile(filename);
		
		AH::PrintSoln(3, score(rounds), assignBadges(rounds));

		return 0;
	}

}
