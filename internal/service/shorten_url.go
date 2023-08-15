package service

import (
	"crypto/sha256"
	"github.com/mr-tron/base58"
)



func ShortenURL(url string) string {  
  hash := sha256.Sum256([]byte(url))
  encoded := base58.Encode(hash[:])

  return encoded[:8]
}
