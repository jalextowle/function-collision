// TODO: These tests do not get full test coverage. They use real ABIs, but should be augmented
package utils

import (
  "reflect"
  "testing"
)

/* Param Decoding Tests */

func TestParam1(t *testing.T) {
  contents := `{
    "name": "",
    "type": "address"
  }` 

  expected := Param{
    Name: "",
    Type: "address",
  }

  actual := DecodeParam([]byte(contents))
  abiTestError(t, contents, expected, actual)
}

func TestParam2(t *testing.T) {
  contents := `{
    "name": "_targetId",
    "type":"bytes32"
  }` 

  expected := Param{
    Name: "_targetId",
    Type: "bytes32",
  }

  actual := DecodeParam([]byte(contents))
  abiTestError(t, contents, expected, actual)
}

/* ABI Decoding Tests */

func TestABI1(t *testing.T) {
  contents := `{
    "constant": false,
    "inputs": [],
    "name": "pause",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  }`

  expected := ABI{
    Constant: false,
    Inputs: []Param{},
    Name: "pause",
    Outputs: []Param{},
    Payable: false,
    StateMutability: "nonpayable",
    Type: "function",
  }

  actual := DecodeABI([]byte(contents))
  abiTestError(t, contents, expected, actual)
}

func TestABI2(t *testing.T) {
  contents := ` {
    "constant": false,
    "inputs": [
      {
        "name": "",
        "type": "address"
      }
    ],
    "name": "changeController",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  }`

  expected := ABI{
    Constant: false,
    Inputs: []Param{
      { Name: "", Type: "address" },
    }, 
    Name: "changeController",
    Outputs: []Param{},
    Payable: false,
    StateMutability: "nonpayable",
    Type: "function",
  }

  actual := DecodeABI([]byte(contents))
  abiTestError(t, contents, expected, actual)
}

/* Contract ABI Decoding Tests */

func TestContractABI1(t *testing.T) {
  contents := `[
  {
    "constant": false,
    "inputs": [],
    "name": "pause",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "",
        "type": "address"
      }
    ],
    "name": "changeController",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  }
]`

  expected := []ABI{
    {
      Constant: false,
      Inputs: []Param{},
      Name: "pause",
      Outputs: []Param{},
      Payable: false,
      StateMutability: "nonpayable",
      Type: "function",
    },
    {
      Constant: false,
      Inputs: []Param{
        { Name: "", Type: "address" },
      }, 
      Name: "changeController",
      Outputs: []Param{},
      Payable: false,
      StateMutability: "nonpayable",
      Type: "function",
    },
  }

  actual := DecodeABIToSlice([]byte(contents))
  abiTestError(t, contents, expected, actual)
}

/* Truffle ABI Decoding Tests */

// TODO
func TestTruffleABI1(t *testing.T) {

}

/* Test Helpers */

// This function simplifies the logic of all of the test functions. The expected and actual values are allowed to be 
// general interfaces so that this function can be used in the tests of all functions, regardless of the return type
func abiTestError(t *testing.T, contents string, expected, actual interface{}) {
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("Expected the decoding of %v to equal %v but instead got %v", contents, expected, actual)
  }
}
