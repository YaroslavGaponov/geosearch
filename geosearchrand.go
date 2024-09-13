package geosearch

import (
	"math/rand"
	"sync"
	"time"
)

type GeoSearchRand struct {
	GeoSearch
	attemps int
}

func GeoSearchRandNew(attemps int, distanceBetweenNeighbours float64) GeoSearchRand {
	return GeoSearchRand{
		GeoSearch: GeoSearchNew(distanceBetweenNeighbours),
		attemps:   attemps,
	}
}

func (gs *GeoSearchRand) Search(point Point) Result {
	start := time.Now()
	var wg sync.WaitGroup

	results := make([]Result, gs.attemps)

	for i := 0; i < gs.attemps; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			seed := gs.objects[rand.Intn(len(gs.objects))]
			results[i] = gs.GeoSearch.Search(point, seed)
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
