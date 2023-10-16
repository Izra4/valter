package handler

import (
	"Valter/db/sqlc"
	"Valter/repository"
	"Valter/service"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Handler(db *sql.DB) (*UserHandler, *ProductHandler, *FeatHandler) {
	queries := sqlc.New(db)
	userRepo := repository.NewUserRepository(queries)
	userServ := service.NewUserService(userRepo)
	userHand := NewUserHandler(userServ)

	prodRepo := repository.NewProductRepository(queries)
	prodServ := service.NewProductService(prodRepo)
	prodHand := NewProductHandler(prodServ)

	featRepo := repository.NewFeatureRepository(queries)
	featServ := service.NewFeatService(featRepo)
	featHand := NewFeatHandler(featServ)
	return userHand, prodHand, featHand

}

func route(r *gin.Engine, uh *UserHandler, ph *ProductHandler) {
	//user side
	r.POST("/register", uh.Register)
	r.POST("/login", uh.Login)
	r.POST("/forgot-pass", uh.ForgotPass)
	r.POST("/verify-code", uh.VerfiyCode)

	//Product
	r.GET("/get-all-products", ph.GetAllProducts)
	r.GET("/product-details/:id", ph.GetProductById)
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})
	uh, ph, fh := Handler(db)
	route(r, uh, ph)
	ph.Dummy()
	fh.FeatDummy()
}
