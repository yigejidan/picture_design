package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"picture_design/common"
	"picture_design/routes"
)

func init() {
	err := common.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err.Error())
	}
}

func main() {
	db := common.InitDB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database,err: %v", err.Error())
	}
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatalf("failed to close database,err: %v", err.Error())
		}
	}(sqlDB)
	gin.SetMode(common.SvrConfig.RunMode)
	gin.DefaultWriter = common.LogWriter()
	r := gin.Default()
	//r.Use(middleware.AuthMiddleware())
	r = routes.UserRoute(r)
	r = routes.PictureRoute(r)
	if err := r.Run(":" + common.SvrConfig.HttpPort); err != nil {
		log.Fatalf("failed to run,err: %v", err.Error())
	}
}
