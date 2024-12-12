#include <cassert>
#include <cstddef>
#include <iostream>
#include <util.hpp>
#include <vector>

void part1();
void part2();

int main() {
    part1();
    part2();
}

void part1() {
    long result = 0;

    std::string input = read_input("input/input09.txt");
    assert(input.size() % 2 == 1);

    std::vector<char> files;
    std::vector<char> spaces;

    for (size_t i = 0; i < input.size(); i += 2) {
        assert(input[i] >= '0' && input[i] <= '9');
        files.push_back(input[i] - '0');

        if (i < input.size() - 1) {
            assert(input[i + 1] >= '0' && input[i + 1] <= '9');
            spaces.push_back(input[i + 1] - '0');
        }
    }

    size_t front_idx = 0;
    size_t back_idx = files.size() - 1;
    size_t space_idx = 0;

    bool is_space = false;
    int i = 0;

    std::cout << "Seria: " << std::endl;
    while (true) {
        if (is_space) {
            if(spaces[space_idx] > 0) {
                if(files[back_idx] > 0) {
                    files[back_idx]--;
                    spaces[space_idx]--;
                    result += i * back_idx;
                    i++;
                } else {
                    back_idx--;
                }
            } else {
                space_idx++;
                is_space = false;
            }
        } else {
            if(files[front_idx] > 0) {
                files[front_idx]--;
                result += i * front_idx;
                i++;
            } else {
                front_idx++;
                is_space = true;
            }
        }

        if(front_idx == back_idx && files[front_idx] == 0) {
            break;
        }
    }
    
    std::cout << result << std::endl;
}

void part2() {
    long result = 0;

    std::cout << result << std::endl;
}
