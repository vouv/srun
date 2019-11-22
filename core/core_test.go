package core

import (
	"fmt"
	"github.com/vouv/srun/store"
	"testing"
)

func TestLogin(t *testing.T) {
	acc, _ := store.LoadAccount()
	fmt.Println(Login(acc.Username, acc.Password))
}
