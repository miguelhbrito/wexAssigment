package wex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	data "github.com/miguelhbrito/wexAssigment/internal/data/wex"
)

type Currency struct {
	Data []struct {
		ExchangeRate float64 `json:"exchange_rate,string"`
		RecordDate   string  `json:"record_date"`
	} `json:"data"`
}

const (
	YYYYMMDD = "2006-01-02"
)

type WexInt interface {
	Save(data.Transaction) (data.Transaction, error)
	List() (data.Transactions, error)
	Get(string) (data.Transaction, error)
}

type core struct {
	data data.WexInt
	log  *log.Logger
}

func NewCore(data data.WexInt, log *log.Logger) WexInt {
	return core{
		data: data,
		log:  log,
	}
}

// Save triangle into db
func (c core) Save(tr data.Transaction) (data.Transaction, error) {
	c.log.Println("checking values to save", tr)
	// Check if price it is a float number
	/*var price interface{} = tr.Price
	if _, ok := price.(int); ok {
		c.log.Println("it is a price format")
	} else {
		c.log.Println("it is not a price format")
		return data.Transaction{}, nil
	}*/

	tr.DataCreated = time.Now()
	transaction, err := c.data.Save(tr)
	if err != nil {
		return data.Transaction{}, err
	}

	return transaction, nil
}

// Get transaction by id from db
func (c core) Get(id string) (data.Transaction, error) {

	tr, err := c.data.Get(id)
	if err != nil {
		return data.Transaction{}, err
	}

	//Subtraction of 6 months
	toCurrencyTime := time.Now().UTC().Add(time.Duration(-262800) * time.Minute).Format(YYYYMMDD)

	requestURL := fmt.Sprintf("https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=exchange_rate,record_date&filter=country_currency_desc:eq:%s,record_date:gte:%s", tr.Country, toCurrencyTime)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("could not create request to get current currency: %s\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.log.Println("error making http request to get current currency:", err)
		return data.Transaction{}, err
	}
	c.log.Println("client: status code:", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.log.Println("could not read response body:", err)
		return data.Transaction{}, err
	}

	var result Currency
	if err := json.Unmarshal(resBody, &result); err != nil {
		c.log.Println("can not unmarshal JSON", err)
		return data.Transaction{}, err
	}

	// Putting the current value into transaction's price
	tr.Price *= result.Data[0].ExchangeRate
	c.log.Println("here new price", tr.Price, result.Data[0].ExchangeRate)

	return tr, nil
}

// List all transactions from db
func (c core) List() (data.Transactions, error) {
	transactions, err := c.data.List()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
