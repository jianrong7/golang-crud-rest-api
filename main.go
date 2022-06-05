package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	router.Use(mux.CORSMethodMiddleware(router))

	c := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:3000"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowCredentials: true,
			// Enable Debugging for testing, consider disabling in production
			Debug: true,
	})

	handler := c.Handler(router)

	// Start the server
	log.Printf(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), handler))
}


func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.HelloWorld).Methods("GET")
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}