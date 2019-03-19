package utils 

import (
  "encoding/hex"
  "github.com/ethereum/go-ethereum/crypto"
  "strings"
)

func FunctionSelector(args ...string) string {
  params := "(" + strings.Join(args[1:], ",") + ")"
  return "0x" + hex.EncodeToString(crypto.Keccak256([]byte(args[0] + params))[:4])
}


func FunctionSelectorABI(functionABI ABI) string {
  content := functionABI.Name + "("
  for idx, input := range(functionABI.Inputs) {
    content += input.Type
    if idx < len(functionABI.Inputs) - 1 {
      content += "," 
    }
  }
  truncated_hash := crypto.Keccak256([]byte(content + ")"))[:4]
  return "0x" + hex.EncodeToString(truncated_hash)
}
