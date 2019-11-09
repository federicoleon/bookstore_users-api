package services

import (
	"github.com/federicoleon/bookstore_users-api/domain/users"
	"github.com/federicoleon/bookstore_users-api/utils/errors"
	"github.com/federicoleon/bookstore_users-api/utils/date_utils"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	dao := &users.User{Id: userId}
	if err := dao.Get(); err != nil {
		return nil, err
	}
	return dao, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
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

func DeleteUser(userId int64) *errors.RestErr {
	dao := &users.User{Id: userId}
	return dao.Delete()
}

func Search(status string) ([]users.User , *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
