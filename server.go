package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jokekerker/gofinal/customer"
	"github.com/jokekerker/gofinal/database"
)

var customers = map[int]*customer.Customer{}

func getAllCustomerHandler(c *gin.Context) {

	customers, err := database.QueryAllCustomer()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func getCustomerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	cs := customer.Customer{}
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	cs, err = database.QueryCustomerById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, cs)
}

func main() {
	r := gin.Default()

	r.GET("/customers", getAllCustomerHandler)
	r.GET("/customer/:id", getAllCustomerHandler)
	// r.GET("/customers", getAllCustomerHandler)
	// r.GET("/customers", getAllCustomerHandler)
	r.Run(":2009")
}
