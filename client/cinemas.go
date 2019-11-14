package client

import (
	"encoding/json"
	"fmt"
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

type Attribute struct {
	ShortTitle string
	Network    Network
	Location   Location
	Distance   string
	Mall       string
	IsFave     string
	IsSale     string
	Address    string
	Tittle     string
	//"subway": [],
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

type Response struct {
	Data []Cinema
}

type CinemasClient struct{}

func (c *CinemasClient) Get(cityId int32) []Cinema {
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

	response := Response{}
	if json.Unmarshal(body, &response) != nil {
		log.Fatalln(err)
	}

	return response.Data
}
