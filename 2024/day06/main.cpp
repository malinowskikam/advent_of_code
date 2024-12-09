#include <cassert>
#include <cstddef>
#include <fstream>
#include <iostream>
#include <unordered_set>
#include <util.hpp>
#include <vector>

#define WALL_FLAG 1
#define VISITED_UP_FLAG 2
#define VISITED_RIGHT_FLAG 4
#define VISITED_DOWN_FLAG 8
#define VISITED_LEFT_FLAG 16
#define VISITED_FLAG 0x1E

#define DIR_UP 0
#define DIR_RIGHT 1
#define DIR_DOWN 2
#define DIR_LEFT 3

void part1();
void part2();

typedef struct point_struct {
    int x, y;

    bool operator==(const point_struct &rhs) const {
        return x == rhs.x && y == rhs.y;
    }
} Point;

template <> struct std::hash<Point> {
    std::size_t operator()(const Point &p) const {
        std::size_t res = 0;
        std::hash<int> h;

        res ^= h(p.x) + 0x9e3779b9 + (res << 6) + (res >> 2);
        res ^= h(p.y) + 0x9e3779b9 + (res << 6) + (res >> 2);

        return res;
    }
};

typedef std::vector<std::vector<unsigned char>> Grid;

Grid build_grid(std::ifstream &is, Point &current);
Point get_next_point(int dir, const Point &current);
bool is_out_of_bounds(const Grid &grid, const Point &point);
std::unordered_set<Point> gather_visited(const Grid &grid, Point current,
                                         int dir);
bool test_grid_with_point(Grid grid, Point current, int dir,
                              Point extra_wall);

int main() {
    part1();
    part2();
}

void part1() {
    int dir = DIR_UP;
    Point current;
    std::ifstream input = open_input("input/input06.txt");
    Grid grid = build_grid(input, current);
    std::unordered_set<Point> visited = gather_visited(grid, current, dir);
    std::cout << visited.size() << std::endl;
}

void part2() {
    int count = 0;

    int dir = DIR_UP;
    Point current;
    std::ifstream input = open_input("input/input06.txt");
    Grid grid = build_grid(input, current);
    std::unordered_set<Point> visited = gather_visited(grid, current, dir);
    for (Point p : visited) {
        if (p != current && test_grid_with_point(grid, current, dir, p)) {
            count++;
        }
    }

    std::cout << count << std::endl;
}

Point get_next_point(int dir, const Point &current) {
    Point next;

    switch (dir) {
    case DIR_UP:
        next = {current.x - 1, current.y};
        break;
    case DIR_RIGHT:
        next = {current.x, current.y + 1};
        break;
    case DIR_DOWN:
        next = {current.x + 1, current.y};
        break;
    case DIR_LEFT:
        next = {current.x, current.y - 1};
        break;
    default:
        panic("Unknown dir");
    }

    return next;
}

bool is_out_of_bounds(const Grid &grid, const Point &point) {
    return point.x < 0 || point.y < 0 || point.x > (int)grid.size() - 1 ||
           point.y > (int)grid[point.x].size() - 1;
}

Grid build_grid(std::ifstream &is, Point &current) {
    Grid grid;
    std::string line;

    int i = 0;
    while (std::getline(is, line)) {
        std::vector<unsigned char> row;

        for (int j = 0; j < line.size(); j++) {
            unsigned char cell = 0;

            switch (line[j]) {
            case '.':
                break;
            case '#':
                cell |= WALL_FLAG;
                break;
            case '^':
                current = {i, j};
                cell |= VISITED_UP_FLAG;
                break;
            default:
                panic("Unknown char");
            }

            row.push_back(cell);
        }

        i++;
        grid.push_back(row);
    }

    return grid;
}

std::unordered_set<Point> gather_visited(const Grid &grid, Point current,
                                         int dir) {
    std::unordered_set<Point> visited;
    visited.insert(current);
    while (true) {
        Point next = get_next_point(dir, current);

        if (is_out_of_bounds(grid, next)) {
            break;
        }

        if (grid[next.x][next.y] & WALL_FLAG) {
            dir = (dir + 1) % 4;
        } else {
            visited.insert(next);
            current = next;
        }
    }

    return visited;
}

bool test_grid_with_point(Grid grid, Point current, int dir, Point extra_wall) {
    grid[extra_wall.x][extra_wall.y] |= WALL_FLAG;
    bool inf_loop = false;

    while (true) {
        Point next = get_next_point(dir, current);

        if (is_out_of_bounds(grid, next)) {
            break;
        }

        if (grid[next.x][next.y] & WALL_FLAG) {
            dir = (dir + 1) % 4;
        } else {
            if (grid[next.x][next.y] & (VISITED_UP_FLAG << dir)) {
                inf_loop = true;
                break;
            }
            grid[next.x][next.y] |= (VISITED_UP_FLAG << dir);
            current = next;
        }
    }

    return inf_loop;
}
