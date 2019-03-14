package utils 

import (
  "testing"
)

func TestSelector1(t *testing.T) {
  func_sig := "identity(uint256)"
  expected := "0xac37eebb"
  actual := Selector("identity", "uint256")
  testError(t, func_sig, expected, actual)
}

func TestSelector2(t *testing.T) {
  func_sig := "initialize(address,address)"
  expected := "0x485cc955"
  actual := Selector("initialize", "address", "address")
  testError(t, func_sig, expected, actual)
}

func testError(t *testing.T, func_sig, expected, actual string) {
  if expected != actual {
    t.Errorf("Expected the hash of \"%v\" to be \"%v\" but instead got \"%v\"", func_sig, expected, actual)
  }
}
