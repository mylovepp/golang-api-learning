package repositories

import (
	"context"
	"mylovepp/ent"
	"mylovepp/ent/account"
)

func GetAccounts(ctx context.Context) ([]*ent.Account, error) {
	return entClient.Account.Query().All(ctx)
}

func GetAccountsByUserId(ctx context.Context, userId int) ([]*ent.Account, error) {
	return entClient.Account.Query().Where(account.UserIDEQ(userId)).All(ctx)
}

func GetAccount(ctx context.Context, id int) (*ent.Account, error) {
	return entClient.Account.Query().Where(account.IDEQ(id)).First(ctx)
}

func AddAccounts(ctx context.Context, accounts []*ent.Account) ([]*ent.Account, error) {
	bulk := []*ent.AccountCreate{}
	for _, account := range accounts {
		bulk = append(bulk, entClient.Account.Create().
			SetAssetID(account.AssetID).
			SetUserID(account.UserID).
			SetAccountNo(account.AccountNo))
	}
	return entClient.Account.CreateBulk(bulk...).Save(ctx)
}
