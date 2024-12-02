#include <iostream>
#include <util.hpp>
#include <vector>
#include <algorithm>
#include <math.h>

void part1();
void part2();

int main() {
    part1();
}

void part1() {
    int sum = 0;

    std::vector<int> left;
    std::vector<int> right;

    int l, r;

    std::ifstream input = open_input("input/input01.txt");
    while (input >> l >> r) {
        left.push_back(l);
        right.push_back(r);
    }
    
    std::sort(left.begin(), left.end());
    std::sort(right.begin(), right.end());

    for (int i = 0; i < left.size(); i++) {
        sum += abs(left[i] - right[i]);
    }

    std::cout << sum << std::endl;
}