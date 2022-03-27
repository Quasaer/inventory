package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Quasaer/goinventory-api/postgres"
	"github.com/Quasaer/goinventory-api/web"
)

func main() {

	connectionString, err := ioutil.ReadFile("/var/openfaas/secrets/inventory-postgres-connection-string")
	if err != nil {
		log.Fatal(err)
	}
	store, err := postgres.NewStore(string(connectionString))

	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":8082", h)
}
