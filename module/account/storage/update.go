package storage

import (
	"context"
	"go-simple-backend/common"
	"go-simple-backend/module/account/model"
)

func (s *sqlStorage) UpdateAccount(ctx context.Context, data *model.Account) error {
	if err := s.db.Updates(data).Error; err != nil {
		return common.NewCustomError(err, "")
	}
	return nil
}
