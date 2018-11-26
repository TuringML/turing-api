package app

import (
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/russross/blackfriday"
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

// Info returns the version of the api currently running
func Info(w http.ResponseWriter, r *http.Request) {
	Response(w, http.StatusOK, map[string]string{"version": version.LongVersion()})
}

// Endpoints will return the documentation regarding each endpoint
func Endpoints(w http.ResponseWriter, r *http.Request) {
	input, err := ioutil.ReadFile("./docs/API.md")
	if err != nil {
		Response(w, http.StatusInternalServerError, "Page not available")
		return
	}

	output := blackfriday.Run(input)
	outputTemplate.Execute(w, struct {
		Path string
		Body template.HTML
	}{
		Path: "TuringML API Documentation",
		Body: template.HTML(string(output)),
	})
}
