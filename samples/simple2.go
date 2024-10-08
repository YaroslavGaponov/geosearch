package main

import (
	"fmt"
	"github.com/YaroslavGaponov/geosearch"
)

const (
	distanceBetweenNeighbours = 2000
)

func main() {
	gs := geosearch.GeoSearchFastNew(distanceBetweenNeighbours)
	gs.AddObject(&geosearch.Object{Id: "Paris", Point: geosearch.Point{Latitude: 48.858374, Longitude: 2.336046}})
	gs.AddObject(&geosearch.Object{Id: "Berlin", Point: geosearch.Point{Latitude: 52.518430, Longitude: 13.370478}})
	gs.AddObject(&geosearch.Object{Id: "Rome", Point: geosearch.Point{Latitude: 41.898199, Longitude: 12.511268}})
	gs.AddObject(&geosearch.Object{Id: "Praha", Point: geosearch.Point{Latitude: 50.092603, Longitude: 14.444329}})
	gs.AddObject(&geosearch.Object{Id: "Poznan", Point: geosearch.Point{Latitude: 52.426060, Longitude: 16.914685}})

	me := geosearch.Point{Latitude: 53.308104, Longitude: 14.416461}

	result1 := gs.Search(me)
	fmt.Printf("Object %s, distance %.2f km, took %v\n",result1.Object.Id, result1.Distance, result1.Took)
}
