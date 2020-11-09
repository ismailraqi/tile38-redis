package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	t38c "github.com/axvq/tile38-client"
)

func main() {
	tile38, err := t38c.New("localhost:9851", t38c.Debug)
	if err != nil {
		log.Fatal(err)
	}
	defer tile38.Close()

	handler := func(event *t38c.GeofenceEvent) {
		b, _ := json.Marshal(event)
		fmt.Printf("event: %s\n", b)
	}
	// add a couple of points named 'truck1' and 'truck2' to a collection named 'fleet'.
	if err := tile38.Keys.Set("fleet", "truck1").Point(33.5123, -112.2693).Do(); err != nil {
		log.Fatal(err)
	}

	if err := tile38.Keys.Set("fleet", "truck2").Point(33.4626, -112.1695).Do(); err != nil {
		log.Fatal(err)
	}
	if err := tile38.Geofence.Nearby("fleet", 33.462, -112.268, 6000).
		Actions(t38c.Enter, t38c.Exit).
		Do(context.Background(), handler); err != nil {
		log.Fatal(err)
	}
}
