package database

import (
	"github.com/jokekerker/gofinal/customer"
)

// var customers = map[int]*customer.Customer{}

func QueryAllCustomer() ([]customer.Customer, error) {
	customers := []customer.Customer{}

	queryDb := `
		select * from customer
	`

	stmt, err := db.Prepare(queryDb)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c := customer.Customer{}
		err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Status)

		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, err
}
