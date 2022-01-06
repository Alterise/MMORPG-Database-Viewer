package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type responseID int

const (
	AUTH_SUCCESS responseID = iota
	AUTH_FAILED
	AUTH_ADMIN
)

func verifyUser(login string, password string) responseID {
	hash := getMD5Hash(password)
	serverMsg := sendCommandToServer(fmt.Sprintf("verify %s %s", login, hash))
	if serverMsg == "success\n" {
		if login == "admin" {
			return AUTH_ADMIN
		} else {
			return AUTH_SUCCESS
		}
	} else {
		return AUTH_FAILED
	}

}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}