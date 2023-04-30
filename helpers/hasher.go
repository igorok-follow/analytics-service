package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (m *Hasher) HashBCrypt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (m *Hasher) CheckBCrypt(new string, old string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(old), []byte(new))
	if err != nil {
		return false, err
	}

	return true, nil
}
