package main

import (
  "encoding/hex"
  "fmt"
  "os"
  "github.com/ethereum/go-ethereum/crypto"
)

func main() {
  args := os.Args
  if len(args) != 3 {
    fmt.Println("Usage: ./collision [selector] [abi_path]")
    return
  }

  printHash(keccak256("abc"))
}

func keccak256(data string) []byte {
  return crypto.Keccak256([]byte(data))
}

func printHash(hash []byte) {
  fmt.Println("0x" + hex.EncodeToString(hash))
}
