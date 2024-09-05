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
	distanceBetweenNeighbours = 2000
)

func main() {
	gs := geosearch.New(distanceBetweenNeighbours)
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

	gs := geosearch.New(distanceBetweenNeighbours)

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

	fmt.Println("Result:")
	for i := 0; i < 5; i++ {
		result := gs.Search(me)
		fmt.Println("attempt", i, result.Object.Id, result.Distance, "km")
	}
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
Result:
attempt 0 Opalenica 0.19742479602703403 km
attempt 1 Opalenica 0.19742479602703403 km
attempt 2 Opalenica 0.19742479602703403 km
attempt 3 Mount Pearl Park 4793.351425097892 km
attempt 4 Opalenica 0.19742479602703403 km
```