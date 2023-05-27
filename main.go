package main

import (
	"github.com/gin-gonic/gin"
	"go-simple-backend/common"
	"go-simple-backend/module/account/transport"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dns := os.Getenv("MYSQL_CONN_STRING")
	var db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db = db.Debug()

	appCtx := common.NewAppCtx(db)
	r := gin.Default()
	r.GET("", transport.CreateAccount(appCtx))

	r.Run()
}
