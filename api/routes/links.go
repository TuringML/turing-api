package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/api/models"
	"github.com/turing-ml/turing-api/api/utils"
)

// GetLinks returns all the links from/to a node
// @Title GetLinks
// @Description Get all the links from/to a node
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Success 200 {array} models.Link Array of Links
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{playground_id}/nodes/{node_id}/links [get]
func GetLinks(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	nodeID, err := strconv.Atoi(c.Param("node_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	nodes, err := models.GetLinks(db, nodeID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, nodes)
}

// GetLink returns a single link given the ID
// @Title GetLink
// @Description Get a single link given the ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Param   link_id     path    string     true        "The link ID"
// @Success 200 {object} models.Link Link object
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{playground_id}/nodes/{node_id}/links/{link_id} [get]
func GetLink(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	linkID, err := strconv.Atoi(c.Param("link_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	link, err := models.GetLink(db, linkID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, link)
}

// CreateLink returns a single link from a node
// @Title CreateLink
// @Description Create a single link from a node
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Success 200 {object} models.Link Link object
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{playground_id}/nodes/{node_id}/links [post]
func CreateLink(c *gin.Context) {
	var l models.Link
	err := c.BindJSON(&l)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return
	}

	db := c.MustGet("DB").(*gorm.DB)
	link, err := models.CreateLink(db, l)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, link)
}

// UpdateLink updates a single link
// @Title UpdateLink
// @Description Updates a single link
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Param   link_id     path    string     true        "The link ID"
// @Success 200 {string} string "link updated"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{playground_id}/nodes/{node_id}/links/{link_id} [put]
func UpdateLink(c *gin.Context) {
	var l models.Link
	err := c.BindJSON(&l)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(gin.ErrorTypeBind)
		return
	}

	db := c.MustGet("DB").(*gorm.DB)

	err = models.UpdateLink(db, l)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "node updated"})
}

// DeleteLink deletes a single link
// @Title DeleteLink
// @Description Deletes a single link
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Param   link_id     path    string     true        "The link ID"
// @Success 200 {string} string "link deleted"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{playground_id}/nodes/{node_id}/links/{link_id} [delete]
func DeleteLink(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	linkID, err := strconv.Atoi(c.Param("link_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	err = models.DeleteLink(db, linkID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "node deleted"})
}
