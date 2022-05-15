package args

import (
	"fmt"

	"github.com/canergulay/goask/app/authentication"
)

const (
	auth = "auth"
)

func (arger *ArgManager) Initialise() {
	switch arger.First {
	case auth:
		authManager := authentication.NewAuthManager(arger.Second)
		fmt.Println(authManager) //
	}

}
