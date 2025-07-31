package main

import (
	"log"
	"net/http"

	"qr-code-generator/pkg/config"
	"qr-code-generator/pkg/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	/* Initialize database connection SQL
	config.Connect()
	db := config.GetDB()
	models.SetDB(db)
	db.AutoMigrate(&models.QRCode{})*/
	config.ConnectFirebase()

	// Use Gorilla Mux
	func SetupRoutes() *mux.Router {
	    router := mux.NewRouter()
	    RegisterQRCodeGeneratorstoreRoutes(router)          // your other routes
	    RegisterQRCodeGeneratorFirebasestoreRoutes(router) // your firebase routes
	    return router
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // React dev server
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)

	log.Println("Starting server on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", handler)) // Use mux router here

}
