package client

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Station struct {
	Id    string
	Name  string
	Lat   float64
	Lng   float64
	Order uint32
}

type Line struct {
	Id       string
	HexColor string `json:"hex_color"`
	Name     string
	Stations []Station
}

type City struct {
	Id    string
	Name  string
	Lines []Line
}

type MetroClient struct{}

func (c *MetroClient) Get(cityId int32) City {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	resp, err := client.Get(fmt.Sprintf("https://api.hh.ru/metro/%d", cityId))
	if err != nil {
		log.Fatalln(err)
	}

	defer Close(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	city := City{}
	if json.Unmarshal(body, &city) != nil {
		log.Fatalln(err)
	}

	return city
}

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}
