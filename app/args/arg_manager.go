package args

import (
	"os"
)

type ArgManager struct {
	First  string
	Second string
}

func NewArgManager() *ArgManager {
	args := os.Args[1:]
	return &ArgManager{First: args[0], Second: args[1]}
}
