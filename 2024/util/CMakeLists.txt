add_library(util STATIC util.cpp util.hpp)
target_include_directories(util PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
