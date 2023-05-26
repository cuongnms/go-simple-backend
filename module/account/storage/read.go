package storage

import (
	"context"
	"go-simple-backend/common"
	"go-simple-backend/module/account/model"
)

func (s *sqlStorage) ReadAccount(ctx context.Context, conditions map[string]interface{}) (*model.Account, error) {
	var data model.Account
	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		return nil, common.NewCustomError(err, "")
	}
	return &data, nil
}
