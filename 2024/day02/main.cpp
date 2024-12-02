#include <cmath>
#include <iostream>
#include <sstream>
#include <string>
#include <util.hpp>

void part1();
void part2();
int sign(int value);

int main() {
    part1();
    part2();
}

void part1() {
    int count = 0;

    std::ifstream input = open_input("input/input02.txt");
    std::string line;
    while (std::getline(input, line)) {
        std::optional<int> last_sign;
        std::optional<int> last_n;
        int n;
        std::istringstream ints(line);
        bool done = true;
        while (ints >> n) {
            if (last_n) {
                // Check sign
                int current_sign = sign(n - *last_n);
                if (last_sign) {
                    if (current_sign != last_sign) {
                        done = false;
                        break;
                    }
                } else {
                    last_sign = current_sign;
                }

                // Check value
                if (abs(n - *last_n) > 3) {
                    done = false;
                    break;
                }
            }
            last_n = n;
        }
        if (done) {
            count++;
        }
    }

    std::cout << count << std::endl;
}

void part2() {
    int count = 0;

    std::cout << count << std::endl;
}

int sign(int value) { return (value > 0) - (value < 0); }