package main

import (
	"fmt"
	"go-todoapp/db"
	"go-todoapp/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migreated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
