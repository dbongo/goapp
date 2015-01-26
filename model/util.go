package model

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// CreateGravatar ...
func CreateGravatar(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	hash := md5.New()
	hash.Write([]byte(email))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
