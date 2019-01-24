package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"time"
)

type Model struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// Userを表す構造体
type User struct {
	Model
	Name string
	Age  int
}

var user User

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// DBに接続
	db := gormConnect()

	vars := mux.Vars(r)
	// 指定した条件を元に複数のレコードを引っ張ってくる
	selectedUser := db.First(&user, vars["id"])
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User No.%v", selectedUser)
}

func gormConnect() *gorm.DB {
	DBMS     := "mysql"
	USER     := "root"
	PASS     := ""
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME   := "go_api_db"

	CONNECT = USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}