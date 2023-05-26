package handlers

import (
	"net/http"
	"randomgifsite/internal/service"
)

func GetCatGif(w http.ResponseWriter, r *http.Request) {

	service.CatGift()

	// gifName, _ := service.CatGift()
	// gifPath := "../../output/" + gifName
	// http.ServeFile(w, r, gifPath)
	// utils.DeleteFile(gifPath)
}
