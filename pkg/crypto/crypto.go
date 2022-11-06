package crypto

import (
	"fmt"
	"math/rand"
)

func GenSalt() (string, error) {
	buff := make([]byte, 16)
	_, err := rand.Read(buff)
	return fmt.Sprintf("%x", buff), err
}
