package routes

import (
	"net/http"
	"strconv"

	"github.com/turing-ml/turing-api/api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/api/utils"
)

// GetNodes returns all the nodes objects of the playground
// @Title GetNodes
// @Description Get all the nodes based on the playground ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Success 200 {array} models.Node Array of Nodes
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes [get]
func GetNodes(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	playgroundID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	nodes, err := models.GetNodes(db, playgroundID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, nodes)
}

// GetNode returns a single node given the ID
// @Title GetNode
// @Description Get a single node given the ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   nodeId     path    string     true        "The node ID"
// @Success 200 {object} models.Node Node object
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{nodeId} [get]
func GetNode(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	playgroundID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	nodeID, err := strconv.Atoi(c.Param("nodeId"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	node, err := models.GetNode(db, playgroundID, nodeID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, node)
}

// CreateNode returns a single node in the playground
// @Title CreateNode
// @Description Create a single node in the playground
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Success 200 {object} models.Node Node object
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{nodeId} [post]
func CreateNode(c *gin.Context) {
	var n models.Node
	// todo: fix post

	db := c.MustGet("DB").(*gorm.DB)

	playgroundID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	node, err := models.CreateNode(db, playgroundID, n)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, node)
}

// UpdateNode updates a single node in the playground given the ID
// @Title UpdateNode
// @Description Updates a single node in the playground given the ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   nodeId     path    string     true        "The node ID"
// @Success 200 {string} string "node updated"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{nodeId} [put]
func UpdateNode(c *gin.Context) {
	var n models.Node
	// todo: fix put

	db := c.MustGet("DB").(*gorm.DB)

	playgroundID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	nodeID, err := strconv.Atoi(c.Param("nodeId"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	err = models.UpdateNode(db, playgroundID, nodeID, n)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "node updated"})
}

// DeleteNode deletes a single node in the playground given the ID
// @Title DeleteNode
// @Description Deletes a single node in the playground given the ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   nodeId     path    string     true        "The node ID"
// @Success 200 {string} string "node deleted"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{nodeId} [put]
func DeleteNode(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	playgroundID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	nodeID, err := strconv.Atoi(c.Param("nodeId"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	err = models.DeleteNode(db, playgroundID, nodeID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "node deleted"})
}
