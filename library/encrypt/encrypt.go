package encrypt

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password *string) (err error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.MinCost)
	if err != nil {
		return
	}
	*password = string(encryptedPassword)
	return
}