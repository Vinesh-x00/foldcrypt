package utiles

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func AskPasswd() string {
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		return "$ERR$"
	}
	return string(bytePassword)
}

func GetHash(key string) string {
	hash := sha512.New()
	hash.Write([]byte(key))
	hexv := hash.Sum(nil)
	return hex.EncodeToString(hexv)

}
