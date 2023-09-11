package service

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/MarluxGitHub/instagram/pkg/domain/rss/domain/models"
)

type RssReader struct {}

func (reader *RssReader) Read(provider *models.Provider) (models.Rss) {
	// http.Get() returns a *http.Response
	response, err := http.Get(provider.URI)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// read the response body
	// io.ReadAll() returns a []byte
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// parse the xml
	// xml.Unmarshal() takes []byte and a pointer to a struct
	var rss models.Rss
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return rss
}
