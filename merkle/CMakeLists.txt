add_library(merkle SHARED merkle.cpp)

target_include_directories(merkle INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/..)