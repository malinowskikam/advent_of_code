#include <cassert>
#include <iostream>
#include <regex>
#include <string>
#include <util.hpp>

void part1();
void part2();

int main() {
    part1();
    part2();
}

void part1() {
    int sum = 0;

    std::regex exp(R"x(mul\((-?\d+),(-?\d+)\))x");
    std::string input = read_input("input/input03.txt");

    std::sregex_iterator begin(input.begin(), input.end(), exp);
    std::sregex_iterator end;

    for (auto it = begin; it != end; ++it) {
        std::smatch match = *it;

        assert(match.size() == 3);
        int l = std::stoi(match[1]);
        int r = std::stoi(match[2]);

        sum += l*r;
    }

    std::cout << sum << std::endl;
}

void part2() {
    int sum = 0;

    std::cout << sum << std::endl;
}
