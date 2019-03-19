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
    fmt.Println("Usage: ./collision [selector] [abi_path]")
    return
  }

  contents, err := ioutil.ReadFile(args[2])
  if err != nil {
    panic(err)
  }

  fmt.Println(contents)
}

func createCollision() {
  // Just a placeholder to disable the warning of not importing utils
  a := utils.FunctionSelector("abc")
  fmt.Println(a)
}
