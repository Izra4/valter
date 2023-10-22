package main

import (
	db2 "Valter/db"
	"Valter/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := db2.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.Default()
	handler.StartEngine(r, db)
	r.Run()
}
