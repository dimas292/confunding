package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput)(User, error)
	Login(input LoginUserInput)(User, error)
	IsEmailAvalable(input CheckUserInput)(bool, error)
}
type service struct {
	repository Repository
}
func NewService(repository Repository) *service{
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error){
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	password_hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(password_hash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginUserInput)(User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func(s *service) IsEmailAvalable(input CheckUserInput)(bool, error){
	email := input.Email

	checkUser, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if checkUser.ID == 0{
		return true, nil 
		
	}

	return false, nil 
}
