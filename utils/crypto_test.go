package utils 

import (
  "testing"
)

func TestFunctionSelector1(t *testing.T) {
  func_sig := "identity(uint256)"
  expected := "0xac37eebb"
  actual := FunctionSelector("identity", "uint256")
  cryptoTestError(t, func_sig, expected, actual)
}

func TestFunctionSelector2(t *testing.T) {
  func_sig := "initialize(address,address)"
  expected := "0x485cc955"
  actual := FunctionSelector("initialize", "address", "address")
  cryptoTestError(t, func_sig, expected, actual)
}

func cryptoTestError(t *testing.T, func_sig, expected, actual string) {
  if expected != actual {
    t.Errorf("Expected the hash of \"%v\" to be \"%v\" but instead got \"%v\"", func_sig, expected, actual)
  }
}
