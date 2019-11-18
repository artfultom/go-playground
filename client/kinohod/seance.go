package kinohod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SeancesData struct {
	Cinemas []struct {
		Title       string
		ShortTitle  string
		Description string
		Website     string
		Mall        string
		Address     string
		Location    struct {
			Longitude float64
			Latitude  float64
		}
		Seances []struct {
			Date       string
			MaxPrice   int
			MinPrice   int
			Formats    []string
			StartTime  string
			IsOrigin   bool
			CanBePayed bool
		}
	}
	Dates []string
}

type seancesResponse struct {
	Data SeancesData
}

func GetSeances(cityId int, movieId int) (*SeancesData, error) {
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
		return nil, err
	} else {
		defer func() {
			err := req.Body.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}

	req.Header.Add("x-requested-with", "XMLHttpRequest")
	resp, err := client.Do(req)
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

	response := seancesResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}
