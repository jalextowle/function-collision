package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "./utils"
)

func main() {
  args := os.Args
  if len(args) != 3 {
    fmt.Println("Usage: ./collision [target_path] [source_path]")
    return
  }

  target, err := ioutil.ReadFile(args[1])
  if err != nil {
    fmt.Println("There was an issue reading the target file")
    return 
  }

  source, err := ioutil.ReadFile(args[2])
  if err != nil {
    fmt.Println("There was an issue reading the source file.")
    return
  }

  target_abi := utils.DecodeTruffleABI(target).Abi
  if target_abi == nil {
    fmt.Println("The target file did not encode a non-empty contract ABI.")
    return 
  } 

  source_abi := utils.DecodeTruffleABI(source).Abi
  if source_abi == nil {
    fmt.Println("The source file did not encode a non-empty contract ABI.")
    return 
  }

  for _, abi := range(target_abi) {
    result := findCollisions(abi, source_abi)
    // It is often useful to redirect the output of programs into files or filters like sed and 
    // grep when scripting. This is the reason that these results are printed on seperate lines.
    for _, collision := range(result) {
      fmt.Println(collision) 
    }
  }
}

func findCollisions(target utils.ABI, source []utils.ABI) []utils.ABI {
  var result []utils.ABI

  if target.Type != "function" {
    return result
  }

  target_selector := utils.FunctionSelectorABI(target)
  for _, function := range(source) {
    if target_selector == utils.FunctionSelectorABI(function) && function.Type == "function" {
      result = append(result, function) 
    }
  }
  return result
}
