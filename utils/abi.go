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

/* ABI Processing */

type parameter struct {
  type_name  string
  stack_size int
}

// This function will return an abi by using the function names within the original abi to 
// create new function abis with different combinations of parameters made up of the 
// specified types.
func ExpandContractABI(param_num int, types []parameter, original []ABI) []ABI {
  if abi == nil {
    return
  }

  // This hash table will be used to ensure that the ABI returned by this contract will not have
  // redundancies 
  abi_set := make(map[ABI]bool)
  // Fill up the hash table with the existing ABI
  for _, abi := range(original) {
    abi_set[abi] = true
    // If the type of the abi object is "function", then add all of the new function abis that can
    // be made with the constraints to the abi
    if abi.Type == "function" {
      combinations := ABICombinations(param_num, types, ABI.Name)
      for _, combination := range(combinations) {
        new_abi := abi
        new_abi.Inputs = combination
        abi_set[new_abi] = true
      }
    }
  }

}

// FIXME: This function has not been tested but it also does not return all of the permutations
//        of these parameter lists. This will need to be done by the functions using these combinations
func ABICombinations(param_num int, types []parameter, name string) [][]Param {
  var result [][]Param
  if types == nil || param_num == 0 {
    return [][]Param{} 
  }

  for i := 0; i <= param_num; i += types.stack_size {
    combinations := ABICombinations(param_num - i, types[1:], name) 
    var type_params []Param 
    for j := 0; j <= i; j += types.stack_size {
      type_params[j] = Param{
        Name: "",
        Type: types[0],
      }
    }
    for idx, _ := range(combinations) {
      combinations[idx] = append(type_params, combinations...)
    }
    result = append(result, combinations) 
  }
  return result
}
