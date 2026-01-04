package main

import (
	"log"
	"net/http"

	myhttp "github.com/Nils13001/Notification-Service/internal/http"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Create router (Chi)
	r := chi.NewRouter()

	// Routes will go here
	r.Get("/api/notification/health", myhttp.HealthHandler)
	r.Post("/api/notification/notify", myhttp.NotifyHandler)

	//Defining starting port
	addr := ":8092"
	log.Printf("notification-service listening on %s", addr)

	// 		Start the server with the router as handler
	//    Using log.Fatal ensures we see startup errors clearly
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}

}
