package database

import (
	"log"
)

func QueryDeleteCustomerByID(customerId int) (string, error) {

	queryDb := `
		delete from customer where id=$1
	`
	stmt, err := db.Prepare(queryDb)

	if err != nil {
		log.Fatal("can't prepare query delete statment", err)
		return "", err
	}
	if _, err := stmt.Exec(customerId); err != nil {
		log.Fatal("error execute update ", err)
		return "", err
	}

	return "customer deleted", nil
}
