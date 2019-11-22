package yandex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"playground/client"
)

type stationsResponse struct {
	Data string
}

type stationsData struct {
	Stations map[string]struct {
		Name      string
		LineId    int
		LabelId   int
		BoardInfo struct {
			Exit []struct {
				Pos []int
			}
		}
		Transfer []struct {
			ToSt int
			Pos  []int
		}
		LinkIds           []int
		IsTransferStation bool
	}
	Lines map[string]struct {
		Name  string
		Color string
	}
	Links map[string]struct {
		Type           string
		FromStationId  int
		ToStationId    int
		WeightTime     int
		WeightTransfer int
	}
	Transfers map[string]struct {
		StationIds []int
	}
	Labels map[string]struct {
		StationIds []int
	}
}

func GetStations() (*stationsData, error) {
	httpClient := client.NewHttpClient()

	resp, err := httpClient.Get(fmt.Sprintf("https://yandex.ru/metro/api/get-scheme-metadata?id=1&lang=ru"))
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

	response := stationsResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	data := stationsData{}
	err = json.Unmarshal([]byte(response.Data), &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
