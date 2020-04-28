package users

import (
	"net/http"
	"simple-crud/db"
	"simple-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

func List(c *gin.Context) {

	var userList []models.Users
	if err := db.DB.Find(&userList).Error; err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, userList)
}

func Add(c *gin.Context) {
	type object struct {
		Name string `binding:"required"`
		Age  int
	}
	var req object
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newUser := models.Users{Name: req.Name, Age: req.Age}
	if err := db.DB.Save(&newUser).Error; err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, newUser.ID)

}

func Delete(c *gin.Context) {
	type object struct {
		ID uint `binding:"required"`
	}
	var req object
	if err := c.ShouldBindWith(&req, binding.JSON); err != nil {
		logrus.Error(err)
		c.JSON(300, err.Error())
		return
	}
	if err := db.DB.Delete(&models.Users{ID: req.ID}).Error; err != nil {
		logrus.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, req.ID)

}
