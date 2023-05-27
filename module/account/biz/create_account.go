package biz

import (
	"context"
	"go-simple-backend/common"
	"go-simple-backend/module/account/model"
)

type CreateAccountStorage interface {
	CreateAccount(ctx context.Context, data *model.Account) error
}
type CreateAccountBiz struct {
	storage CreateAccountStorage
}

func NewCreateAccountBiz(storage CreateAccountStorage) *CreateAccountBiz {
	return &CreateAccountBiz{storage: storage}
}

func (a *CreateAccountBiz) CreateAccount(ctx context.Context, data *model.Account) error {
	if err := data.Validate(); err != nil {
		return common.NewCustomError(err, "")
	}
	if err := a.storage.CreateAccount(ctx, data); err != nil {
		return common.NewCustomError(err, "")
	}
	return nil
}
