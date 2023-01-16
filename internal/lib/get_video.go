package lib

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	goq "github.com/PuerkitoBio/goquery"
)

// в коде присутствуют имена ошибок по типу ErrStatusNotOK или ErrNilQueryPointer.
// эти ошибки объявлены отдельно в файле errors.go

// GetVideoURL returns URL of video in WebM format
func (c *Client) GetVideoURL(ctx context.Context) (string, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL, nil)
	if err != nil {
		return "", err
	}

	req = req.WithContext(ctx)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", ErrStatusNotOK
	}
	defer resp.Body.Close()

	doc, err := goq.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	query := doc.Find("source") // ищем теги <source>
	if query == nil {
		return "", ErrNilQueryPointer
	} else if query.Nodes == nil {
		if c.Debug {
			fmt.Printf("%v, %v\n", *query, query.Nodes)
		}
		return "", ErrNilNodesArray
	} else if len(query.Nodes) == 0 {
		return "", ErrEmptyNodesArray
	}
	node := query.Last().Get(0) // берём последний тег из списка (в последнем находится webm-файл с котом)
	if node == nil {
		return "", ErrNilNodePointer
	} else if node.Attr == nil {
		return "", ErrNilAttrArray
	} else if len(node.Attr) == 0 {
		return "", ErrEmptyAttrArray
	}
	var url string
	for _, attr := range node.Attr {
		if attr.Key == "src" {
			url = attr.Val
			continue
		}
	}
	if url == "" {
		return "", ErrSrcAttrNotFound
	}
	return url, nil
}

// GetVideo gets video in WebM format
// Be careful: it uses the same context for both GetVideoURL
// and GetVideo
func (c *Client) GetVideo(ctx context.Context) ([]byte, error) {
	url, err := c.GetVideoURL(ctx)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, ErrStatusNotOK
	}
	defer resp.Body.Close()

	dat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

// SaveVideoToTemp saves video to Client's TempDir.
// Returns filename and error
func (c *Client) SaveVideoToTemp(dat []byte) (string, error) {
	// в качестве имени видео будет использоваться первые шесть символов хеша
	hash := md5.Sum(dat)
	filename := fmt.Sprintf(
		"%s/%s.webm",
		c.TempDir,
		hex.EncodeToString(hash[:])[:6],
	)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(filename), 0770)
		if err != nil {
			return "", err
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = f.Write(dat)
	if err != nil {
		return "", err
	}
	return filename, nil
}
