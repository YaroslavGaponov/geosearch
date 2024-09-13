package geosearch

import (
	"time"
)

type Point struct {
	Latitude  float64
	Longitude float64
}

type Object struct {
	Id string
	Point
	Neighbours []*Object
}

type Result struct {
	Object   *Object
	Distance float64
	Took     time.Duration
}

type GeoSearch struct {
	objects                   []*Object
	distanceBetweenNeighbours float64
}

func GeoSearchNew(distanceBetweenNeighbours float64) GeoSearch {
	return GeoSearch{
		objects:                   make([]*Object, 0),
		distanceBetweenNeighbours: distanceBetweenNeighbours,
	}
}

func (gs *GeoSearch) AddObject(obj *Object) {
	if len(gs.objects) > 0 {
		for i := 0; i < len(gs.objects); i++ {
			distance := Haversine(obj.Point, gs.objects[i].Point)
			if distance <= gs.distanceBetweenNeighbours {
				obj.Neighbours = append(obj.Neighbours, gs.objects[i])
				gs.objects[i].Neighbours = append(gs.objects[i].Neighbours, obj)
			}
		}
	}
	gs.objects = append(gs.objects, obj)
}

func (gs *GeoSearch) Search(point Point, obj *Object) Result {

	start := time.Now()

	distance := Haversine(point, obj.Point)

	for {
		found := false
		for _, obj2 := range obj.Neighbours {
			distance2 := Haversine(point, obj2.Point)
			if distance2 < distance {
				obj = obj2
				distance = distance2
				found = true
			}
		}
		if !found {
			break
		}

	}
	return Result{Object: obj, Distance: distance, Took: time.Since(start)}
}
