package geosearch

import (
	"math/rand"
	"sync"
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
	attemps                   int
	distanceBetweenNeighbours float64
}

func New(attemps int, distanceBetweenNeighbours float64) GeoSearch {
	return GeoSearch{
		objects:                   make([]*Object, 0),
		attemps:                   attemps,
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

func (gs *GeoSearch) SearchOne(point Point) Result {

	start := time.Now()

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
	return Result{Object: obj, Distance: distance, Took: time.Since(start)}
}

func (gs *GeoSearch) Search(point Point) Result {
	start := time.Now()
	result := gs.SearchOne(point)
	k := 1
	halfAttempts := gs.attemps >> 1
	for i := 1; i < gs.attemps; i++ {
		result2 := gs.SearchOne(point)
		if result.Object.Id == result2.Object.Id {
			k++
			if k > halfAttempts {
				break
			}
		} else if result.Distance > result2.Distance {
			result = result2
			k = 1
		}
	}
	result.Took = time.Since(start)
	return result
}

func (gs *GeoSearch) Search2(point Point) Result {
	start := time.Now()
	var wg sync.WaitGroup

	results := make([]Result, gs.attemps)

	for i := 0; i < gs.attemps; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = gs.SearchOne(point)
		}(i)
	}

	wg.Wait()

	result := results[0]
	for i := 1; i < len(results); i++ {
		if result.Distance > results[i].Distance {
			result = results[i]
		}
	}

	result.Took = time.Since(start)
	return result
}
