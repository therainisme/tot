add_library(timer SHARED timer.cpp)

target_include_directories(timer INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/..)