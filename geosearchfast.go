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

	lat := int(obj.Latitude) + 90
	lng := int(obj.Longitude) + 180
	idx := lat*1000 + lng
	gs.seed[idx] = obj
}

func (gs *GeoSearchFast) Search(point Point) Result {
	start := time.Now()

	direct := [...][2]int{{0, 0}, {0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, -1}}

	lat := int(point.Latitude) + 90
	lng := int(point.Longitude) + 180

	dir := 0
	distance := 1
	var result Result
	for {
		nextLat := lat + distance*direct[dir][0]
		nextLng := lng + distance*direct[dir][1]
		idx := nextLat*1000 + nextLng
		seed, found := gs.seed[idx]
		if found {
			result = gs.GeoSearch.Search(point, seed)
			break
		}
		dir++
		if dir >= len(direct) {
			dir = 0
			distance++
		}
	}
	result.Took = time.Since(start)
	return result
}
