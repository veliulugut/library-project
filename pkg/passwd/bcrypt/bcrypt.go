package bcrypt

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"library/pkg/passwd"
)

var _ passwd.Interface = (*Bcrypt)(nil)

type Bcrypt struct {
	secret string
	cost   int
}

func NewBcrypt(secret string, cost int) *Bcrypt {
	return &Bcrypt{
		secret: secret,
		cost:   cost,
	}
}

func (b *Bcrypt) Generate(password string) (string, error) {
	var (
		generated []byte
		err       error
	)

	if generated, err = bcrypt.GenerateFromPassword([]byte(b.secret+password+b.secret), b.cost); err != nil {
		return "", fmt.Errorf("bcrypt / generate :%w", err)
	}

	return string(generated), nil

}

func (b *Bcrypt) Compare(hashed, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(b.secret+password+b.secret)); err != nil {
		return fmt.Errorf("bcrypt / compare :%w", err)
	}

	return nil
}
