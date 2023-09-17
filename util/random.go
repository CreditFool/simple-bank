package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmonpqrstuvwxyz"

func RandomInt(min int64, max int64) int64 {
  return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
  var output strings.Builder
  k := len(alphabet)

  for i := 0; i < n; i++ {
    c := alphabet[rand.Intn(k)]
    output.WriteByte(c)
  }

  return output.String()
}

func RandomOwner() string {
  return RandomString(6)
}

func RandomMoney() int64 {
  return RandomInt(0, 1000000)
}


func RandomCurrency() string {
  currencies := []string{"USD", "IDR", "JPY"}
  n := len(currencies)
  return currencies[rand.Intn(n)]
}
