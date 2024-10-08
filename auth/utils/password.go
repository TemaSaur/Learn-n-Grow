package auth

import (
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) ([]byte, error) {
	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}
