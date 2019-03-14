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
}

func createCollision() {
  // Just a placeholder to disable the warning of not importing utils
  a := utils.Selector("abc")
  fmt.Println(a)
}
