package authentication

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EKSI    = "eksi"
	TWITTER = "twitter"
)

type AuthManager struct {
	username string
	password string
}

func NewAuthManager(targetedPlatform string) AuthManager {
	reader := bufio.NewReader(os.Stdin)
	username := readCredentialInput(reader, targetedPlatform, "username")
	password := readCredentialInput(reader, targetedPlatform, "password")
	return AuthManager{username: username, password: password}
}

func readCredentialInput(reader *bufio.Reader, targetedPlatform string, targetedVariable string) string {
	fmt.Printf("please enter your username for %s : \n", targetedPlatform)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("please enter a valid %s \n", targetedVariable)
		os.Exit(1)
	}
	return value
}
