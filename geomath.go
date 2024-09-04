package geosearch

import (
	"math"
)

const EarthRadiusKm = 6371

func Haversine(p1, p2 Point) float64 {

	lat1Rad := degreesToRadians(p1.Latitude)
	lon1Rad := degreesToRadians(p1.Longitude)
	lat2Rad := degreesToRadians(p2.Latitude)
	lon2Rad := degreesToRadians(p2.Longitude)

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := EarthRadiusKm * c

	return distance
}

func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
