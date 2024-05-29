package main

import (
	"math"
	"time"

	"github.com/Admiral-Simo/toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(*types.OBUData) (float64, error)
}

type CalculatorService struct {
	previousCoordinates map[int]*types.OBUData // Map to store previous coordinates for each OBU
	distances           map[int]float64        // Map to store distances for each OBU
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{
		previousCoordinates: make(map[int]*types.OBUData),
		distances:           make(map[int]float64),
	}
}

func (s *CalculatorService) CalculateDistance(data *types.OBUData) (float64, error) {
	var distance float64

	if prevCoords, exists := s.previousCoordinates[data.OBUID]; exists {
		distance = calculateDistanceHelper(data.Lat, data.Long, prevCoords.Lat, prevCoords.Long)
		s.distances[data.OBUID] += distance
	}
	s.previousCoordinates[data.OBUID] = data

	distanceData := &types.Distance{
		OBUID: data.OBUID,
		Value: distance,
		Unix:  time.Now().Unix(),
	}

	// here should be a POST request to the aggregator microservice
	// POST(distanceData)
	_ = distanceData

	return s.distances[data.OBUID], nil
}

// calculateDistanceHelper calculates the great-circle distance between two points
// given their latitude and longitude using the Haversine formula.
func calculateDistanceHelper(lat1, long1, lat2, long2 float64) float64 {
	const earthRadius = 6371 // Earth's radius in kilometers

	// Convert degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	long1Rad := long1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	long2Rad := long2 * math.Pi / 180

	// Haversine formula
	dlat := lat2Rad - lat1Rad
	dlong := long2Rad - long1Rad

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dlong/2)*math.Sin(dlong/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Distance in kilometers
	distance := earthRadius * c

	return distance
}
