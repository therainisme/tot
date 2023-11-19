#pragma once

#include <chrono>
#include <cstdio>
#include <string>

class Timer;

class Timer {
 public:
  explicit Timer(const std::string &name = "", const int indent = 0) noexcept;
  ~Timer();

  Timer(const Timer &&) = delete;
  Timer(const Timer &) = delete;
  Timer &operator=(const Timer &) = delete;
  Timer &operator=(const Timer &&) = delete;

 private:
  const std::string name_;
  const int indent_;
  const std::chrono::time_point<std::chrono::steady_clock> start_;
};

Timer::Timer(const std::string &name, const int indent) noexcept
    : name_(name), indent_(indent), start_(std::chrono::steady_clock::now()) {}

Timer::~Timer() {
  auto end = std::chrono::steady_clock::now();
  auto duration =
      std::chrono::duration_cast<std::chrono::milliseconds>(end - start_);
  for (int i = 0; i < indent_; ++i) {
    printf("  ");
  }

  if (!name_.empty()) {
    printf("%s: ", name_.c_str());
  }
  printf("%ld ms\n", duration.count());
}