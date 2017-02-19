package main

import (
	"os"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/eugene0707/gettek/routers"

	"github.com/astaxie/beego"
	"log"
	"github.com/eugene0707/gettek/models"
	"github.com/eugene0707/gettek/common"
)

func main() {
	common.CurrentDB().AutoMigrate(&models.Driver{}, &models.Metric{})

	port := os.Getenv("PORT");
	if port == "" {
		port = beego.AppConfig.String("httpport")
	}

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/api/v1/doc"] = "swagger"

	beego.Run(":" + port)
}
