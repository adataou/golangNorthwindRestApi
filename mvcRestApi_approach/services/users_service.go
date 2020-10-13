package services

import (
	"golangNorthwindRestApi/mvcRestApi_approach/domain/users"
	"golangNorthwindRestApi/utils/crypto_utils"
	"golangNorthwindRestApi/utils/date_utils"
	"golangNorthwindRestApi/utils/rest_errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, rest_errors.RestErr)
	GetUser(userId int64) (*users.User, rest_errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr)
	DeleteUser(userId int64) rest_errors.RestErr
	SearchUser(status string) (users.Users, rest_errors.RestErr)
	LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr)
}

type usersService struct{}

func (s *usersService) CreateUser(user users.User) (*users.User, rest_errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password)
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, rest_errors.RestErr) {
	dao := &users.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, rest_errors.RestErr) {
	current := &users.User{Id: user.Id}
	if err := current.Get(); err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userId int64) rest_errors.RestErr {
	user := users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, rest_errors.RestErr) {
	user := &users.User{}
	return user.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
