#include "AH.h"

namespace Day05
{

	typedef std::tuple<size_t, size_t, size_t> Move;

	void ParseInput(std::vector<std::string>& v, const std::vector<std::string>& ss)
	{
		for (size_t j = 0; j < 8; ++j) {
			for (size_t i = 1; i < ss[j].size(); i += 4) {
				auto c = ss[j].at(i);
				if (c != ' ')
					v[i / 4] += c;
			}
		}
	}

	std::vector<Move> ParseMoves(const std::vector<std::string>& ss)
	{
		std::vector<Move> vec;
		for (size_t i = 10; i < ss.size(); ++i)
		{
			const auto ps = AH::Split(ss[i], ' ');
			const Move tpl = 
				std::make_tuple(std::stoi(ps[1]),std::stoi(ps[3]),std::stoi(ps[5]));
			vec.push_back(tpl);
		}
		return vec;
	}

	void ApplyMove(std::vector<std::string>& v, const Move m, const bool b)
	{
		const size_t n = std::get<0>(m);
		const size_t f = std::get<1>(m) - 1;
		const size_t t = std::get<2>(m) - 1;

		std::string from_new = v[f].substr(n, v[f].length());
		std::string moveable  = v[f].substr(0, n);
		if (b)
			std::reverse(moveable.begin(), moveable.end());
		std::string to_new   = moveable + v[t];

		v[f] = from_new;
		v[t] = to_new;
	}

	int Run(const std::string& filename)
	{
		const auto lines = AH::ReadTextFile(filename);
		std::vector<std::string> bricks1 = { "", "", "", "", "", "", "", "", "" };
		std::vector<std::string> bricks2 = { "", "", "", "", "", "", "", "", "" };
		ParseInput(bricks1, lines);
		ParseInput(bricks2, lines);

		const auto moves = ParseMoves(lines);

		for (auto & mv : moves)
		{
			ApplyMove(bricks1, mv, true);
			ApplyMove(bricks2, mv, false);
		}
		std::string part1 = "";
		for (auto b : bricks1)
			part1 += b.at(0);

		std::string part2 = "";
		for (auto b : bricks2)
			part2 += b.at(0);
		
		AH::PrintSoln(5, part1, part2);

		return 0;
	}

}
