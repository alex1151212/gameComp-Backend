package utils

import bcrypt "golang.org/x/crypto/bcrypt"

func Encode(data string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func Compare(hashPassword string, loginPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(loginPassword))
	return err == nil
}
