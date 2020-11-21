package database

import (
	"log"

	"github.com/jokekerker/gofinal/customer"
)

func QueryUpdateCustomerByID(customerId int, cs customer.Customer) (customer.Customer, error) {

	updateDb := `
		update customer set name=$2, email=$3, status=$4 where id=$1;
	`
	stmt, err := db.Prepare(updateDb)

	if err != nil {
		log.Fatal("can't prepare statement update", err)
		return cs, err
	}

	if _, err := stmt.Exec(customerId, &cs.Name, &cs.Email, &cs.Status); err != nil {
		log.Fatal("error execute update ", err)
		return cs, err
	}

	return cs, nil
}
