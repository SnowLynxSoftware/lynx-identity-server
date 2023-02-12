package util

import (
	"SnowLynxSoftware/lynx-identity-server/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(config.AppConfig.AppKey+password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(config.AppConfig.AppKey+password))
	return err == nil
}
