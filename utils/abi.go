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
