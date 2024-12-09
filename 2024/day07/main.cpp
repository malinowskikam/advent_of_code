#include <cassert>
#include <iostream>
#include <string>
#include <util.hpp>
#include <vector>

void part1();
void part2();

bool check_line(const std::vector<long> &numbers, long value);
bool check_line_req(const std::vector<long> &numbers, long value,
                    long current_value,
                    std::vector<long>::size_type current_index);

int main() {
    part1();
    part2();
}

void part1() {
    long sum = 0;

    std::string line;
    std::ifstream is = open_input("input/input07.txt");
    while (std::getline(is, line)) {
        std::vector<std::string> split_val_and_nums = split(line, ": ");
        long value = std::stol(split_val_and_nums[0]);
        std::vector<long> numbers;
        for (std::string s : split(split_val_and_nums[1], " ")) {
            numbers.push_back(std::stol(s));
        }

        if (check_line(numbers, value)) {
            sum += value;
            std::cout << sum << std::endl;
        }
    }

    std::cout << sum << std::endl;
}

void part2() {
    int count = 0;

    std::cout << count << std::endl;
}

bool check_line(const std::vector<long> &numbers, long value) {
    assert(numbers.size() > 0 && numbers.size() <= 32);

    int n = 1 << (numbers.size() - 1);
    for (int i = 0; i < n; i++) {
        long result = numbers[0];
        int val = i;
        for (auto it = numbers.begin() + 1; it != numbers.end(); ++it) {
            if (val & 1) {
                result += *it;
            } else {
                result *= *it;
            }
            val >>= 1;
        }
        if (result == value) {
            return true;
        }
    }
    return false;
}
