package entities

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/tnnz20/godemy-be/pkg/errs"
	"github.com/tnnz20/godemy-be/pkg/jwt"
)

const (
	ROLE_Teacher string = "teacher"
	ROLE_Student string = "student"
)

type Users struct {
	ID         uuid.UUID
	Email      string
	Password   string
	Name       string
	Date       time.Time
	Address    string
	Gender     string
	ProfileImg string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

func NewUsersRegister(email, password, name string) Users {
	return Users{
		ID:        uuid.New(),
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewUsersUpdate(date time.Time, name, address, gender, profileImg string) Users {
	return Users{
		Name:       name,
		Date:       date,
		Address:    address,
		Gender:     gender,
		ProfileImg: profileImg,
		UpdatedAt:  time.Now(),
	}
}

func NewUsersLogin(email, password string) Users {
	return Users{
		Email:    email,
		Password: password,
	}
}

// IsEmailAlreadyExists is a method to check if the email is already exists in the database
func (u Users) IsEmailAlreadyExists() bool {
	return u.ID != uuid.Nil
}

// ValidateAuth is a method to validate the user authentication payload
func (u Users) ValidateAuth() (err error) {
	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if err := u.ValidatePassword(); err != nil {
		return err
	}

	if err := u.ValidateName(); err != nil {
		return err
	}

	return
}

// ValidateCreateUsers is a method to validate the user update payload
func (u Users) ValidateUpdateUsers() (err error) {

	if err := u.ValidateName(); err != nil {
		return err
	}

	if err := u.ValidateAddress(); err != nil {
		return err
	}

	if err := u.ValidateGender(); err != nil {
		return err
	}

	return
}

func (u Users) ValidateEmail() (err error) {
	if u.Email == "" {
		return errs.ErrEmailRequired
	}

	splitEmail := strings.Split(u.Email, "@")
	if len(splitEmail) != 2 {
		return errs.ErrInvalidEmail
	}

	return
}

func (u Users) ValidatePassword() (err error) {
	if u.Password == "" {
		return errs.ErrPasswordRequired
	}

	if len(u.Password) < 8 {
		return errs.ErrInvalidLengthPassword
	}

	return
}

func (u Users) ValidateName() (err error) {
	if u.Name == "" {
		return errs.ErrNameRequired
	}

	if len(u.Name) < 3 {
		return errs.ErrInvalidLengthName
	}

	return
}

func (u Users) ValidateAddress() (err error) {
	if len(u.Address) < 5 {
		return errs.ErrInvalidLengthAddress
	}

	return
}

func (u Users) ValidateGender() (err error) {
	validGender := map[string]bool{
		"male":   true,
		"female": true,
	}

	if _, ok := validGender[u.Gender]; !ok {
		return errs.ErrInvalidGender
	}

	return
}

// HashingPassword is a method to hash the password
func (u *Users) HashingPassword() (err error) {
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(encryptedPass)
	return nil
}

func (u Users) VerifyPasswordFromEncrypted(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
}

func (u Users) VerifyPasswordFromPlain(encrypted string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(u.Password))
}

func (u Users) GenerateToken(role, secret string) (tokenString string, err error) {
	return jwt.GenerateToken(u.ID.String(), role, secret)
}
