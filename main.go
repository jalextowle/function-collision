package main

import (
  "fmt"
  "os"
  "./utils"
)

func main() {
  args := os.Args
  if len(args) != 3 {
    fmt.Println("Usage: ./collision [selector] [abi_path]")
    return
  }

  utils.PrintHash(utils.Keccak256("abc"))
}
