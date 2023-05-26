package routes

import (
	"randomgifsite/internal/transport/handlers"

	"github.com/gorilla/mux"
)

var RandomGifSiteRouter = func(router *mux.Router) {
	router.HandleFunc("/randomgif/cat", handlers.GetCatGif).Methods("GET")
}
