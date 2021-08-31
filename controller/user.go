package controller

import (
	"gingorm/database"
	"gingorm/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	db.AutoMigrate(&model.User{})
	return &UserRepo{Db: db}
}

func (repository *UserRepo) CreateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	err := repository.Db.Create(&user).Error;
	if (err != nil) {
	   c.JSON(400, gin.H{"error": err})
	   return
	}
	c.JSON(200, user);
}

func (repository *UserRepo) GetUsers(c *gin.Context){
	var user []model.User
	err := repository.Db.Find(&user).Error
	if (err != nil) {
	   c.JSON(400, gin.H{"error": err})
	   return
	}
	c.JSON(200, user)
}

func (repository *UserRepo) GetUser(c *gin.Context){
	var user model.User
	id := c.Query("id");
	err := repository.Db.Where("id = ?",id).First(&user).Error
	if (err != nil) {
	   c.JSON(400, gin.H{"error": err})
	   return
	}
	c.JSON(200, user)
}

func (repository *UserRepo) EditUser(c *gin.Context){
	var user model.User
	c.BindJSON(&user)
	err := repository.Db.Save(&user)
	if (err != nil) {
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, user)
}

func (repository *UserRepo) DeleteUser(c *gin.Context){
	var user model.User
	id := c.Query("id");
	err := repository.Db.Where("id =?",id).Delete(&user).Error;
	if (err != nil){
		c.JSON(400, gin.H{"error": err})
		return
	}
	c.JSON(200, user)
}