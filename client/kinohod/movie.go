package kinohod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"playground/client"
)

//type Network struct {
//	Id    string
//	Alias string
//}

//type Location struct {
//	Longitude string
//	Latitude  string
//}

//type subway struct {
//	Id       string
//	Name     string
//	Color    string
//	Distance string
//}

type attribute struct {
	Title  string
	ImdbId string
	//Network    Network
	//Location   Location
	//Distance   string
	//Mall       string
	//IsFave     string
	//IsSale     string
	//Address    string
	//Subway     []Subway
	//"goodies": [],
	//"city": {...},
	//"isAdv": "0",
	//"labels": [],
}

type movie struct {
	Id         string
	Attributes attribute
	//Link       string
}

type response struct {
	Data []movie
}

func GetMovies() ([]movie, error) {
	httpClient := client.NewHttpClient()

	resp, err := httpClient.Get(fmt.Sprintf("https://api.kinohod.ru/api/restful/v1/movies"))
	if err != nil {
		return nil, err
	} else {
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
