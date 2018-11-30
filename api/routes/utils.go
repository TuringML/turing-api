package routes

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"github.com/turing-ml/turing-api/api/utils"
	"github.com/turing-ml/turing-api/pkg/version"
)

var outputTemplate = template.Must(template.New("base").Parse(`
<html>
  <head>
    <title>{{ .Path }}</title>
  </head>
  <body>
    {{ .Body }}
  </body>
</html>
`))

// Version returns the APIs version
func Version(c *gin.Context) {
	utils.Response(c, http.StatusOK, version.LongVersion())
}

// Endpoints will return the documentation regarding each endpoint
func Endpoints(c *gin.Context) {
	input, err := ioutil.ReadFile("./docs/API.md")
	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	output := blackfriday.Run(input)
	outputTemplate.Execute(c.Writer, struct {
		Path string
		Body template.HTML
	}{
		Path: "TuringML API Documentation",
		Body: template.HTML(string(output)),
	})
}
