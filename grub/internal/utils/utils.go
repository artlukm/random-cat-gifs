package utils

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/gocolly/colly"
)

func ParceUrl(path string) (string, error) {
	ur, err := url.ParseRequestURI(path)
	if err != nil {
		return "", err
	}
	return ur.String(), nil
}

func GetGifUrl() []string {

	var Links []string
	c := colly.NewCollector(

		colly.AllowedDomains("randomcatgifs.com"),
	)

	c.OnHTML("source", func(e *colly.HTMLElement) {
		link := e.Attr("src")
		Links = append(Links, link)

	})

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://randomcatgifs.com")
	return Links
}

func DownloadFile(url, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return nil
	}
	return nil
}

func GetFileName(url string) (string, error) {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	return path.Base(r.URL.Path), nil
}

func DeleteFile(filepath string) error {
	return os.Remove(filepath)
}
