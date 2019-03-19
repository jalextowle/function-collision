package main

import (
  "reflect"
  "testing"
  "./utils"
)

func TestEmptyTarget(t *testing.T) {
  target := utils.ABI{}
  source := []utils.ABI{
    { Name: "a", Inputs: []utils.Param{}, Type: "function" },
  }

  var expected []utils.ABI
  actual := findCollisions(target, source)
  printErrorMessage(t, expected, actual)
}

func TestTrivialMatch(t *testing.T) {
  target := utils.ABI{
    Name: "a",
    Inputs: []utils.Param{},
    Type: "function",
  }

  source := []utils.ABI{
    { Name: "a", Inputs: []utils.Param{}, Type: "function" },
  }

  expected := source
  actual := findCollisions(target, source)
  printErrorMessage(t, expected, actual)
}

/* Test Helpers */

func printErrorMessage(t *testing.T, expected, actual []utils.ABI) {
  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("Expected the list of function selector collisions to be %v but instead got %v", expected, actual)
  }
}
