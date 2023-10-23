package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Env struct {
	log *log.Logger
	db  *sql.DB
}

func NewGin(log *log.Logger, db *sql.DB) http.Handler {
	router := gin.New()

	env := &Env{
		db:  db,
		log: log,
	}

	//Starting postgres and core
	wex := NewWexHandler(env)

	router.GET("/wex", wex.list)
	router.POST("/wex", wex.save)
	router.GET("/wex/:id", wex.get)

	return router
}
