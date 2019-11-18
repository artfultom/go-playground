package imdb

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type rating struct {
	Source string
	Value  string
}

type response struct {
	Title   string
	Ratings []rating
}

type Client struct{}

func (c *Client) Get(imdbId int) response {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	resp, err := client.Get(fmt.Sprintf("http://www.omdbapi.com/?i=tt%d&apikey=???", imdbId))
	if err != nil {
		log.Fatalln(err)
	}

	defer Close(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
