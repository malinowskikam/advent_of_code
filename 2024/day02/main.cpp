#include <cmath>
#include <iostream>
#include <ostream>
#include <sstream>
#include <string>
#include <util.hpp>

void part1();
void part2();
int sign(int value);
bool check_line(std::string &line);
bool check_line_with_dampner(std::string &line,
                             std::optional<int> dropped_index);

int main() {
    part1();
    part2();
}

void part1() {
    int count = 0;

    std::ifstream input = open_input("input/input02.txt");
    std::string line;
    while (std::getline(input, line)) {
        if (check_line(line)) {
            count++;
        }
    }

    std::cout << count << std::endl;
}

void part2() {
    int count = 0;

    std::ifstream input = open_input("input/input02.txt");
    std::string line;
    while (std::getline(input, line)) {
        if (check_line_with_dampner(line, {})) {
            count++;
        }
    }

    std::cout << count << std::endl;
}

bool check_line(std::string &line) {
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
                    return false;
                }
            } else {
                last_sign = current_sign;
            }

            // Check value
            if (abs(n - *last_n) > 3) {
                return false;
            }
        }
        last_n = n;
    }
    return true;
}

bool check_line_with_dampner(std::string &line,
                             std::optional<int> dropped_index) {
    std::optional<int> last_sign;
    std::optional<int> last_n;
    int n;
    std::istringstream ints(line);
    bool done = true;

    int i = 0;
    while (ints >> n) {
        if (dropped_index == i) {
            i++;
            continue;
        }

        if (last_n) {
            // Check sign
            int current_sign = sign(n - *last_n);
            if (last_sign) {
                if (current_sign != last_sign) {
                    // Check if the line is correct if we drop any of the
                    // previous values (just last 2 is not enough)
                    bool with_damped = false;
                    for (int j = 0; j <= i && !dropped_index && !with_damped;
                         j++) {
                        with_damped =
                            with_damped || check_line_with_dampner(line, j);
                    }
                    return with_damped;
                }
            } else {
                last_sign = current_sign;
            }

            // Check value
            if (abs(n - *last_n) > 3) {
                // Check if the line is correct if we drop any of the previous
                // values (just last 2 is not enough)
                bool with_damped = false;
                for (int j = 0; j <= i && !dropped_index && !with_damped; j++) {
                    with_damped =
                        with_damped || check_line_with_dampner(line, j);
                }
                return with_damped;
            }
        }
        last_n = n;
        i++;
    }
    return true;
}

int sign(int value) { return (value > 0) - (value < 0); }
