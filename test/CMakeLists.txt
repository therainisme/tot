add_executable(timer_test timer_test.cpp)

add_dependencies(timer_test gtest)
include_directories(${gtest_SOURCE_DIR}/include ${gtest_SOURCE_DIR})
target_link_libraries(timer_test GTest::gtest_main)
target_link_libraries(timer_test timer)