package storage

import (
	"context"
	"go-simple-backend/module/account/model"
)

func (s *sqlStorage) CreateAccount(ctx context.Context, data *model.Account) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
