package model

import (
	"context"
	"github.com/SekiguchiKai/try_gin_fw_go/util"
	"github.com/mjibson/goon"
	"github.com/satori/go.uuid"
	"time"
)

// ユーザー
type User struct {
	ID        string    `json:"id" goon:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
}

// 新規のUserインスタンスの生成
func NewUser(name string) *User {
	u := new(User)
	u.ID = uuid.NewV4().String()
	u.Name = name
	u.CreatedAt = time.Now()

	return u
}

// DatastoreにUserを登録する
func (u *User) Post(c context.Context) (*User, error) {
	g := goon.FromContext(c)
	// DatastoreにUserのEntityを登録
	if _, err := g.Put(u); err != nil {
		util.ErrorLog(c, err.Error())
		return u, nil
	}

	return u, nil
}

//  DatastoreからUserを取得する
func (u *User) Get(c context.Context) (*User, error) {
	g := goon.FromContext(c)
	// key(userID)を指定してDatastoreからEntityを取得する、エラーの場合のハンドリングを同時に行う
	if err := g.Get(u); err != nil {
		util.ErrorLog(c, err.Error())
		return nil, err
	}

	return u, nil
}
