#include "merkle/merkle.h"

#include <gtest/gtest.h>

#include <iostream>

TEST(MerkleTest, Merkle) {
  merkle::Tree::Hash hash(
      "fa8f44eabb728d4020e7f33d1aa973faaef19de6c06679bccdc5100a3c01f54a");
  merkle::Tree tree;
  tree.insert(hash);
  auto root = tree.root();
  auto path = tree.path(0);
  assert(path->verify(root));
}