package database

import "github.com/jokekerker/gofinal/customer"

func QueryCreateCustomer(cs customer.Customer) (customer.Customer, error) {

	insertDb := `
		insert into customer 
		(name, email, status) 
		values 
		($1, $2, $3) 
		returning id;
	`

	row := db.QueryRow(insertDb, &cs.Name, &cs.Email, &cs.Status)
	err := row.Scan(&cs.ID)

	return cs, err

}
