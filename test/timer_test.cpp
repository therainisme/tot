#include "timer/timer.h"

#include <gtest/gtest.h>

#include <cstdio>

TEST(TimerTest, Timer) {
  { Timer timer("Call A"); }
  { Timer timer("Do something in A", 1); }
  { Timer timer("End A"); }

  { Timer timer; }
}