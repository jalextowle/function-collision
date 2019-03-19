package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "./utils"
//  "reflect"
)

func main() {
  // TODO: Refactor this to allow a variable number of arguments, or a config file
  //       to list all reasonable files
  args := os.Args
  if len(args) != 3 {
    fmt.Println("Usage: ./collision [target_path] [source_path]")
    return
  }

  target, err := ioutil.ReadFile(args[1])
  if err != nil {
    fmt.Println("Error: Issue reading the target file")
    return 
  }

  source, err := ioutil.ReadFile(args[2])
  if err != nil {
    fmt.Println("Error: Issue reading the source file")
    return
  }

  target_abi := utils.DecodeTruffleABI(target).Abi
  source_abi := utils.DecodeTruffleABI(source).Abi

  fmt.Println(target_abi)
  fmt.Println(source_abi)
}
