#include "AH.h"

namespace Day10
{

	typedef std::pair<std::string, int> Instruction;

	Instruction ParseInput(const std::string s)
	{
		const auto ps = AH::Split(s, ' ');

		if (ps[0] == "addx")
			return make_pair(ps[0], std::stoi(ps[1]));
		else
			return make_pair(ps[0], 0);
	}

	std::vector<int> RunInstructions(const std::vector<Instruction> & is)
	{
		int reg = 1;
		std::vector<int> values{reg};
		for (auto & ins : is)
		{
			values.push_back(reg);
			if (ins.first == "addx")
			{
				reg += ins.second;
				values.push_back(reg);
			}
		}

		return values;
	}

	int Part1(const std::vector<int> & rs, const std::vector<size_t> & is)
	{
		int total = 0;
		for (auto & i : is)
			total += (int(i) * rs[i - 1]);

		return total;
	}

	std::string RenderCRT(const std::vector<int> & rs)
	{
		std::string crt = "";
		int pixel = 0;
		for (size_t i = 0; i < rs.size(); ++i, ++pixel)
		{
			const auto p = pixel % 40;
			const auto r = rs[i];
			if (((r - 1) <= p) && (p <= (r + 1)))
				crt += "#";
			else
				crt += " ";
		}

		return crt;
	}

	int Run(const std::string& filename)
	{
		auto lines = AH::ReadTextFile(filename);
		std::vector<Instruction> is;
		for (auto l : lines)
			is.push_back(ParseInput(l));

		const auto regs = RunInstructions(is);
		const std::vector<size_t> indices{20, 60, 100, 140, 180, 220};
		const auto crt = RenderCRT(regs);

		AH::PrintSoln(10, Part1(regs, indices), 0);

		for (int i = 0; i < 6; ++i)
			std::cout << crt.substr(i*40, 40) << std::endl;

		return 0;
	}

}
