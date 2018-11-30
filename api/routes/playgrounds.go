package routes

import (
	"net/http"
	"strconv"

	"github.com/turing-ml/turing-api/api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/api/utils"
)

// GetPlaygrounds returns all the playground objects of the user in input
// @Title GetPlaygrounds
// @Description Get all the playgrounds based on the user ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The user ID"
// @Success 200 {array} models.Playground Array of playgrounds
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds [get]
func GetPlaygrounds(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	// TODO: fix me with auth0 id
	playground, err := models.GetPlaygrounds(db, "auth0|id")
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, playground)
}

// GetPlayground return a single playground object given the ID in input
// @Title GetPlayground
// @Description Get a specific playground based on the ID in input of a single user
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   playground_id     path    string     true        "The playground ID"
// @Success 200 {object} models.Playground
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{playground_id} [get]
func GetPlayground(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	playgroundID, err := strconv.Atoi(c.Param("playground_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	playground, err := models.GetPlayground(db, "auth0|id", playgroundID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, playground)
}

// CreatePlayground creates a new playground in the database
// @Title CreatePlayground
// @Description Create a new playground based on the parameters in input
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   name     query    string     true        "The playground name"
// @Success 200 {object} models.Playground
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds [post]
func CreatePlayground(c *gin.Context) {
	var p models.Playground
	err := c.BindJSON(&p)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	playground, err := models.CreatePlayground(db, "auth0|id", p)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, playground)
}

// UpdatePlayground updates the playground's properties in the database
// @Title UpdatePlayground
// @Description Update the playground of a specific user
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   playground_id     path    string     true        "The playground ID"
// @Param   name     query    string     true        "The playground name"
// @Success 200 {string} string	"playground updated"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /users/{user_id}/playgrounds/{playground_id} [put]
func UpdatePlayground(c *gin.Context) {
	var p models.Playground
	err := c.BindJSON(&p)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	playgroundID, err := strconv.Atoi(c.Param("playground_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	err = models.UpdatePlayground(db, "auth0|", playgroundID, p)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "playground updated"})
}

// DeletePlayground deletes the specified playground from the database
// @Title DeletePlayground
// @Description Delete the playground of a specific user
// @Accept  json
// @Produce  json
// @Param   user_id     path    string     true        "The user ID"
// @Param   playground_id     path    string     true        "The playground ID"
// @Success 200 {string} string	"playground deleted"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /users/{user_id}/playgrounds/{playground_id} [delete]
func DeletePlayground(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	playgroundID, err := strconv.Atoi(c.Param("playground_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	err = models.DeletePlayground(db, "auth0|", playgroundID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "playground deleted"})
}
