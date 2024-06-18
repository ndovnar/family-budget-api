package dto

import (
	"github.com/ndovnar/family-budget-api/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Balance   float64            `bson:"balance"`
	IsDeleted bool               `bson:"deleted"`
	Dates     Dates              `bson:"dates,omitempty"`
}

func ModelAccountToDtoAccount(account *model.Account) *Account {
	return &Account{
		ID:        primitive.NewObjectID(),
		Name:      account.Name,
		Balance:   account.Balance,
		IsDeleted: account.IsDeleted,
		Dates:     ModelDatesToDtoDates(account.Dates),
	}
}

func DtoAccountToModelAccount(account *Account) *model.Account {
	return &model.Account{
		ID:        account.ID.Hex(),
		Name:      account.Name,
		Balance:   account.Balance,
		IsDeleted: account.IsDeleted,
		Dates:     DtoDatesToModelDates(account.Dates),
	}
}
