package main

import (
	"fmt"

	t38c "github.com/axvq/tile38-client"
)

func main() {
	client, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	if err := client.Keys.Set("fleet", "truck1").Point(33.5123, -112.2693).Do(); err != nil {
		panic(err)
	}

	if err := client.Keys.Set("fleet", "truck2").Point(33.4626, -112.1695).
		// optional params
		Field("speed", 20).
		Expiration(20).
		Do(); err != nil {
		panic(err)
	}

	// search 6 kilometers around a point. returns one truck.
	response, err := client.Search.Nearby("fleet", 33.462, -112.268, 6000).
		Where("speed", 0, 100).
		Match("truck*").
		Format(t38c.FormatPoints).Do()
	if err != nil {
		panic(err)
	}

	// truck1 {33.5123 -112.2693}
	fmt.Println(response.Points[0].ID, response.Points[0].Point)
}
