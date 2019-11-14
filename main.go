package main

import (
	"fmt"
	"playground/client"
)

func main() {
	mClient := client.MetroClient{}

	fmt.Println(mClient.Get(1))

	cClient := client.CinemasClient{}

	fmt.Println(cClient.Get(1))
}
