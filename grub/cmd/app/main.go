package main

import (
	"randomgifsite/internal/service"
)

func main() {
	// r := mux.NewRouter()
	// routes.RandomGifSiteRouter(r)
	// http.Handle("/", r)
	// log.Fatal(http.ListenAndServe("localhost:8080", r))
	// fmt.Println("Server start")

	// service.DownloadVideo(200)

	service.VideoToGif()
}
