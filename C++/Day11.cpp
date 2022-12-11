#include "AH.h"

namespace Day11
{

	struct Monkey
	{
		std::vector<uint64_t> items;
		const std::string operation;
		const uint64_t op_value;
		const uint64_t test;
		const size_t if_true;
		const size_t if_false;
		uint64_t counter;

		Monkey(const std::vector<uint64_t> items, const std::string op, const uint64_t val,
			const uint64_t test, const size_t t, const size_t f);
		std::pair<uint64_t, size_t> MonkeySee(const bool worry);
	};

	Monkey::Monkey(const std::vector<uint64_t> items, const std::string op, const uint64_t val,
			const uint64_t test, const size_t t, const size_t f) :
		items(items), operation(op), op_value(val), test(test), if_true(t), if_false(f), counter(0)
		{}

	std::pair<uint64_t, size_t> Monkey::MonkeySee(const bool worry)
	{
		counter++;

		auto value = items[0];
		items.erase(items.begin());

		if (operation == "+") {
			if (op_value == 0)
				value += value;
			else
				value += op_value;
		} else {
			if (op_value == 0) 
				value *= value;
			else 
				value *= op_value;
		}

		if (worry) {
			value /= 3;
		}

		if ((value % test) == 0)
			return std::make_pair(value, if_true);
		else
			return std::make_pair(value, if_false);
	}

	std::vector<Monkey> ParseInput(const std::vector<std::string> & ss, const char c)
	{
		std::vector<Monkey> ms;

		for (auto s : ss)
		{
			const auto lines = AH::Split(s, c);
			const auto strItems = AH::SplitOnString(lines[1].substr(18), ", ");
			std::vector<uint64_t> tempItems;
			for (auto item : strItems)
				tempItems.push_back(std::stoi(item));

			const std::string tempOp = lines[2].substr(23,1);
			std::string tempOpValueStr = lines[2].substr(25);
			uint64_t tempOpValue = 0;
			if (tempOpValueStr.front() != 'o')
				tempOpValue = std::stoi(tempOpValueStr);

			const uint64_t tempDiv   = std::stoi(lines[3].substr(21));
			const size_t tempTrue  = std::stoi(lines[4].substr(29));
			const size_t tempFalse = std::stoi(lines[5].substr(30));

			ms.emplace_back(tempItems, tempOp, tempOpValue, tempDiv, tempTrue, tempFalse);
		}

		return ms;
	}

	void MonkeyDo(std::vector<Monkey> & ms, const bool worry)
	{
		uint64_t reducer = 1;
		for (const auto & m : ms)
			reducer *= m.test;

		for (auto & m : ms)
		{
			const uint64_t item_count = m.items.size();
			for (uint64_t i = 0; i < item_count; ++i)
			{
				const auto monkeyMove = m.MonkeySee(worry);
				ms[monkeyMove.second].items.push_back(monkeyMove.first % reducer);
			}
		} 
	}

	uint64_t MonkeyBusiness(const std::vector<Monkey> & ms)
	{
		std::vector<uint64_t> business;
		for (const auto m : ms)
			business.push_back(m.counter);

		std::sort(business.begin(), business.end(), std::greater{});
		return business[0] * business[1];
	}

	uint64_t Run(const std::string& filename)
	{
		const auto inputLines = AH::ReadTextFile(filename);
		const auto lines = AH::ParseLineGroups(inputLines, '|');
		auto monkeys1 = ParseInput(lines, '|');
		auto monkeys2 = ParseInput(lines, '|');
		
		for (uint64_t i = 0; i < 20; ++i) {
			MonkeyDo(monkeys1, true);
		}
		for (uint64_t i = 0; i < 10000; ++i) {
			MonkeyDo(monkeys2, false);
		}

		AH::PrintSoln(11, MonkeyBusiness(monkeys1), MonkeyBusiness(monkeys2));


		return 0;
	}

}
