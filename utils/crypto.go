package utils 

import (
  "encoding/hex"
  "github.com/ethereum/go-ethereum/crypto"
  "strings"
)

func Selector(args ...string) string {
  params := "(" + strings.Join(args[1:], ",") + ")"
  return "0x" + hex.EncodeToString(crypto.Keccak256([]byte(args[0] + params))[:4])
}
