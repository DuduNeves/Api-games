package controllers

import (
	"strconv"

	"github.com/DuduNeves/dev-games.git/database"
	"github.com/DuduNeves/dev-games.git/models"
	"github.com/gin-gonic/gin"
)

func GetGame(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var games models.Games
	err = db.First(&games, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find game:" + err.Error(),
		})
		return
	}

	c.JSON(200, games)
}

func GetAllGames(c *gin.Context) {

	db := database.GetDatabase()

	var games []models.Games
	err := db.Find(&games).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list games:" + err.Error(),
		})

		return
	}

	c.JSON(200, games)
}

func CreateGames(c *gin.Context) {
	db := database.GetDatabase()

	var game models.Games

	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON:" + err.Error(),
		})
		return
	}

	err = db.Create(game).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create game:" + err.Error(),
		})

		return
	}

	c.JSON(200, game)
}

func UpdateGames(c *gin.Context) {

	db := database.GetDatabase()

	var game models.Games

	err := c.ShouldBindJSON(&game)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON:" + err.Error(),
		})
		return
	}

	err = db.Save(game).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update game:" + err.Error(),
		})

		return
	}

	c.JSON(200, game)
}

func DeleteGames(c *gin.Context) {

	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Games{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot delete a book",
		})

		return
	}

	c.Status(204)
}
