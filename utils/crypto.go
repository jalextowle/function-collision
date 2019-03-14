package utils 

import (
  "encoding/hex"
  "fmt"
  "github.com/ethereum/go-ethereum/crypto"
)

func Keccak256(data string) []byte {
  return crypto.Keccak256([]byte(data))
}

// TODO: Delete if unneeded
func PrintHash(hash []byte) {
  fmt.Println("0x" + hex.EncodeToString(hash))
}
