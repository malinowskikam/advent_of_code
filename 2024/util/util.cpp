#include <util.hpp>

#include <fstream>
#include <iostream>
#include <string>

void panic(std::string message) {
    std::cerr << message << std::endl;
    exit(1);
}

std::ifstream open_input(std::string filename) {
    std::ifstream file(filename);

    if (!file.is_open()) {
        panic(std::string("Failed to open input file: ") + filename);
    }

    return std::ifstream(filename);
}