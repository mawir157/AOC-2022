#include "AH.h"

namespace Day02
{
	
	std::pair<int, int> score(const std::string & s)
	{
		const auto players = AH::Split(s, ' ');
		const int player1 = int(players[0][0]) - 65;
		const int player2 = int(players[1][0]) - 88;

		const auto result1 = (player2 - player1 + 4) % 3;
		const int final1 = 3 * result1;

		const int final2 = 3 * player2;

		int play2 = (player1 + 3 + (player2 - 1)) % 3;

		return std::make_pair<int, int>(1 + player2 + final1, 1 + play2 + final2);
	}

	int Run(const std::string& filename)
	{
		const auto rounds = AH::ReadTextFile(filename);
		int part1 = 0;
		int part2 = 0;

		for (auto & r : rounds)
		{
			const auto p = score(r);
			part1 += p.first;
			part2 += p.second;
		}
		
		AH::PrintSoln(2, part1, part2);

		return 0;
	}

}
