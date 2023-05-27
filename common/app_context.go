package common

import "gorm.io/gorm"

type AppContext interface {
	GetDatabaseConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppCtx(db *gorm.DB) *appCtx {
	return &appCtx{
		db: db,
	}
}

func (a *appCtx) GetDatabaseConnection() *gorm.DB {
	return a.db
}
