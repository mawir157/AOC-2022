#include "AH.h"

namespace Day12
{

	typedef std::pair<int, int> Pos;
	typedef std::vector<std::vector<int>> Grid;

	Grid MakeGrid(const std::vector<std::string>& ss, Pos & start, Pos & end)
	{
		Grid g;
		for (size_t i = 0; i < ss.size(); ++i)
		{
			std::vector<int> row;
			for (size_t j = 0; j < ss[i].size(); ++j)
			{
				const auto r = ss[i].at(j);
				if (r == 'S')
				{
					row.push_back(0);
					start = std::make_pair(j, i);
				}
				else if (r == 'E')
				{
					row.push_back(26);
					end = std::make_pair(j, i);
				}
				else
				{
					row.push_back(int(r) - 96);
				}
			}
			g.push_back(row);
		}

		return g;
	}

	std::vector<Pos> GetNbrs(const Pos p, const Grid & g)
	{
		const auto target = g[p.second][p.first] + 1;
		std::vector<Pos> nbrs;

		if (p.first > 0)
			if(g[p.second][p.first - 1] <= target)
				nbrs.emplace_back(p.first - 1, p.second);

		if (p.first < (int(g[0].size()) - 1))
			if(g[p.second][p.first + 1] <= target)
				nbrs.emplace_back(p.first + 1, p.second);

		if (p.second > 0)
			if(g[p.second - 1][p.first] <= target)
				nbrs.emplace_back(p.first, p.second - 1);

		if (p.second < (int(g.size()) - 1))
			if(g[p.second + 1][p.first] <= target)
				nbrs.emplace_back(p.first, p.second + 1);

		return nbrs;
	}

	Pos minAlt(std::map<Pos, int> & flagged)
	{
		int min = -1;
		Pos pMin = std::make_pair(0,0);
		for (auto const& [key, val] : flagged)
		{
			if ((min == -1) || (val < min))
			{
				min = val;
				pMin = key;
			}
		}
		flagged.erase (pMin);

		return pMin; 
	}

	std::map<Pos, int> GraphTraverse(const Grid & g, const Pos source, const Pos target)
	{
		std::map<Pos, int> dist;
		std::map<Pos, int> marked;

		for (size_t y = 0; y < g.size(); ++y)
		{
			for (size_t x = 0; x < g[0].size(); ++x)
			{
				const Pos p = std::make_pair(x,y);
				dist[p] = 1000000;
			}
		}
		dist[source] = 0;
		marked[source] = 0;

		while (marked.size() > 0)
		{
			const Pos u = minAlt(marked);

			if (u == target) {
				return dist;
			}

			const int distU = dist[u];

			const auto moves = GetNbrs(u, g);
			for (auto m : moves)
			{
				const auto new_dist = distU + 1;
				if (new_dist < dist[m])
				{
					dist[m] = new_dist;
					marked[m] = new_dist;
				}
			}
		}

		return dist;
	}

	uint64_t Run(const std::string& filename)
	{
		const auto inputLines = AH::ReadTextFile(filename);
		Pos start, end;
		const auto g = MakeGrid(inputLines, start, end);

		const std::map<Pos, int> d = GraphTraverse(g, start, end);

		int part2 = 1000000;
		for (size_t i = 0; i < g.size(); ++i)
		{
			for (size_t j = 0; j < g[i].size(); ++j)
			{
				if (g[i][j] == 1)
				{
					const auto p = std::make_pair(j, i);
					const auto d2 = GraphTraverse(g, p, end);
					if (d2.at(end) < part2)
						part2 = d2.at(end);
				}
			}
		}

		AH::PrintSoln(12, d.at(end), part2);

		return 0;
	}

}
