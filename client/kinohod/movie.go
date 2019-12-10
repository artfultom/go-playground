package kinohod

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"playground/client"
)

type MovieData struct {
	Id         string
	Attributes struct {
		Title          string
		ImdbId         string
		ImdbRating     string
		ProductionYear string
		AnnotationFull string
		Genres         []struct {
			Name string
		}
	}
}

type moviesResponse struct {
	Data []MovieData
}

type movieResponse struct {
	Data MovieData
}

func GetMovies() ([]MovieData, error) {
	httpClient := client.NewHttpClient()

	resp, err := httpClient.Get(fmt.Sprintf("https://api.kinohod.ru/api/restful/v1/movies?attributes[]=movie.full"))
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

	response := moviesResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

func GetMovie(id int) (*MovieData, error) {
	httpClient := client.NewHttpClient()

	resp, err := httpClient.Get(fmt.Sprintf("https://api.kinohod.ru/api/restful/v1/movies/%d?attributes[]=movie.full", id))
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

	response := movieResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response.Data, nil
}
