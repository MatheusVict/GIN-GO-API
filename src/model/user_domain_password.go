package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (user *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(user.password))
	user.password = hex.EncodeToString(hash.Sum(nil))
}
