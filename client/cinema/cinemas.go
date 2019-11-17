package cinema

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Network struct {
	Id    string
	Alias string
}

type Location struct {
	Longitude string
	Latitude  string
}

type Subway struct {
	Id       string
	Name     string
	Color    string
	Distance string
}

type Attribute struct {
	ShortTitle string
	Network    Network
	Location   Location
	Distance   string
	Mall       string
	IsFave     string
	IsSale     string
	Address    string
	Title      string
	Subway     []Subway
	//"goodies": [],
	//"city": {...},
	//"isAdv": "0",
	//"labels": [],
}

type Cinema struct {
	Id         string
	Link       string
	Structure  string
	Attributes Attribute
}

type response struct {
	Data []Cinema
}

type Client struct{}

func (c *Client) Get(cityId int) []Cinema {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	resp, err := client.Get(fmt.Sprintf("https://api.kinohod.ru/api/restful/v1/cinemas?city=%d", cityId))
	if err != nil {
		log.Fatalln(err)
	}

	defer Close(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := response{}
	err = xml.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.Data
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
