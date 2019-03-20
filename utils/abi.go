package utils

import (
  "encoding/json"
)

// Represents either an input of an output JSON
// object within a contract ABI
type Param struct {
  Name                 string
  Type                 string
}

// Represents a contract ABI JSON object
type ABI struct {
  Constant             bool
  Inputs               []Param
  Name                 string
  Outputs              []Param
  Payable              bool
  StateMutability      string
  Type                 string
}

// Represents a truffle ABI object that is produced when using the command
// `truffle compile`.
type TruffleABI struct {
  ContractName         string
  Abi                  []ABI 
  Bytecode             string 
  DeployedBytecode     string
  SourceMap            string 
  DeployedSourceMap    string
  Source               string
  SourcePath           string
  // TODO: Replace with an AST object struct
  Ast                  interface{}
  // TODO: Replace with the network object
  Networks             interface{}
  SchemaVersion        string
  UpdatedAt            string
}

/* Decoding */

func DecodeParam(contents []byte) Param {
  var param Param
  json.Unmarshal(contents, &param)
  return param
}

func DecodeABI(contents []byte) ABI {
  var abi ABI
  json.Unmarshal(contents, &abi)
  return abi
}

func DecodeABIToSlice(contents []byte) []ABI {
  var abi []ABI
  json.Unmarshal(contents, &abi)
  return abi
}

func DecodeTruffleABI(contents []byte) TruffleABI {
  var abi TruffleABI
  json.Unmarshal(contents, &abi) 
  return abi
}

/* ABI Creation */

type SolidityType struct {
  StackSize int 
  Name      string
}

// TODO: Use the permutation generators to perumte the ABIs
// Create a contract ABI that soley consists of functions that have the same name as
// the specfied name. This ABI will consist of all possible combinations of operator overloading 
// this function name using the specified types that is under the specified maximum stack 
// size.
func GenerateABIFromName(function_name string, max_stack_size int, types []SolidityType) (result []ABI) {
  if max_stack_size == 0 || len(types) == 0 {
    return []ABI{}
  }

  var combinations [][]string
  // Collect a list of all of the different combinations
  for i := 0; i < len(types); i++ {
    // If the current type's stack size is zero, skip it because it will cause an infinite loop
    if types[i].StackSize != 0 {
      name := types[i].Name 
      for j := len(combinations[i]); j < max_stack_size; j += types[i].StackSize {
        combinations[i][j] = name
      }
    }
  }

  var generated_permutations [][][]int
  var permutated_combinations [][]string

  // TODO: Much of this can probably be refactored into somewhere else
  // Permute the combinations and scan out redundancies
  for _, v := range(combinations) {
    if generated_permutations[len(v)] == nil {
      generated_permutations[len(v)] = generatePermutations(len(v))
    }
    permutated := applyPermutations(v, generated_permutations[len(v)])
    permutated_combinations = append(permutated_combinations, permutated...)
  }

  // TODO: Filter out redundant permutations
  // Generate the ABIs
  return 
}

// For all permutations of N numbers, you can add another number and insert it at every point to get the list of permutions.
// So we start with 1 number, and bootstrap from there
func generatePermutations(permutation_size int) (result [][]int) {
  if permutation_size < 1 {
    return [][]int{}
  }

  result = [][]int{{0}} 
  for i := 1; i < permutation_size; i++ {
    var joins [][]int
    for j := 0; j < len(result); j++ {
      joins = append(joins, joinPermutation(result[j], i)...)
    }
    result = joins
  }
  return
}

func joinPermutation(permutation []int, new_member int) (result [][]int) {
  for i := 0; i <= len(permutation); i++ {
    new_permutation := append(permutation[:i], append([]int{new_member}, permutation[i:]...)...)
    result = append(result, new_permutation)
  }
  return
}

func applyPermutation(original []string, permutation []int) (result []string) {
  if len(permutation) == len(original) {
    // TODO: How should this be handled?
    return original
  }

  for i, v := range(permutation) {
    result[i] = original[v]
  }
  return
}

func applyPermutations(original []string, permutations [][]int) (result [][]string) {
  for _, v := range(permutations) {
    result = append(result, applyPermutation(original, v))
  }
  return 
}
