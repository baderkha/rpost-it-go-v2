package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	HashPassword(passowrd string) string
	CompareHash(userInput string, hashedPassword string) bool
}

type Bcrypt struct {
	Rounds uint
}

func (b *Bcrypt) HashPassword(password string) string {
	cost := bcrypt.DefaultCost
	if b.Rounds != 0 {
		cost = int(b.Rounds)
	}
	hashByte, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashByte)
}

func (b *Bcrypt) CompareHash(userInput string, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userInput)) == nil
}
