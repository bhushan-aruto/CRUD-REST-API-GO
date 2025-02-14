package controller

import (
	"log"
	"net/http"

	"github.com/bhushan-aruto/ice_cream_shop_rest_api/models"
	"github.com/bhushan-aruto/ice_cream_shop_rest_api/repository"
	"github.com/gin-gonic/gin"
)

func GetAllFlavourController(c *gin.Context) {
	flavours, err := repository.GetAllFlavours()
	if err != nil {
		log.Println("error occured ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch the error"})
		return
	}
	c.JSON(http.StatusOK, flavours)

}
func AddFlavourController(c *gin.Context) {
	var flavor models.Flavour
	if err := c.ShouldBindJSON(&flavor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "inavlid JSON"})
		return
	}
	if err := repository.AddFlavour(flavor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add the flavor"})
		return
	}
	c.JSON(http.StatusOK, flavor)

}
func GetFlavourByIdController(c *gin.Context) {
	id := c.Param("id")
	flavor, err := repository.GetFlavourById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, flavor)

}

func UpdateFlavourController(c *gin.Context) {
	id := c.Param("id")

	var updatedflavour models.Flavour
	if err := c.ShouldBindJSON(&updatedflavour); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	if err := repository.Updateflavour(id, updatedflavour); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "flavour updated succesfully"})
}

func DeleteFlavourController(c *gin.Context) {
	id := c.Param("id")
	if err := repository.DeleteFlavour(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Flavour  is deleted succusfully"})
}
