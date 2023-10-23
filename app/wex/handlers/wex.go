package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	wex "github.com/miguelhbrito/wexAssigment/internal/core/wex"
	data "github.com/miguelhbrito/wexAssigment/internal/data/wex"
)

type WexHandler struct {
	core wex.WexInt
	log  *log.Logger
}

func NewWexHandler(env *Env) *WexHandler {

	db := data.WexPostgres{
		Db: env.db,
	}

	core := wex.NewCore(
		db,
		env.log,
	)

	return &WexHandler{
		core: core,
		log:  env.log,
	}
}

func (h WexHandler) save(c *gin.Context) {
	h.log.Printf("%s %s -> %s", c.Request.Method, c.Request.URL, c.Request.RemoteAddr)
	h.log.Printf("receive request to save a new transaction")

	var newTransaction data.Transaction

	// Call BindJSON to bind the received JSON to
	// newTransaction.
	if err := c.BindJSON(&newTransaction); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tr, err := h.core.Save(newTransaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, tr)
}

func (h WexHandler) get(c *gin.Context) {
	h.log.Printf("%s %s -> %s", c.Request.Method, c.Request.URL, c.Request.RemoteAddr)
	h.log.Printf("receive request to get a transaction")

	id := c.Param("id")

	h.log.Println("here id", id)
	transaction, err := h.core.Get(id)
	if err != nil {
		h.log.Printf("Error to get a transaction from db: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h WexHandler) list(c *gin.Context) {
	h.log.Printf("%s %s -> %s", c.Request.Method, c.Request.URL, c.Request.RemoteAddr)
	h.log.Printf("receive request to list all transactions")

	transactions, err := h.core.List()
	if err != nil {
		h.log.Printf("Error to list transactions from db: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Transactions: ": transactions,
	})
}
