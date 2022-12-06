#include "AH.h"

namespace Day06
{
	
	size_t GetMarker(const std::string & s, const size_t n)
	{
		for (size_t i = 0; i < s.size() - n; ++i)
		{
			const auto p = s.substr(i, n);
			std::unordered_set<char> st;
			for (auto c : p)
				st.insert(c);

			if (st.size() == n)
				return i + n;
		}

		return 0;
	}

	int Run(const std::string& filename)
	{
		const auto input = AH::ReadTextFile(filename);
		
		AH::PrintSoln(6, GetMarker(input[0], 4), GetMarker(input[0], 14));

		return 0;
	}

}
