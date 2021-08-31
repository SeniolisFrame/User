package main

import (
	"fmt"
	"gingorm/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){

	
	err := godotenv.Load("db.env");
   	if err != nil {
      fmt.Println("Error loading .env file")
	  return
    }

	user := controller.New()
	r := gin.Default();
	r.GET("/",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message": "Hello",
		})
	});
	
	r_user := r.Group("/user")
	{
		r_user.POST("/create",user.CreateUser)
		r_user.GET("/gets",user.GetUsers)
		r_user.GET("/get",user.GetUser)
		r_user.PUT("/edit",user.EditUser)
		r_user.DELETE("/delete",user.DeleteUser)
	}
	
	r.Run(":8080");
}
