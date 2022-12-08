#include "AH.h"

namespace Day08
{

	typedef std::vector<std::vector<unsigned int>> Forest;

	Forest ParseInput(const std::vector<std::string>& ss)
	{
		Forest f(ss.size());
		for (size_t r = 0; r < ss.size(); ++r) {
			for (size_t c = 0; c < ss[r].size(); ++c) {
				auto h = ss[r].substr(c,1);
				f[r].push_back(std::stoi(h));
			}
		}

		return f;
	}

	bool isVisible(const Forest f, const size_t x, const size_t y)
	{
		const auto height = f[x][y];

		// l
		bool visible = true;
		for (int i = x - 1; i >= 0; --i)
		{
			if (f[i][y] >= height)
			{
				visible = false;
				break;
			}
		}
		if (visible)
			return true;

		// r
		visible =  true;
		for (int i = x + 1; i < int(f.size()); ++i)
		{
			if (f[i][y] >= height)
			{
				visible = false;
				break;
			}
		}
		if (visible)
			return true;

		// u
		visible =  true;
		for (int i = y - 1; i >= 0; --i)
		{
			if (f[x][i] >= height)
			{
				visible = false;
				break;
			}

		}
		if (visible)
			return true;

		// d
		visible =  true;
		for (int i = y + 1; i < int(f[x].size()); ++i)
		{
			if (f[x][i] >= height)
			{
				visible = false;
				break;
			}
		}

		return visible;
	}

	int CountVisible(const Forest f)
	{
		int count = 0;
		for (size_t r = 0; r < f.size(); ++r)
			for (size_t c = 0; c < f[r].size(); ++c)
				count += isVisible(f, r, c) ? 1 : 0;

		return count;
	}

	int Scenic(const Forest f, const size_t x, const size_t y)
	{
		const auto height = f[x][y];
		int score = 1;

		// l 
		int count = 0;
		for (int i = x - 1; i >= 0; --i)
		{
			count++;
			if (f[i][y] >= height)
				break;
		}
		score *= count;

		// r
		count = 0;
		for (int i = x + 1; i < int(f.size()); ++i)
		{
			count++;
			if (f[i][y] >= height)
				break;
		}
		score *= count;

		// u
		count = 0;
		for (int i = y - 1; i >= 0; --i)
		{
			count++;
			if (f[x][i] >= height)
				break;
		}
		score *= count;

		// d
		count = 0;
		for (int i = y + 1; i < int(f[x].size()); ++i)
		{
			count++;
			if (f[x][i] >= height)
				break;
		}
		score *= count;

		return score;
	}

	int MostScenic(const Forest f)
	{
		int most_scenic = 0;
		for (size_t r = 0; r < f.size(); ++r)
			for (size_t c = 0; c < f[r].size(); ++c)
			{
				const int scenic = Scenic(f, r, c);
				if (scenic > most_scenic)
					most_scenic = scenic;
			}

		return most_scenic;	
	}

	int Run(const std::string& filename)
	{
		const auto lines = AH::ReadTextFile(filename);
		const auto forest = ParseInput(lines);
		
		AH::PrintSoln(8, CountVisible(forest), MostScenic(forest));

		return 0;
	}

}
