package core

import (
	"fmt"
	"github.com/monigo/srun/model"
	"github.com/monigo/srun/store"
	"reflect"
	"testing"
)

func TestLogin(t *testing.T) {
	acc, err := store.LoadAccount()
	fmt.Println(Login())
}
