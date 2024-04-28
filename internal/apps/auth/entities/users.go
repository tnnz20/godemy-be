package entities

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/jwt"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// IsEmailAlreadyExists is a method to check if the email is already exists in the database
func (u User) IsEmailAlreadyExists() bool {
	return u.ID != uuid.Nil
}

// Validate is a method to validate the user
func (u User) Validate() (err error) {
	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	if err := u.ValidateRole(); err != nil {
		return err
	}

	return
}

func (u User) ValidateEmail() (err error) {
	if u.Email == "" {
		return errs.ErrEmailRequired
	}

	splitEmail := strings.Split(u.Email, "@")
	if len(splitEmail) != 2 {
		return errs.ErrInvalidEmail
	}

	return
}

func (u User) ValidatePassword() (err error) {
	if u.Password == "" {
		return errs.ErrPasswordRequired
	}

	if len(u.Password) < 8 {
		return errs.ErrInvalidLengthPassword
	}

	return
}

func (u User) ValidateRole() (err error) {
	if u.Role == "" {
		return errs.ErrRoleRequired
	}

	if u.Role != "student" && u.Role != "teacher" {
		return errs.ErrInvalidRole
	}

	return
}

// HashingPassword is a method to hash the password
func (u *User) HashingPassword() (err error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(encryptedPass)
	return nil
}

func (u User) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

func (u User) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(u.Password))
}

func (u User) GenerateToken(secret string) (tokenString string, err error) {
	return jwt.GenerateToken(u.ID.String(), u.Role, secret)
}
