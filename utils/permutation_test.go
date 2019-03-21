package utils

import (
  "reflect"
  "testing"
)

/* Permutation Generation Comprehensive Tests */

func TestZeroPermutations(t *testing.T) {
  max_size := 0
  expected := [][]int{}
  actual := GeneratePermutations(max_size)
  testPrintArray(t, max_size, expected, actual)
}

func TestOnePermutations(t *testing.T) {
  max_size := 1 
  expected := [][]int{{0}}
  actual := GeneratePermutations(max_size)
  testPrintArray(t, max_size, expected, actual)
}

func TestTwoPermutations(t *testing.T) {
  max_size := 2 
  expected := [][]int{
    {1, 0}, 
    {0, 1},
  }
  actual := GeneratePermutations(max_size)
  testPrintArray(t, max_size, expected, actual)
}

func TestThreePermutations(t *testing.T) {
  max_size := 3 
  expected := [][]int{
    {2, 1, 0},
    {1, 2, 0},
    {1, 0, 2},
    {2, 0, 1},
    {0, 2, 1},
    {0, 1, 2},
  }
  actual := GeneratePermutations(max_size)
  testPrintArray(t, max_size, expected, actual)
}

func TestFourPermutations(t *testing.T) {
  max_size := 4 
  expected := [][]int{
    // Join of {2, 1, 0} and 3
    {3, 2, 1, 0},
    {2, 3, 1, 0},
    {2, 1, 3, 0},
    {2, 1, 0, 3},
    // Join of {1, 2, 0} and 3
    {3, 1, 2, 0},
    {1, 3, 2, 0},
    {1, 2, 3, 0},
    {1, 2, 0, 3},
    // Join of {1, 0, 2} and 3
    {3, 1, 0, 2},
    {1, 3, 0, 2},
    {1, 0, 3, 2},
    {1, 0, 2, 3},
    // Join of {2, 0, 1} and 3
    {3, 2, 0, 1},
    {2, 3, 0, 1},
    {2, 0, 3, 1},
    {2, 0, 1, 3},
    // Join of {0, 2, 1} and 3
    {3, 0, 2, 1},
    {0, 3, 2, 1},
    {0, 2, 3, 1},
    {0, 2, 1, 3},
    // Join of {0, 2, 1} and 3
    {3, 0, 1, 2},
    {0, 3, 1, 2},
    {0, 1, 3, 2},
    {0, 1, 2, 3},
  }
  actual := GeneratePermutations(max_size)
  testPrintArray(t, max_size, expected, actual)
}

/* 
Permutation Generation Length Tests 

The size of the output for the GeneratePermutations function gets too large for me to write comprehensive unit tests by hand,
so for the time being I am just testing the length of the output of permutations of length greater than 4
*/

func TestFivePermuationsLength(t *testing.T) {
  max_size := 5
  expected := factorial(max_size)
  actual := len(GeneratePermutations(max_size))
  testPrintInt(t, max_size, expected, actual)
}

func TestSixPermuationsLength(t *testing.T) {
  max_size := 6 
  expected := factorial(max_size)
  actual := len(GeneratePermutations(max_size))
  testPrintInt(t, max_size, expected, actual)
}

func TestSevenPermuationsLength(t *testing.T) {
  max_size := 7 
  expected := factorial(max_size)
  actual := len(GeneratePermutations(max_size))
  testPrintInt(t, max_size, expected, actual)
}

/* Helpers */

func factorial(base int) int {
  if base < 0 {
    return 0
  }

  if base == 0 || base == 1 {
    return 1
  }

  return base * factorial(base - 1)
}

func testPrintArray(t *testing.T, max_size int, expected, actual [][]int) {
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("The generation of permutations with a max size of %v should have returned %v but instead returned %v", max_size, expected, actual)
  }
}

func testPrintInt(t *testing.T, max_size int, expected, actual int) {
  if actual != expected {
    t.Errorf("The generation of permutations with a max size of %v should have returned %v but instead returned %v", max_size, expected, actual)
  }
}
