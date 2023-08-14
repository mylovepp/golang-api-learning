package services

import (
	"context"
	"mylovepp/ent"
	"mylovepp/repositories"
)

type UserService struct{}

func (UserService) GetUsers() []*ent.User {
	users, err := repositories.GetUsers(context.Background())
	if err != nil {
		println(err.Error())
		return nil
	} else {
		return users
	}
}

func (UserService) GetUser(id int) *ent.User {
	user, err := repositories.GetUser(context.Background(), id)

	if err != nil {
		println(err.Error())
		return nil
	} else {
		return user
	}
}

func (UserService) AddUser(user *ent.User) {
	result, err := repositories.AddUser(context.Background(), user)
	if err != nil {
		println(err.Error())
	} else {
		user.ID = result.ID
	}
}

func (UserService) AddUserWithAccounts(user *ent.User, accounts []*ent.Account) error {
	result, resultAccounts, err := repositories.AddUserWithAccounts(context.Background(), user, accounts)
	if err != nil {
		println(err.Error())
		return err
	} else {
		user.ID = result.ID

		for index, account := range resultAccounts {
			accounts[index] = account
		}
	}

	return nil
}

func (UserService) UpdateUser(user *ent.User) bool {
	user, err := repositories.UpdateUser(context.Background(), user)
	if err != nil {
		println(err.Error())
		return false
	} else {
		println(user)
		return true
	}
}

func (UserService) DeleteUser(id int) bool {
	err := repositories.DeleteUser(context.Background(), id)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return true
	}
}
