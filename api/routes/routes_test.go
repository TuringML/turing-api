package routes

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/turing-ml/turing-api/api/middleware"
	"github.com/turing-ml/turing-api/api/utils"
	"github.com/turing-ml/turing-api/pkg/database"
	testfixtures "gopkg.in/testfixtures.v2"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var testDb *gorm.DB

func TestMain(m *testing.M) {
	tearUp()
	c := m.Run()
	tearDown()
	os.Exit(c)
}

func tearUp() {
	gin.SetMode(gin.TestMode)

	router = gin.Default()
	router.RedirectTrailingSlash = true

	// Testing Database
	testDb = database.OpenConnection("root", "root", "127.0.0.1", "", false)

	testDb = testDb.Exec("CREATE DATABASE IF NOT EXISTS testing")
	testDb = testDb.Exec("USE testing")

	// Run migrations
	utils.RunMigration(testDb)

	// Load fixtures
	fixtures, err := testfixtures.NewFolder(testDb.DB(), &testfixtures.MySQL{}, "../fixtures")
	if err != nil {
		panic(err)
	}

	err = fixtures.Load()
	if err != nil {
		panic(err)
	}

	// Set up Middleware
	router.Use(middleware.DB(testDb))
}

func tearDown() {
	router = nil
	testDb.Exec("DROP DATABASE testing")
}

// MockRequest will send a request to the server. Used for testing purposes
func MockRequest(method, path string, body io.Reader) (int, *bytes.Buffer, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return -1, nil, err
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	return w.Code, w.Body, nil
}
