package seed

import (
	users "github.com/d1agnoze/kitchen/services/auth/models"
	faker "github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Options struct {
	Count int
}

func CreateUserSeeds(db *gorm.DB, opt Options) error {
	for i := 0; i < opt.Count; i++ {
		x := users.User{
			Name:         faker.Name(),
			Email:        faker.Email(),
			PasswordHash: generateHash("123456"),
			Address:      faker.GetRealAddress().Address,
		}

		err := db.Save(&x).Error

		if err != nil {
			return err
		}
	}

	return nil
}

func generateHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return faker.Password()
	}
	return string(bytes)
}
