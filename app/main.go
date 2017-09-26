package main

import (
	"github.com/SekiguchiKai/try_gin_fw_go/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// エントリポイント
func init() {
	// デフォルトのミドルウェアを元にrouterを生成
	http.Handle("/", getMainEngine())
}

// routerを設定
func getMainEngine() *gin.Engine {
	router := gin.Default()

	const USER = "/user"
	router.POST(USER, api.CreateUser)
	router.GET(USER + "/:id", api.GetUser)

	return router
}
