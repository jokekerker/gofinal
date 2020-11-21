package database

import (
	"github.com/jokekerker/gofinal/customer"
)

func QueryCustomerById(customerId int) (customer.Customer, error) {
	defer db.Close()

	cs := customer.Customer{}

	queryDb := `
		select id, name, email, status from customer where id=$1
	`

	stmt, err := db.Prepare(queryDb)

	if err != nil {
		return cs, err
	}
	row := stmt.QueryRow(customerId)

	if err != nil {
		return cs, err
	}

	err = row.Scan(&cs.ID, &cs.Name, &cs.Email, &cs.Status)

	return cs, err
}
