package main

import (
	db2 "Valter/db"
	"Valter/handler"
	"Valter/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
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
	test := utility.SupabaseClient()
	r := gin.Default()
	r.GET("/get", func(c *gin.Context) {
		result := test.GetPublicUrl("product_image", "assistBook.jpg")
		fmt.Println("=========================================")
		log.Println(result)
		fmt.Println("=========================================")
		utility.HttpSuccessResponse(c, "s", result)
	})
	handler.StartEngine(r, db)
	r.Run()
}
