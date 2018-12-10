package api

import (
	"github.com/gin-gonic/gin"
	"github.com/turing-ml/turing-api/api/middleware"
	"github.com/turing-ml/turing-api/api/routes"
	"github.com/turing-ml/turing-api/api/utils"
	"github.com/turing-ml/turing-api/pkg/database"
	"github.com/turing-ml/turing-api/pkg/vault"
)

// SetupRouter defines all the endpoints of the APIs
func SetupRouter(secret, dbUsername, dbPassword, dbHost, dbName, vaultToken, vaultAddr string) *gin.Engine {

	r := gin.Default()
	r.RedirectTrailingSlash = true

	// Init Vault
	v := vault.New(vaultToken, vaultAddr)

	// Database migrations
	db := database.OpenConnection(dbUsername, dbPassword, dbHost, dbName, true)
	utils.RunMigration(db)

	// Set up Middleware
	r.Use(middleware.Vault(v))
	r.Use(middleware.DB(db))

	r.GET("/version", routes.Version)
	r.GET("/docs", routes.Endpoints)

	// Playground Endpoints
	p := r.Group("/playgrounds")
	p.GET("", routes.GetPlaygrounds)
	p.POST("", routes.CreatePlayground)
	p.GET("/:id", routes.GetPlayground)
	p.PUT("/:id", routes.UpdatePlayground)
	p.DELETE("/:id", routes.DeletePlayground)

	// Nodes Endpoints
	p.GET("/:id/nodes", routes.GetNodes)
	p.POST("/:id/nodes", routes.CreateNode)
	p.GET("/:id/nodes/:node_id", routes.GetNode)
	p.PUT("/:id/nodes/:node_id", routes.UpdateNode)
	p.DELETE("/:id/nodes/node_id", routes.DeleteNode)

	// Fields Endpoints
	p.GET("/:id/nodes/:node_id/fields", routes.GetFields)
	p.POST("/:id/nodes/:node_id/fields", routes.CreateField)
	p.GET("/:id/nodes/:node_id/fields/:field_id", routes.GetField)
	p.PUT("/:id/nodes/:node_id/fields/:field_id", routes.UpdateField)
	p.DELETE("/:id/nodes/node_id/fields/:field_id", routes.DeleteField)

	// Links Endpoints
	p.GET("/:id/nodes/:node_id/links", routes.GetLinks)
	p.POST("/:id/nodes/:node_id/links", routes.CreateLink)
	p.GET("/:id/nodes/:node_id/links/:link_id", routes.GetLink)
	p.PUT("/:id/nodes/:node_id/links/:link_id", routes.UpdateLink)
	p.DELETE("/:id/nodes/node_id/links/:link_id", routes.DeleteLink)

	return r
}

// Serve will serve the APIs on a specific address
func Serve(addr, secret, dbUsername, dbPassword, dbHost, dbName, vaultToken, vaultAddr string) error {
	r := SetupRouter(secret, dbUsername, dbPassword, dbHost, dbName, vaultToken, vaultAddr)
	return r.Run(addr)
}
