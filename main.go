// @APIVersion 0.0.1
// @APITitle TuringML APIs
// @APIDescription This is the RESTful APIs of TuringML
// @BasePath https://api.turingml.org/
// @Contact davideberdin@gmail.com
// @TermsOfServiceUrl https://turingml.org/
// @License MIT
// @LicenseUrl http://turingml.org/
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/turing-ml/turing-api/cmd"
)

func main() {
	cmd.Execute()
}
