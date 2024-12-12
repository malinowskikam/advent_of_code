#include <boost/container_hash/hash.hpp>
#include <boost/unordered_map.hpp>
#include <boost/unordered_set.hpp>
#include <iostream>
#include <util.hpp>

void part1();
void part2();

typedef struct point_struct {
    int x, y;

    point_struct operator-(const point_struct &rhs) const {
        return point_struct{x - rhs.x, y - rhs.y};
    }

    point_struct operator+(const point_struct &rhs) const {
        return point_struct{x + rhs.x, y + rhs.y};
    }

    bool operator==(const point_struct &rhs) const {
        return x == rhs.x && y == rhs.y;
    }
} Point;

namespace boost {
template <> struct hash<Point> {
    std::size_t operator()(const Point &s) const {
        std::size_t seed = 0;

        boost::hash_combine(seed, s.x);
        boost::hash_combine(seed, s.y);

        return seed;
    }
};
}; // namespace boost

typedef struct grid_struct {
    int width, height;

    bool is_in_bounds(const Point &p) const {
        return p.x >= 0 && p.y >= 0 && p.x < height && p.y < width;
    }
} Grid;

int main() {
    part1();
    part2();
}

void part1() {
    long result = 0;

    Grid grid{0, 0};

    boost::unordered_set<Point> antinodes;
    boost::unordered_map<char, boost::unordered_set<Point>> antennas;

    std::string line;
    std::ifstream is = open_input("input/input08.txt");
    while (std::getline(is, line)) {
        int row_len = 0;
        for (char c : line) {
            if (c != '.') {
                antennas[c].insert(Point{grid.height, row_len});
            }

            row_len++;
            if (grid.height == 0) {
                grid.width = row_len;
            }
        }
        grid.height++;
    }

    for (const auto &p : antennas) {
        auto antennas_set = p.second;
        for (const auto &a1 : antennas_set) {
            for (const auto &a2 : antennas_set) {
                if (a1 != a2) {
                    Point node1 = a1 + (a1 - a2);
                    if (grid.is_in_bounds(node1)) {
                        antinodes.insert(node1);
                    }

                    Point node2 = a2 + (a2 - a1);
                    if (grid.is_in_bounds(node2)) {
                        antinodes.insert(node2);
                    }
                }
            }
        }
    }

    result = antinodes.size();
    std::cout << result << std::endl;
}

void part2() {
    long result = 0;

    Grid grid{0, 0};

    boost::unordered_set<Point> antinodes;
    boost::unordered_map<char, boost::unordered_set<Point>> antennas;

    std::string line;
    std::ifstream is = open_input("input/input08.txt");
    while (std::getline(is, line)) {
        int row_len = 0;
        for (char c : line) {
            if (c != '.') {
                antennas[c].insert(Point{grid.height, row_len});
            }

            row_len++;
            if (grid.height == 0) {
                grid.width = row_len;
            }
        }
        grid.height++;
    }

    for (const auto &p : antennas) {
        auto antennas_set = p.second;
        for (const auto &a1 : antennas_set) {
            for (const auto &a2 : antennas_set) {
                if (a1 != a2) {
                    Point node1 = a1;
                    Point diff1 = a1 - a2;
                    while (grid.is_in_bounds(node1)) {
                        antinodes.insert(node1);
                        node1 = node1 + diff1;
                    }

                    Point node2 = a2;
                    Point diff2 = a2 - a1;
                    while (grid.is_in_bounds(node2)) {
                        antinodes.insert(node2);
                        node2 = node2 + diff2;
                    }
                }
            }
        }
    }

    result = antinodes.size();
    std::cout << result << std::endl;
}
