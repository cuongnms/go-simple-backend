package model

import (
	"go-simple-backend/common"
	"strings"
	"time"
)

type Account struct {
	Id        int8       `json:"id" gorm:"column:id"`
	Owner     string     `json:"owner" gorm:"column:owner"`
	Balance   int8       `json:"balance" gorm:"column:balance"`
	Currency  string     `json:"currency" gorm:"column:currency"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"created_at"`
}

func (Account) TableName() string {
	return "accounts"
}

func (data *Account) Validate() error {
	if strings.TrimSpace(data.Currency) == "" || strings.TrimSpace(data.Owner) == "" {
		return common.NewCustomError(nil, "Valid data")
	}
	return nil
}
