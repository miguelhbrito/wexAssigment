package wex

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type WexInt interface {
	Save(Transaction) (Transaction, error)
	List() (Transactions, error)
	Get(string) (Transaction, error)
}

type WexPostgres struct {
	Db *sql.DB
}

func (wp WexPostgres) Save(tr Transaction) (Transaction, error) {
	transaction := Transaction{
		Id:          uuid.New().String(),
		Description: tr.Description,
		Price:       tr.Price,
		Country:     tr.Country,
		DataCreated: tr.DataCreated,
	}

	sqlStatement := `INSERT INTO transaction VALUES ($1, $2, $3, $4, $5)`
	_, err := wp.Db.Exec(sqlStatement, transaction.Id, transaction.Description, transaction.Price, transaction.Country, transaction.DataCreated)
	if err != nil {
		log.Error().Err(err).Msgf("Error to insert an new transaction into db")
		return Transaction{}, err
	}

	return transaction, nil
}

func (wp WexPostgres) List() (Transactions, error) {
	var trs Transactions
	sqlStatement := `SELECT id, description, price, country, created_at FROM transaction`
	rows, err := wp.Db.Query(sqlStatement)
	if err != nil {
		log.Error().Err(err).Msg("Error to get all transactions from db")
		return nil, err
	}

	for rows.Next() {
		var tr Transaction
		err := rows.Scan(&tr.Id, &tr.Description, &tr.Price, &tr.Country, &tr.DataCreated)
		if err != nil {
			log.Error().Err(err).Msg("Error to extract result from row")
		}
		trs = append(trs, tr)
	}

	return trs, nil
}

func (wp WexPostgres) Get(id string) (Transaction, error) {

	var tr Transaction
	sqlStatement := `SELECT id, description, price, country, created_at FROM transaction WHERE id = $1`
	result := wp.Db.QueryRow(sqlStatement, id)
	err := result.Scan(&tr.Id, &tr.Description, &tr.Price, &tr.Country, &tr.DataCreated)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).Msg("Error, no rows in result")
			return Transaction{}, err
		}
		log.Error().Err(err).Msg("Error to extract result from row")
		return Transaction{}, err
	}

	return tr, nil
}
