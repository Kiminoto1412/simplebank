package api

import (
	db "github.com/Kiminoto1412/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add custom own validation to gin validation
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// currency => name to use from oneof=USD EUR CAD to currency
		v.RegisterValidation("currency", validCurrency)
	}

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	// transfer
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// gin.H is a shortcut for map[string]any
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
