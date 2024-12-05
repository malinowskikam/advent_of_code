#include <cstddef>
#include <iterator>
#include <util.hpp>

#include <fstream>
#include <iostream>
#include <string>

void panic(const std::string& message) {
    std::cerr << message << std::endl;
    exit(1);
}

std::ifstream open_input(const std::string& filename) {
    std::ifstream file(filename);

    if (!file.is_open()) {
        panic(std::string("Failed to open input file: ") + filename);
    }

    return std::ifstream(filename);
}

std::string read_input(const std::string& filename) {
    std::ifstream input = open_input(filename);

    return std::string(std::istreambuf_iterator<char>(input),
                       std::istreambuf_iterator<char>());
}

std::vector<std::string> split(const std::string& str, const std::string& delim) {
    size_t pos_start = 0, pos_end, delim_len = delim.length();
    std::string token;
    std::vector<std::string> res;

    while ((pos_end = str.find(delim, pos_start)) != std::string::npos) {
        token = str.substr (pos_start, pos_end - pos_start);
        pos_start = pos_end + delim_len;
        res.push_back (token);
    }

    res.push_back (str.substr (pos_start));
    return res;
}