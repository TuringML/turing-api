package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/api/models"
	"github.com/turing-ml/turing-api/api/utils"
)

// GetFields returns all the fields of a node
// @Title GetFields
// @Description Get all the fields of a node
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Success 200 {array} models.Link Array of Links
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{node_id}/fields [get]
func GetFields(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	nodeID, err := strconv.Atoi(c.Param("node_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	nodes, err := models.GetFields(db, nodeID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, nodes)
}

// GetField returns a single field given the ID
// @Title GetField
// @Description Get a single field given the ID
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Param   field_id     path    string     true        "The field ID"
// @Success 200 {object} models.Field Field object
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{node_id}/fields/{field_id} [get]
func GetField(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	nodeID, err := strconv.Atoi(c.Param("node_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	fieldID, err := strconv.Atoi(c.Param("field_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	field, err := models.GetField(db, nodeID, fieldID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, field)
}

// CreateField returns a single field for a node
// @Title CreateField
// @Description Create a single field for a node
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Success 200 {object} models.Field Field object
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{node_id}/fields [post]
func CreateField(c *gin.Context) {
	var f models.Field
	// todo: fix post

	db := c.MustGet("DB").(*gorm.DB)

	nodeID, err := strconv.Atoi(c.Param("node_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	field, err := models.CreateField(db, nodeID, f)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, field)
}

// UpdateField updates a single field of a node
// @Title UpdateField
// @Description Updates a single field of a node
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Param   field_id     path    string     true        "The field ID"
// @Success 200 {string} string "field updated"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{node_id}/fields/{field_id} [put]
func UpdateField(c *gin.Context) {
	var f models.Field
	// todo: fix put

	db := c.MustGet("DB").(*gorm.DB)

	nodeID, err := strconv.Atoi(c.Param("node_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	fieldID, err := strconv.Atoi(c.Param("field_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	err = models.UpdateField(db, nodeID, fieldID, f)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "field updated"})
}

// DeleteField deletes a single field
// @Title DeleteField
// @Description Deletes a single field
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "The playground ID"
// @Param   node_id     path    string     true        "The node ID"
// @Param   field_id     path    string     true        "The field ID"
// @Success 200 {string} string "link deleted"
// @Failure 500 {string} string	"Internal Server Error"
// @Router /playgrounds/{id}/nodes/{node_id}/fields/{field_id} [delete]
func DeleteField(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	nodeID, err := strconv.Atoi(c.Param("node_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	fieldID, err := strconv.Atoi(c.Param("field_id"))
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	err = models.DeleteField(db, nodeID, fieldID)
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, map[string]string{"message": "field deleted"})
}
