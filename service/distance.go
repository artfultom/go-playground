package service

import (
	"github.com/RyanCarrier/dijkstra"
	"log"
	"playground/client/yandex"
)

func GetPath(source int, destination int) (int64, []int) {
	stations, _ := yandex.GetStations()

	graph := dijkstra.NewGraph()

	for _, station := range stations.Stations {
		graph.AddVertex(station.LabelId)
	}

	for _, link := range stations.Links {
		graph.AddArc(link.FromStationId, link.ToStationId, link.WeightTime)
	}

	best, err := graph.Shortest(source, destination)
	if err != nil {
		log.Fatal(err)
	}

	return best.Distance / 60, best.Path
}
