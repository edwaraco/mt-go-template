package main

import (
	"log"
	"net/http"

	mt_app "github.com/mitribu/mt-residentials/pkg/residential_units/application"
	mt_rest "github.com/mitribu/mt-residentials/pkg/residential_units/infrastructure/http/rest"
	mt_storage "github.com/mitribu/mt-residentials/pkg/residential_units/infrastructure/storage"
)

func setup() *mt_rest.Server {
	res_store := mt_storage.NewInMemoryRepository()
	res_app := mt_app.NewResidentialApp(res_store)
	res_route := mt_rest.NewServer(res_app)
	return res_route
}

func main() {
	log.Printf("Starting server...")
	res_route := setup()
	// TODO: It should retrieving from env-vars
	log.Printf("Listen server on 0.0.0.0:8000...")
	http.ListenAndServe(":8000", res_route)
}
