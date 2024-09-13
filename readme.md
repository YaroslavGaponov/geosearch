Geosearch
=====


# Example (random seed)

Basic application for using the geosearch library with random seed

## Playground

[Example](https://go.dev/play/p/rnrD9uTs-W7)

## Run

```sh
make simple
```

## Code

```go
package main

import (
	"fmt"
	"github.com/YaroslavGaponov/geosearch"
)

const (
	attemps = 5
	distanceBetweenNeighbours = 2000
)

func main() {
	gs := geosearch.GeoSearchRandNew(attemps, distanceBetweenNeighbours)
	gs.AddObject(&geosearch.Object{Id: "Paris", Point: geosearch.Point{Latitude: 48.858374, Longitude: 2.336046}})
	gs.AddObject(&geosearch.Object{Id: "Berlin", Point: geosearch.Point{Latitude: 52.518430, Longitude: 13.370478}})
	gs.AddObject(&geosearch.Object{Id: "Rome", Point: geosearch.Point{Latitude: 41.898199, Longitude: 12.511268}})
	gs.AddObject(&geosearch.Object{Id: "Praha", Point: geosearch.Point{Latitude: 50.092603, Longitude: 14.444329}})
	gs.AddObject(&geosearch.Object{Id: "Poznan", Point: geosearch.Point{Latitude: 52.426060, Longitude: 16.914685}})

	me := geosearch.Point{Latitude: 52.308104, Longitude: 16.416461}

	result1 := gs.Search(me)
	fmt.Printf("Object %s, distance %.2f km, took %v\n",result1.Object.Id, result1.Distance, result1.Took)
}

```

## Result

```sh
Object Poznan, distance 36.28 km, took 116.413µs
```


# Example (fast seed)

Basic application for using the geosearch library with fast seed

## Playground

[Example](https://go.dev/play/p/UanpJoaGL4b)

## Run

```sh
make simple2
```

## Code

```go
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

	me := geosearch.Point{Latitude: 52.308104, Longitude: 16.416461}

	result1 := gs.Search(me)
	fmt.Printf("Object %s, distance %.2f km, took %v\n",result1.Object.Id, result1.Distance, result1.Took)
}
```

## Result

```sh
Object Poznan, distance 36.28 km, took 2.005µs
```
