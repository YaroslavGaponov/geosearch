package geosearch

import (
	"math/rand"
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
}

type GeoSearch struct {
	objects []*Object
}

func New() GeoSearch {
	return GeoSearch{
		objects: make([]*Object, 0),
	}
}

func (gs *GeoSearch) AddObject(obj *Object) {
	if len(gs.objects) > 0 {
		idx := 0
		distance := Haversine(obj.Point, gs.objects[0].Point)
		for i := 1; i < len(gs.objects); i++ {
			dist := Haversine(obj.Point, gs.objects[i].Point)
			if distance > dist {
				idx = i
				distance = dist
			}
		}
		obj.Neighbours = append(obj.Neighbours, gs.objects[idx])
		gs.objects[idx].Neighbours = append(gs.objects[idx].Neighbours, obj)
	}
	gs.objects = append(gs.objects, obj)
}

func (gs *GeoSearch) Search(point Point) Result {

	idx := rand.Intn(len(gs.objects))

	obj := gs.objects[idx]
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
	return Result{Object: obj, Distance: distance}
}

