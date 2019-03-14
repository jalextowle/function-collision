package utils 

import (
  "testing"
)

func TestKeccak256(t *testing.T) {
  // This is the hash of the string "abc" returned by using the keccak256 function in Solidity
  expected := "4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45"
  actual := string(Keccak256("abc"))

  if actual != expected {
    t.Errorf("Expected the hash of \"abc\" to be %v but instead got %v", expected, actual)
  }
}
