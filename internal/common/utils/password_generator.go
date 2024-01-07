package utils

import "golang.org/x/crypto/bcrypt"

func NormalizePassword(p string) []byte {
	return []byte(p)
}

// 비밀번호 해시 생성
func GeneratePassword(p string) string {
	bytePw := NormalizePassword(p)

	hash, err := bcrypt.GenerateFromPassword(bytePw, bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	return string(hash)
}

// 비밀번호 비교
func ComparePassword(hashedPwd string, inputPwd string) bool {
	byteHash := NormalizePassword(hashedPwd)
	byteInput := NormalizePassword(inputPwd)
	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
