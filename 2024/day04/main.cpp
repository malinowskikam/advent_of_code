#include <iostream>
#include <string>
#include <unordered_map>
#include <util.hpp>
#include <vector>

void part1();
void part2();
int get_xmas_count_from_pos(std::vector<std::string> &grid, int i, int j);
bool check_mas_direction(std::vector<std::string> &grid, int i, int j, int dir_i, int dir_j);
bool check_x_mas(std::vector<std::string> &grid, int i, int j);

int main() {
    part1();
    part2();
}

void part1() {
    int count = 0;
    
    std::vector<std::string> grid;

    std::ifstream input = open_input("input/input04.txt");
    std::string line;
    while (std::getline(input, line)) {
        grid.push_back(line);
    }

    for (int i = 0; i < grid.size(); i++) {
        for (int j = 0; j < grid[i].size(); j++) {
            count += get_xmas_count_from_pos(grid, i, j);
        }
    }
    
    std::cout << count << std::endl;
}

void part2() {
    int count = 0;
    
    std::vector<std::string> grid;

    std::ifstream input = open_input("input/input04.txt");
    std::string line;
    while (std::getline(input, line)) {
        grid.push_back(line);
    }

    for (int i = 0; i < grid.size(); i++) {
        for (int j = 0; j < grid[i].size(); j++) {
            if (check_x_mas(grid, i, j)) {
                count++;
            }
        }
    }
    
    std::cout << count << std::endl;
}

/**
 * @brief Check how many instances of word XMAS are present in all 8 directions
 * 
 * @param grid Main grid
 * @param i X row
 * @param j X column
 * @return int Number of instances of word XMAS starting in position i,j
 */
int get_xmas_count_from_pos(std::vector<std::string> &grid, int i, int j) {
    if (grid[i][j] != 'X') {
        return 0;
    }

    int count = 0;
    for (int dir_i = -1; dir_i <= 1; dir_i++) {
        for (int dir_j = -1; dir_j<= 1; dir_j++) {
            if ((dir_i != 0 || dir_j != 0) && check_mas_direction(grid, i, j, dir_i, dir_j)) {
                count ++;
            }
        }
    }

    return count;
}

/**
 * @brief Check if the letters M A S are in the required positions
 * 
 * @param grid Main grid
 * @param i X Row
 * @param j X Column
 * @param dir_i Direction of checking (row). Must be -1, 0 or 1.
 * @param dir_j Direction of checking (column). Must be -1, 0 or 1.

 * @return true if the letters M A and S are in the required positions
 */
bool check_mas_direction(std::vector<std::string> &grid, int i, int j, int dir_i, int dir_j) {
    int m_pos_i = i + dir_i;
    int m_pos_j = j + dir_j;
    int a_pos_i = i + dir_i * 2;
    int a_pos_j = j + dir_j * 2;
    int s_pos_i = i + dir_i * 3;
    int s_pos_j = j + dir_j * 3;
    
    
    return s_pos_i >= 0 && s_pos_i < grid.size()
        && s_pos_j >= 0 && s_pos_j < grid[s_pos_i].size()
        && grid[m_pos_i][m_pos_j] == 'M'
        && grid[a_pos_i][a_pos_j] == 'A'
        && grid[s_pos_i][s_pos_j] == 'S';
}

/**
 * @brief Check if valid X-MAS pattern is centered on position i,j
 * 
 * @param grid Main grid
 * @param i row position of center A
 * @param j column position of center A
 * @return true if there is a valid X-MAS pattern
 */
bool check_x_mas(std::vector<std::string> &grid, int i, int j) {
    if (grid[i][j] != 'A' || i < 1 || i > grid.size() - 2 || j < 1 && j > grid[i].size() - 2) {
        return false;
    }

    std::unordered_map<char, int> counts;

    counts[grid[i-1][j-1]]++;
    counts[grid[i-1][j+1]]++;
    counts[grid[i+1][j-1]]++;
    counts[grid[i+1][j+1]]++;


    return counts['M'] == 2 && counts['S'] == 2 && grid[i-1][j-1] != grid[i+1][j+1];
}