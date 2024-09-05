Geosearch
=====


# Example #1

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
	gs := geosearch.New(attemps, distanceBetweenNeighbours)
	gs.AddObject(&geosearch.Object{Id: "Paris", Point: geosearch.Point{Latitude: 48.858374, Longitude: 2.336046}})
	gs.AddObject(&geosearch.Object{Id: "Berlin", Point: geosearch.Point{Latitude: 52.518430, Longitude: 13.370478}})
	gs.AddObject(&geosearch.Object{Id: "Rome", Point: geosearch.Point{Latitude: 41.898199, Longitude: 12.511268}})
	gs.AddObject(&geosearch.Object{Id: "Praha", Point: geosearch.Point{Latitude: 50.092603, Longitude: 14.444329}})
	gs.AddObject(&geosearch.Object{Id: "Poznan", Point: geosearch.Point{Latitude: 52.426060, Longitude: 16.914685}})

	me := geosearch.Point{Latitude: 52.308104, Longitude: 16.416461}

	result1 := gs.Search(me)
	fmt.Printf("Object %s, distance %0.2f km, took %v\n",result1.Object.Id, result1.Distance, result1.Took)

	result2 := gs.Search2(me)
	fmt.Printf("Object %s, distance %0.2f km, took %v\n",result2.Object.Id, result2.Distance, result2.Took)

}
```

## Result

```sh
Object Poznan, distance 36.28 km, took 23.704µs
Object Poznan, distance 36.28 km, took 21.992µs
```

# Example #2

## Code

```go
package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/YaroslavGaponov/geosearch"
)

const (
	Width = 50

	ZIP_FILE = "./worldcities.zip"
	CSV_FILE = "worldcities.csv"

	attemps                   = 5
	distanceBetweenNeighbours = 400
)

func main() {

	zipFile, err := zip.OpenReader(ZIP_FILE)
	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}
	defer zipFile.Close()

	var file *zip.File
	for _, f := range zipFile.File {
		if f.Name == CSV_FILE {
			file = f
			break
		}
	}
	if file == nil {
		log.Fatalf("file %s is not found", CSV_FILE)
	}

	reader, err := file.Open()
	if err != nil {
		log.Fatalf("failed to open file %s", err)
	}
	defer reader.Close()

	gs := geosearch.New(attemps, distanceBetweenNeighbours)

	fmt.Println("Creating geo structure...")
	counter := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		city := trim(parts[1])
		lat, err := strconv.ParseFloat(trim(parts[2]), 64)
		if err != nil {
			continue
		}
		lng, err := strconv.ParseFloat(trim(parts[3]), 64)
		if err != nil {
			continue
		}
		print(city)
		gs.AddObject(&geosearch.Object{Id: city, Point: geosearch.Point{Latitude: lat, Longitude: lng}})
		counter++
	}
	printEnd("Done")
	fmt.Println(counter, "points are loaded")

	me := geosearch.Point{Latitude: 52.308104, Longitude: 16.416461}

	result := gs.Search(me)
	fmt.Printf("Object %s, distance %0.2f km, took %v\n", result.Object.Id, result.Distance, result.Took)

	result2 := gs.Search2(me)
	fmt.Printf("Object %s, distance %0.2f km, took %v\n", result2.Object.Id, result2.Distance, result2.Took)
}

func trim(s string) string {
	return s[1 : len(s)-1]
}

func print(s string) {
	fmt.Print(s, strings.Repeat(" ", Width-len(s)), "\r")
}
func printEnd(s string) {
	fmt.Println(s, strings.Repeat(" ", Width-len(s)))
}
```

## Result

```sh
Creating geo structure...
Done                                               
47868 points are loaded
Object Opalenica, distance 0.20 km, took 2.7209ms
Object Opalenica, distance 0.20 km, took 785.347µs
```