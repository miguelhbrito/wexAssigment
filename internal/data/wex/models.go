package wex

import (
	"time"
)

type Transaction struct {
	Id          string    `db:"id" json:"id"`
	Description string    `db:"description" json:"description"`
	Price       float64   `db:"price" json:"price"`
	Country     string    `db:"country" json:"country"`
	DataCreated time.Time `db:"created_at" json:"dataCreated"`
}

type Transactions []Transaction
