package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/turing-ml/turing-api/models"
)

// Collector will test the collector and get the content of the first file
// @Title Collector
// @Description Test the collector and return the file detected and the content of it
// @Accept  json
// @Param   s3        query   object     false        "AWS S3 collector object"
// @Param   gcs        query   object     false        "Google Cloud Storage collector object"
// @Param   adl        query   object     false        "Azure Data Lake collector object"
// @Success 200 {string}  string
// @Failure 500 {string} string    "Internal Server Error"
// @Resource /user
// @Router /users/ [post]
func (a *App) Collector(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}

	var c models.Collectors
	err = json.Unmarshal(b, &c)
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}

	files, err := c.ListFiles()
	if err != nil {
		Response(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(files)
	Response(w, http.StatusOK, files)
}
