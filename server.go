package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
		log.Fatal("get all customer error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func getCustomerByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	cs := customer.Customer{}
	if err != nil {
		log.Fatal("get params id error", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	cs, err = database.QueryCustomerById(id)

	if err != nil {
		log.Fatal("get customer by id error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	c.JSON(http.StatusOK, cs)
}

func createCustomerHandler(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal("get request body error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	cs := customer.Customer{}
	err = json.Unmarshal(jsonData, &cs)
	if err != nil {
		log.Fatal("Unmarshal json error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	cs, err = database.QueryCreateCustomer(cs)

	if err != nil {
		log.Fatal("query create customer error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusCreated, cs)
}

func updateCustomerHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("get params id error", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal("get request body error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	cs := customer.Customer{ID: id}

	err = json.Unmarshal(jsonData, &cs)
	if err != nil {
		log.Fatal("Unmarshal json error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	cs, err = database.QueryUpdateCustomerByID(id, cs)
	if err != nil {
		log.Fatal("query update error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, cs)
}

func deleteCustomerByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	message, err := database.QueryDeleteCustomerByID(id)

	if err != nil {
		log.Fatal("query delete error", err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

func main() {
	r := gin.Default()

	r.GET("/customers", getAllCustomerHandler)
	r.GET("/customers/:id", getCustomerByIDHandler)
	r.POST("/customers", createCustomerHandler)
	r.PUT("/customers/:id", updateCustomerHandler)
	r.DELETE("/customers/:id", deleteCustomerByIDHandler)
	r.Run(":2009")
}
