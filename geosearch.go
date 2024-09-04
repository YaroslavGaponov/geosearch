package geosearch

import (
	"math/rand"
	"sort"
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
	objects     []*Object
	maxDistance float64
}

func New(length int, maxDistance float64) GeoSearch {
	return GeoSearch{
		objects:     make([]*Object, 0, length),
		maxDistance: maxDistance,
	}
}

func (gs *GeoSearch) AddObject(obj *Object) {
	if len(gs.objects) > 0 {
		for _, obj2 := range gs.objects {
			dist := Haversine(obj.Point, obj2.Point)
			if dist <= gs.maxDistance {
				obj.Neighbours = append(obj.Neighbours, obj2)
				obj2.Neighbours = append(obj2.Neighbours, obj)
			}
		}
	}
	gs.objects = append(gs.objects, obj)
}

func (gs *GeoSearch) Search(point Point, attempt int) []Result {

	result := make([]Result, 0)
	skipped := make(map[string]struct{})

	for i := 0; i < attempt; i++ {

		idx := rand.Intn(len(gs.objects))

		list := []*Object{gs.objects[idx]}

		for len(list) > 0 {
			list2 := []*Object{}
			for _, obj := range list {
				if _, found := skipped[obj.Id]; !found {
					distance := Haversine(point, obj.Point)
					if distance <= gs.maxDistance {
						result = append(result, Result{Object: obj, Distance: distance})
						list2 = append(list2, obj.Neighbours...)
					}
					skipped[obj.Id] = struct{}{}
				}
			}
			list = list2
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Distance < result[j].Distance
	})
	return result
}
