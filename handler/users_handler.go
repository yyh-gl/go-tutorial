// Filename: handler/users_handler.go

package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

// Model共通分を表す構造体
type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Userを表す構造体
type User struct {
	Model
	Name string
	Age  int
}

// User作成用のパラメータを表す構造体
type UserParams struct {
	Name string
	Age  int
}

func UsersShowHandler(w http.ResponseWriter, r *http.Request) {
	// DBに接続
	db := gormConnect()

	vars := mux.Vars(r)

	// 指定IDのユーザを取得
	var user User
	db.First(&user, vars["id"])

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Show this user -> %v", user.Name)
}

func UsersCreateHandler(w http.ResponseWriter, r *http.Request) {
	// DBに接続
	db := gormConnect()
	defer db.Close() // dbを使い終わったらクローズ

	// リクエストBodyをJSONにパース
	decoder := json.NewDecoder(r.Body)
	var userParams UserParams
	error := decoder.Decode(&userParams) // userParamsとbody内の対応するキーの値を入れてくれる
	if error != nil {
		w.Write([]byte("json decode error" + error.Error() + "\n"))
	}
	defer r.Body.Close() // bodyを使い終わったらクローズ

	//INSERT実行部分
	var user User
	user.Name = userParams.Name
	user.Age = userParams.Age
	db.Create(&user)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Create this user -> %v", user.Name)
}

func gormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := "mysql"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME   := "go_api_db"

	db,err := gorm.Open(DBMS, USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME)

	if err != nil {
		panic(err.Error())
	}
	return db
}
