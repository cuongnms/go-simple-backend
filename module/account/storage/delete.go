package storage

import (
	"context"
	"go-simple-backend/common"
	"go-simple-backend/module/account/model"
)

func (s *sqlStorage) DeleteAccount(ctx context.Context, id int8, conditions map[string]interface{}) error {
	if err := s.db.Where(conditions).Delete(&model.Account{}, id).Error; err != nil {
		return common.NewCustomError(err, "")
	}
	return nil
}
