package middleware

import "golang.org/x/crypto/bcrypt"

func ValidPassword(originPwd string, pwd string, salt string) bool {
	done := pwd + salt
	err := bcrypt.CompareHashAndPassword([]byte(originPwd), []byte(done))
	if err != nil {
		return false
	}
	return true
}

func BuildPassSalt(pwd string, salt string) (string, error) {
	done := pwd + salt
	hash, err := bcrypt.GenerateFromPassword([]byte(done), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}