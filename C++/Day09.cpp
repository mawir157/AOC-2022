#include "AH.h"

namespace Day09
{

	typedef std::pair<int, int> Pos;
	typedef std::pair<std::string, int> Move;

	Move ParseInput(const std::string s)
	{
		const auto ps = AH::Split(s, ' ');

		return make_pair(ps[0], std::stoi(ps[1]));
	}

	void update(Pos & t, const Pos h)
	{
		const auto dx = h.first - t.first;
		const auto dy = h.second - t.second;

		if ((std::abs(dx) > 1) || (std::abs(dy) > 1))
		{
			t.first += AH::sgn(dx);
			t.second += AH::sgn(dy);
		}

		return;
	}

	void PullChain(std::vector<Pos> & chain, std::map<Pos, int> & history, const Move m)
	{
		for (int i = 0; i < m.second; ++i)
		{
			if (m.first == "U")
				chain[0].second -= 1;
			else if (m.first == "D")
				chain[0].second += 1;
			else if (m.first == "L")
				chain[0].first -= 1;
			else if (m.first == "R")
				chain[0].first += 1;

			for (size_t j = 1; j < chain.size(); ++j)
				update(chain[j], chain[j - 1]);

			history[chain.back()] += 1;
		}
	}

	int Run(const std::string& filename)
	{
		auto lines = AH::ReadTextFile(filename);
		std::vector<Move> moves;
		for (auto l : lines)
			moves.push_back(ParseInput(l));

		std::map<Pos, int> hist1;
		std::map<Pos, int> hist2;

		std::vector<Pos> chain1;
		for (int i = 0; i < 2; ++i)
			chain1.emplace_back(0 ,0);

		std::vector<Pos> chain2;
		for (int i = 0; i < 10; ++i)
			chain2.emplace_back(0 ,0);

		for (auto m : moves)
		{
			PullChain(chain1, hist1, m);
			PullChain(chain2, hist2, m);
		}

		AH::PrintSoln(9, hist1.size(), hist2.size());

		return 0;
	}

}
