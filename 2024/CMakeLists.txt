cmake_minimum_required(VERSION 3.31)
project(advent_of_code_2023 CXX)
set (CMAKE_CXX_STANDARD 23)

find_package(Boost 1.45.0 REQUIRED) 

add_compile_options("-Wall" "-Wextra")

#Util library
add_subdirectory(util)

file(GLOB DAY_DIRS "day??")
foreach(DAY_DIR ${DAY_DIRS})
    get_filename_component(DAY_NAME ${DAY_DIR} NAME)
    add_executable(${DAY_NAME} ${DAY_DIR}/main.cpp)
    target_link_libraries(${DAY_NAME} util)
    target_include_directories(${DAY_NAME} PRIVATE ${CMAKE_SOURCE_DIR}/util)
    include_directories(${Boost_INCLUDE_DIRS})
    target_link_libraries(${DAY_NAME} ${Boost_LIBRARIES})
endforeach()
