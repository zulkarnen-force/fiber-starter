// usecase/user_usecase.go
package usecase

import (
	"errors"
	"fmt"
	"time"

	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/zulkarnen-force/fiber-starter/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo       domain.UserRepository
	jwtSecret  string
	tokenExpiry time.Duration
}

func NewUserUsecase(repo domain.UserRepository, jwtSecret string, tokenExpiry time.Duration) domain.UserUsecase {
	return &userUsecase{
		repo:       repo,
		jwtSecret:  jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

func (u *userUsecase) Register(user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return u.repo.Create(user)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.GetByEmail(email)
	if err != nil {
		return "", err
	}

	fmt.Println(user.Password);
	fmt.Println(password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, jtoken.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(u.tokenExpiry).Unix(),
	})
	return token.SignedString([]byte(u.jwtSecret))
}

func (u *userUsecase) GetUser(id uint) (*domain.User, error) {
	return u.repo.GetByID(id)
}

func (u *userUsecase) UpdateUser(user *domain.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.repo.Delete(id)
}
