package utils 

import (
  "testing"
)

/* Function Selector from String Arguments */

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

/* Function Selector from ABI */

func TestFunctionSelectorABI1(t *testing.T) {
  func_sig := "identity(uint256)"
  expected := "0xac37eebb"
  abi := ABI{
    Name: "identity",
    Inputs: []Param{
      { Name: "", Type: "uint256" }, 
    }, 
  }
  actual := FunctionSelectorABI(abi)
  cryptoTestError(t, func_sig, expected, actual)
}

func TestFunctionSelectorABI2(t *testing.T) {
  func_sig := "initialize(address,address)"
  expected := "0x485cc955"
  abi := ABI{
    Name: "initialize",
    Inputs: []Param{
      { Name: "", Type: "address" },
      { Name: "", Type: "address" },
    },
  }
  actual := FunctionSelectorABI(abi)
  cryptoTestError(t, func_sig, expected, actual)
}

/* Test Helpers */

func cryptoTestError(t *testing.T, func_sig, expected, actual string) {
  if expected != actual {
    t.Errorf("Expected the hash of \"%v\" to be \"%v\" but instead got \"%v\"", func_sig, expected, actual)
  }
}
