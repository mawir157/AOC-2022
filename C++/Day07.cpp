#include "AH.h"

namespace Day07
{

	struct Node
	{
		std::string name;
		unsigned int size;
		bool is_file;
		Node * parent;
		std::vector<Node *> children;

		int dirSize() const;
		Node * changeDir(const std::string s);
		void addNode(const std::string name, const unsigned int size, const bool is_file);
		void getDirSizes(std::map<std::string, unsigned int> & m) const;
	};

	int Node::dirSize() const
	{
		if (is_file)
		{
			return size;
		}
		else
		{
			unsigned int sum = 0;
			for (const auto c : children)
				sum += c->dirSize();

			return sum;
		}
	}

	Node * Node::changeDir(const std::string s)
	{
		if (s == "..")
			return parent;

		for (const auto c : children)
			if (c->name == s)
				return c;

		return this;
	}

	void Node::addNode(const std::string name, const unsigned int size, const bool is_file)
	{
		Node * n = new Node;
		n->name = name;
		n->size = size;
		n->is_file = is_file;
		n->parent = this;
		n->children.clear();
		children.push_back(n);
	}

	void parseCLI(const std::vector<std::string> ss, Node * n)
	{
		for (const auto s: ss)
		{
			if (s == "$ ls")
			{
				// do nothing
			}
			else if (s.substr(0, 4) == "$ cd")
			{
				const auto data = AH::Split(s, ' ');
				n = n->changeDir(data[2]);
			}
			else
			{
				const auto data = AH::Split(s, ' ');
				if (data[0] == "dir")
					n->addNode(data[1], 0, false);
				else
					n->addNode(data[1], std::stoi(data[0]), true);
			}
		}
	}

	void Node::getDirSizes(std::map<std::string, unsigned int> & m) const
	{
		if (is_file)
			return;
		
		std::string full_name = "root";
		if (parent != nullptr)
			full_name = parent->name + "|" + name;

		m[full_name] = dirSize();

		for (const auto c: children) 
			c->getDirSizes(m);

		return;
	}

	int Run(const std::string& filename)
	{
		auto lines = AH::ReadTextFile(filename);
		lines.erase(lines.begin());
		Node root = Node();
		root.name = "root";
		root.size = 0;
		root.is_file = false;
		root.parent = nullptr;
		root.children.clear();

		std::map<std::string, unsigned int> m;

		parseCLI(lines, & root);
		root.getDirSizes(m);

		unsigned int part1 = 0;
		unsigned int part2 = m["root"];
		const unsigned int toDelete = m["root"] - 40000000;

		for (auto const& [k, v] : m)
		{
			if ((v > toDelete) && (v < part2))
				part2 = v;

			if (v <= 100000)
				part1 += v;
		}

		AH::PrintSoln(7, part1, part2);

		return 0;
	}

}
