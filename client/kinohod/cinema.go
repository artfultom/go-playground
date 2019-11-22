package kinohod

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type CinemasData struct {
	Id         string
	Link       string
	Structure  string
	Attributes struct {
		ShortTitle string
		Network    struct {
			Id    string
			Alias string
		}
		Location struct {
			Longitude string
			Latitude  string
		}
		Distance string
		Mall     string
		IsFave   string
		IsSale   string
		Address  string
		Title    string
		Subway   []struct {
			Id       string
			Name     string
			Color    string
			Distance string
		}
	}
}

type cinemasResponse struct {
	Data []CinemasData
}

func GetCinemas(cityId int) ([]CinemasData, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSHandshakeTimeout: 5 * time.Second,
		},
	}

	resp, err := client.Get(fmt.Sprintf("https://api.kinohod.ru/api/restful/v1/cinemas?city=%d", cityId))
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

	response := cinemasResponse{}
	err = xml.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}
