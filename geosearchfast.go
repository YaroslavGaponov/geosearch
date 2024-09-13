package geosearch

import (
	"time"
)

type GeoSearchFast struct {
	GeoSearch
	seed map[int]*Object
}

func GeoSearchFastNew(distanceBetweenNeighbours float64) GeoSearchFast {
	return GeoSearchFast{
		GeoSearch: GeoSearchNew(distanceBetweenNeighbours),
		seed:      make(map[int]*Object),
	}
}

func (gs *GeoSearchFast) AddObject(obj *Object) {
	gs.GeoSearch.AddObject(obj)

	a := int(obj.Latitude) + 90
	b := int(obj.Longitude) + 180
	idx := a*1000 + b
	gs.seed[idx] = obj
}

func (gs *GeoSearchFast) Search(point Point) Result {
	start := time.Now()

	a := int(point.Latitude) + 90
	b := int(point.Longitude) + 180
	idx := a*1000 + b

	dir := 1
	var result Result
	for {
		seed, found := gs.seed[idx]
		if found {
			result = gs.GeoSearch.Search(point, seed)
			break
		} else {
			idx += dir
			dir = -(dir + 1)
		}
	}
	result.Took = time.Since(start)
	return result
}
