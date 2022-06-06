package token

import (
	"fmt"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	a, _ := GetToken("password", time.Now().Unix(), 1, 1)
	time.Sleep(time.Second * 5)
	u, err := ParseToken(a, "password")
	fmt.Println(err)
	fmt.Println(u)
}
