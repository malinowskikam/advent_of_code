#include <cassert>
#include <iostream>
#include <string>
#include <unordered_map>
#include <unordered_set>
#include <util.hpp>
#include <vector>

void part1();
void part2();
bool check_update(std::unordered_map<int, std::unordered_set<int>> &rules,
                  const std::vector<int> &before, int current);
void sort_update(std::unordered_map<int, std::unordered_set<int>> &rules,
                 std::vector<int> &update);

int main() {
    part1();
    part2();
}

void part1() {
    int sum = 0;

    // Rules are represented by mapping a page to a set containing all pages
    // that cannot be after it.
    std::unordered_map<int, std::unordered_set<int>> rules;

    std::ifstream input = open_input("input/input05.txt");
    std::string line;
    while (std::getline(input, line)) {
        if (line == "") {
            break;
        }

        std::vector<std::string> parts = split(line, "|");
        assert(parts.size() == 2);
        rules[std::stoi(parts[1])].insert(std::stoi(parts[0]));
    }

    while (std::getline(input, line)) {
        std::vector<int> update;
        std::vector<std::string> parts = split(line, ",");
        bool valid = true;
        for (int i = 0; i < parts.size(); i++) {
            int current = std::stoi(parts[i]);
            if (!check_update(rules, update, current)) {
                valid = false;
                break;
            }
            update.push_back(current);
        }

        if (valid) {
            sum += update[update.size() / 2];
        }
    }

    std::cout << sum << std::endl;
}

void part2() {
    int sum = 0;

    // Rules are represented by mapping a page to a set containing all pages
    // that cannot be after it.
    std::unordered_map<int, std::unordered_set<int>> rules;

    std::ifstream input = open_input("input/input05.txt");
    std::string line;
    while (std::getline(input, line)) {
        if (line == "") {
            break;
        }

        std::vector<std::string> parts = split(line, "|");
        assert(parts.size() == 2);
        rules[std::stoi(parts[1])].insert(std::stoi(parts[0]));
    }

    while (std::getline(input, line)) {
        std::vector<int> update;
        std::vector<std::string> parts = split(line, ",");
        bool valid = true;
        for (int i = 0; i < parts.size(); i++) {
            int current = std::stoi(parts[i]);
            if (valid && !check_update(rules, update, current)) {
                valid = false;
            }
            update.push_back(current);
        }

        if (!valid) {
            sort_update(rules, update);
            for (int i = 0; i < update.size(); i++) {
                for (int j = 0; j < i; j++) {
                    if (rules[update[j]].contains(update[i])) {
                        // cos nie tak

                        std::cout << "Dupa" << std::endl;
                    }
                }
            }
            sum += update[update.size() / 2];
        }
    }

    std::cout << sum << std::endl;
    ;
}

bool check_update(std::unordered_map<int, std::unordered_set<int>> &rules,
                  const std::vector<int> &before, int current) {
    for (int i = 0; i < before.size(); i++) {
        if (rules[before[i]].contains(current)) {
            return false;
        }
    }
    return true;
}

void sort_update(std::unordered_map<int, std::unordered_set<int>> &rules,
                 std::vector<int> &update) {
    for (int i = 0; i < update.size(); i++) {
        int current = update[i];

        // check for first invalid index
        int j = 0;
        while (j < i && !rules[update[j]].contains(current)) {
            j++;
        }

        // moving all elements from j to i - 1 one position forward
        for (int k = i - 1; k >= j; k--) {
            update[k + 1] = update[k];
        }

        // inserting current at j
        update[j] = current;
    }
}