package service

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"randomgifsite/internal/utils"
	"strings"

	"github.com/mowshon/moviego"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func CatGift() (string, error) {
	links := utils.GetGifUrl

	gifUrl, err := utils.ParceUrl(links()[1])
	if err != nil {
		return "", err
	}

	gifname, err := utils.GetFileName(gifUrl)
	if err != nil {
		return "", err
	}

	// Проверка гифнейма на совпадения с уже сохраненными
	if fileIsNew(gifname) {
		gifPath := "../../output/" + gifname
		utils.DownloadFile(gifUrl, gifPath)
		fmt.Printf("save video %s \n", gifname)
		numberOfSave()
	}

	return gifname, nil

}

func fileIsNew(uploadedFile string) (saveFile bool) {

	saveFile = true

	files, err := ioutil.ReadDir("../../output/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if uploadedFile == file.Name() {
			log.Printf("File with that name already exists: %s", file.Name())
			numberOfMatches()
			saveFile = false
			return
		}
	}

	return
}

func amountFils(path string) (num int) {

	files, err := ioutil.ReadDir(path) // "../../output/"
	if err != nil {
		log.Fatal(err)
	}

	for _, _ = range files {
		num++

	}

	return num
}

func DownloadVideo(num int) {
	for i := 0; i < num; i++ {
		CatGift()
	}

	fmt.Printf("number of matches: %d \n", numMatch)
	fmt.Printf("number save: %d \n", numSave)

	fmt.Printf("number all video: %d", amountFils("../../output/"))
}


func EditSizeVideo() {

	files, err := ioutil.ReadDir("../../output/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		vidpath := "../../output/" + file.Name() // путь до видео

		editvideo, err := moviego.Load(vidpath)
		if err != nil {
			log.Println("video don't load")
		}
		newVidPath := "../../ouput_resize/" + file.Name()

		if file.Size() > 1048576 { // if video > 1mb -> resize
			err = editvideo.Resize(250, 250).Output(newVidPath).Run()
			if err != nil {
				log.Println("video don't save")
			}
			log.Printf("Video resize and save: %s", file.Name())
		} else {
			editvideo.Output(newVidPath).Run()
			log.Printf("Video save: %s", file.Name())
		}
		os.Remove(vidpath)
	}

}

func VideoToGif() {

	files, err := ioutil.ReadDir("../../ouput_resize/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		videoPath := "../../ouput_resize/" + file.Name()

		fileName := strings.TrimSuffix(file.Name(), ".webm")

		output := "../../gif/" + fileName + ".gif"

		err := Convert(videoPath, output)
		if err != nil {
			fmt.Println("error convert video to gif")
		}

	}

}

// Convert converts video to gif	go get github.com/u2takey/ffmpeg-go
func Convert(from, to string) error {

	cmd := ffmpeg.Input(from, ffmpeg.KwArgs{}).Output(to)

	return cmd.Run()
}

var numMatch int

func numberOfMatches() {

	numMatch++

}

var numSave int

func numberOfSave() {
	numSave++
}
