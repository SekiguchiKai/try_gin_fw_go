package api

import (
	"encoding/json"
	"github.com/SekiguchiKai/try_gin_fw_go/model"
	"github.com/SekiguchiKai/try_gin_fw_go/util"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"io/ioutil"
	"log"
	"net/http"
)

type UserName struct {
	Name string `json:"name"`
}

// ユーザーを新規生成
// *gin.Contextは、現在のHTTP requestのcontextを表す。
// *gin.Contextは、request objects、response objects、 path、 path parameters、 data、registered handlerを持っている。
func CreateUser(c *gin.Context) {
	log.Println("CreateUser is called.")

	ac := appengine.NewContext(c.Request)

	// request bodyからJSONを読み込む
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		util.ErrorLog(ac, err.Error())
	}
	defer c.Request.Body.Close()

	util.DebugLog(ac, string(body))

	un := &UserName{}
	// JSONをUnmarshal
	if err := json.Unmarshal(body, un); err != nil {
		util.ErrorLog(ac, err.Error())
	}

	// Userのインスタンスを生成
	u := model.NewUser(un.Name)
	// DatastoreにUserを登録
	if _, err := u.Post(ac); err != nil {
		util.ErrorLog(ac, err.Error())
	}
	util.DebugLog(ac, u.Name)

	// Bindは、リクエストボディを与えられた引数と結びつける。
	//if err := c.Bind(u); err != nil {
	//	ErrorLog(ac, err.Error())
	//	return err
	//}
	// 引数で与えられたstatus codeと共に、構造体をJSONにして返す
	c.JSON(http.StatusCreated, u)
}

// ユーザーを取得
func GetUser(c *gin.Context) {
	log.Println("GetUser is called.")
	ac := appengine.NewContext(c.Request)
	// DatastoreからUserをGetする際に、データを格納するためのインスタンスを生成
	u := new(model.User)
	// クエリパラメータからidを取得し、UserのIDに格納
	u.ID = c.Param("id")
	util.DebugLog(ac, u.ID)

	// DatastoreからUserのデータを取得
	u, err := u.Get(ac)
	if err != nil {
		util.ErrorLog(c.Request.Context(), err.Error())
	}
	// 引数で与えられたstatus codeと共に、構造体をJSONにして返す
	c.JSON(http.StatusOK, u)

}
