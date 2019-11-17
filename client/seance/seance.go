package seance

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Location struct {
	Longitude float64
	Latitude  float64
}

type Seance struct {
	Date       string
	MaxPrice   int
	MinPrice   int
	Formats    []string
	StartTime  string
	IsOrigin   bool
	CanBePayed bool

	/**
	"seatCategories":[
	     {
	        "name":"Стандартное",
	        "id":45361021,
	        "price":420
	     }
	  ],
	"hall":{
	     "name":"Зал 7",
	     "id":5958
	  },
	"isVip":false,
	"isImax":false,
	"time":"18:05",
	*/
}

type Cinema struct {
	Title       string
	ShortTitle  string
	Description string
	Website     string
	Mall        string
	Address     string
	Location    Location
	Seances     []Seance
}

type CinemaDate struct {
	Cinemas []Cinema
	Dates   []string
}

type Response struct {
	Data CinemaDate
}

type Client struct{}

func (c *Client) Get(cityId int, movieId int) CinemaDate {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	// apikey=7426d96b-9f48-3f37-a77d-705e533c4696
	// date=2019-11-12
	req, err := http.NewRequest("GET", fmt.Sprintf("https://kinohod.ru/widget/movie/cinemas?cityId=%d&movieId=%d", cityId, movieId), nil)
	if err != nil {
		log.Fatalln(err)
	}

	//defer Close(req.Body)

	req.Header.Add("x-requested-with", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	//defer Close(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := Response{}
	err = json.Unmarshal(body, &response)
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
