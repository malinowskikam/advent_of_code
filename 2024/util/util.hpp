#include <fstream>
#include <vector>

void panic(const std::string& message);
std::ifstream open_input(const std::string& filename);
std::string read_input(const std::string& filename);
std::vector<std::string> split(const std::string& s, const std::string& delimiter);

