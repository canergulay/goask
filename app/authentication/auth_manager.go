package authentication

import (
	"bufio"
	"fmt"
	"os"

	authentication "github.com/canergulay/goask/app/authentication/auth_entity"
)

const (
	EKSI    = "eksi"
	TWITTER = "twitter"
)

func NewAuthEntity(targetedPlatform string) authentication.AuthEntity {
	reader := bufio.NewReader(os.Stdin)
	username := readCredentialInput(reader, targetedPlatform, "username")
	password := readCredentialInput(reader, targetedPlatform, "password")
	authEntity := CreateAuthEntity(targetedPlatform, username, password)
	return authEntity
}

func CreateAuthEntity(platform string, username string, password string) authentication.AuthEntity {
	switch platform {
	case EKSI:
		eksiAuthEntity := &authentication.EksiAuthEntity{}
		eksiAuthEntity.InitialiseEntity(username, password)
		return eksiAuthEntity
	case TWITTER:
		os.Exit(1)
		return nil
	}
	return nil
}

func readCredentialInput(reader *bufio.Reader, targetedPlatform string, targetedVariable string) string {
	fmt.Printf("please enter your %s for %s : \n", targetedVariable, targetedPlatform)
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("please enter a valid %s \n", targetedVariable)
		os.Exit(1)
	}
	return value
}
