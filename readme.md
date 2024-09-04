Geosearch
=====


# Example

## Code

```go
package main

import (
	"fmt"
	"github.com/YaroslavGaponov/geosearch"
)


func main() {
	gs := geosearch.New()
	gs.AddObject(&geosearch.Object{Id: "Paris", Point: geosearch.Point{Latitude: 48.858374, Longitude: 2.336046}})
	gs.AddObject(&geosearch.Object{Id: "Berlin", Point: geosearch.Point{Latitude: 52.518430, Longitude: 13.370478}})
	gs.AddObject(&geosearch.Object{Id: "Rome", Point: geosearch.Point{Latitude: 41.898199, Longitude: 12.511268}})
	gs.AddObject(&geosearch.Object{Id: "Praha", Point: geosearch.Point{Latitude: 50.092603, Longitude: 14.444329}})
	gs.AddObject(&geosearch.Object{Id: "Poznan", Point: geosearch.Point{Latitude: 52.426060, Longitude: 16.914685}})

	me := geosearch.Point{Latitude: 52.308104, Longitude: 16.416461}
    
	result := gs.Search(me)
	fmt.Println(result.Object.Id, result.Distance, "km")
}

```

## Result

```sh
Poznan 36.28095488081731 km
```