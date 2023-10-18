package handler

import (
	"Valter/db/sqlc"
	"Valter/middleware"
	"Valter/repository"
	"Valter/service"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Handler(db *sql.DB) (*UserHandler, *ProductHandler, *FeatHandler, *BookingHandler) {
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

	bookRepo := repository.NewBookingRepository(queries)
	bookServ := service.NewBookingService(bookRepo)
	bookhand := NewBookingHandler(bookServ, prodServ)
	return userHand, prodHand, featHand, bookhand

}

func route(r *gin.Engine, uh *UserHandler, ph *ProductHandler, bh *BookingHandler) {
	//user side
	r.POST("/register", uh.Register)
	r.POST("/login", uh.Login)
	r.GET("/user-profile", middleware.Auth(), uh.GetDataUser)
	r.POST("/forgot-pass", uh.ForgotPass)
	r.POST("/verify-code", uh.VerfiyCode)

	//Product
	r.GET("/get-all-products", ph.GetAllProducts)
	r.GET("/product-details/:id", ph.GetProductById)

	//Booking
	r.POST("/bookings/:id", middleware.Auth(), bh.CreateBooking)
}

func StartEngine(r *gin.Engine, db *sql.DB) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
	uh, ph, _, bh := Handler(db)
	route(r, uh, ph, bh)
	//ph.Dummy()
	//fh.FeatDummy()
}
