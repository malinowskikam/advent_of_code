#include <cassert>
#include <iostream>
#include <string>
#include <util.hpp>
#include <vector>
#include <cmath>

void part1();
void part2();

bool check_line(const std::vector<long> &numbers, long value);
bool check_line2(const std::vector<long> &numbers, long value);

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
        }
    }

    std::cout << sum << std::endl;
}

void part2() {
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

        if (check_line2(numbers, value)) {
            sum += value;
        }
    }

    std::cout << sum << std::endl;
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

bool check_line2(const std::vector<long> &numbers, long value) {
    assert(numbers.size() > 0 && numbers.size() <= 32);

    int n = pow(3, numbers.size() - 1);
    for (int i = 0; i < n; i++) {
        long result = numbers[0];
        int val = i;
        
        for (auto it = numbers.begin() + 1; it != numbers.end(); ++it) {
            switch (val % 3) {
            case 0:
                result += *it;
                break;
            case 1:
                result *= *it;
                break;
            case 2:
                result = std::stol(std::to_string(result) + std::to_string(*it));
                break;
            }

            val /= 3;
        }

        if (result == value) {
            return true;
        }
    }
    return false;
}
