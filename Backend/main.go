package main

import (
	"github.com/gin-gonic/gin"
	"github.com/panupongKanin/ProjectSA-arm/controller"
	"github.com/panupongKanin/ProjectSA-arm/entity"
)

func main() {
	entity.SetupDatabase()
	r := gin.Default()
	r.Use(CORSMiddleware())

	//User Router
	r.GET("/GetUser", controller.GetUser)
	r.POST("/CreateUser", controller.CreateUser)

	//Nutrition Router
	r.GET("/nutritions", controller.ListNutritions)
	r.GET("/nutrition/:id", controller.GetNutrition)
	//Manage Router
	r.POST("/CreateManage", controller.CreateManage)
	r.GET("/ListManage", controller.ListManage)

	// login User Route
	r.POST("/signup", controller.CreateUser)
	r.POST("/login", controller.Login)
	r.GET("/user/:id", controller.GetUser)
	
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
