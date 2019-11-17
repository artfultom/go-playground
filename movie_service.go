package main

import (
	"log"
	"playground/client/kinohod"
)

func GetPopularMovies() []Movie {
	movies, err := kinohod.GetMovies()
	if err != nil {
		log.Fatal(err)
	}

	var result []Movie

	for index, movie := range movies {
		result[index] = Movie{
			Id:     movie.Id,
			Name:   movie.Attributes.Title,
			Rating: 0,
		}
	}

	return result
}

func Test() {
	log.Println('1')
}
