package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateAccessCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06d", r.Intn(1000000))
}
