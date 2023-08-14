package repositories

import (
	"context"
	"mylovepp/ent"
	"mylovepp/ent/user"
)

func GetUsers(ctx context.Context) ([]*ent.User, error) {
	return entClient.User.Query().All(ctx)
}

func GetUser(ctx context.Context, id int) (*ent.User, error) {
	return entClient.User.Query().Where(user.IDEQ(id)).First(ctx)
}

func AddUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return entClient.User.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetBirthDate(user.BirthDate).
		Save(ctx)
}

func AddUserWithAccounts(ctx context.Context, user *ent.User, accounts []*ent.Account) (*ent.User, []*ent.Account, error) {
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if v := recover(); v != nil {
			println("rollback transaction")
			tx.Rollback()
			panic(v)
		}
	}()
	newUser, userErr := tx.User.Create().
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetBirthDate(user.BirthDate).
		Save(ctx)
	if userErr != nil {
		tx.Rollback()
		return nil, nil, userErr
	}

	bulk := []*ent.AccountCreate{}
	for _, account := range accounts {
		bulk = append(bulk, tx.Account.Create().
			SetAssetID(account.AssetID).
			SetUserID(newUser.ID).
			SetAccountNo(account.AccountNo))
	}
	newAccounts, accountError := tx.Account.CreateBulk(bulk...).Save(ctx)
	if accountError != nil {
		tx.Rollback()
		return nil, nil, accountError
	}
	tx.Commit()

	return newUser, newAccounts, nil
}

func UpdateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return entClient.User.UpdateOneID(user.ID).
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetBirthDate(user.BirthDate).
		Save(ctx)
}

func DeleteUser(ctx context.Context, id int) error {
	return entClient.User.DeleteOneID(id).Exec(ctx)
}
