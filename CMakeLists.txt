cmake_minimum_required(VERSION 3.11...3.21)
project(
        ToT
        VERSION 1.0.0
        DESCRIPTION "Tools of Therainisme"
        LANGUAGES CXX
)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_subdirectory(third_party/googletest)
add_subdirectory(timer)
add_subdirectory(test)