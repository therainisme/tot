include_directories(${gtest_SOURCE_DIR}/include ${gtest_SOURCE_DIR})

add_executable(merkle_test merkle_test.cpp)

add_dependencies(merkle_test gtest)
target_link_libraries(merkle_test GTest::gtest_main)
target_link_libraries(merkle_test merkle)

add_executable(timer_test timer_test.cpp)

add_dependencies(timer_test gtest)

target_link_libraries(timer_test GTest::gtest_main)
target_link_libraries(timer_test timer)