package services

import (
	"context"
	"mylovepp/ent"
	"mylovepp/repositories"
)

func GetAccountsByUserId(userId int) ([]*ent.Account, error) {
	results, err := repositories.GetAccountsByUserId(context.Background(), userId)
	if err != nil {
		println(err.Error())
		return nil, err
	} else {
		return results, nil
	}
}

func AddAccounts(user *ent.User, trans context.Context) ([]*ent.Account, error) {
	accounts := []*ent.Account{}
	accounts = append(accounts, &ent.Account{AssetID: 1, UserID: user.ID, AccountNo: "1234567890"})
	accounts = append(accounts, &ent.Account{AssetID: 2, UserID: user.ID, AccountNo: "1234567891"})
	accounts = append(accounts, &ent.Account{AssetID: 3, UserID: user.ID, AccountNo: "1234567892"})

	results, err := repositories.AddAccounts(trans, accounts)
	if err != nil {
		println(err.Error())
		return nil, err
	} else {
		return results, nil
	}
}
