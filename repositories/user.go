package repositories

import (
	"github.com/johnrazeur/gin-boilerplate/models"
	"golang.org/x/crypto/bcrypt"
)

// UserRepository the user repository
type UserRepository struct {
	Repository
}

// Create the user
func (r *UserRepository) Create(user *models.User) error {
	_, err := r.GetFirstByEmail(user.Email)

	if err == nil {
		return ErrUserExists
	} else if err != ErrDataNotFound {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// replace the plaintext password with ciphertext password
	user.Password = string(hash)

	db := r.Db.Create(user)

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return ErrKeyConflict
	}

	return nil
}

// Login the user if the password is correct
func (r *UserRepository) Login(email, password string) (*models.User, error) {
	user, err := r.GetFirstByEmail(email)

	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, ErrWrongPassword
	}

	return user, nil
}

// GetFirstByEmail gets the user by his email
func (r *UserRepository) GetFirstByEmail(email string) (*models.User, error) {
	user := &models.User{}
	db := r.Db.Where("email=?", email).First(user)
	if db.RecordNotFound() {
		return user, ErrDataNotFound
	} else if db.Error != nil {
		return user, db.Error
	}

	return user, nil
}
