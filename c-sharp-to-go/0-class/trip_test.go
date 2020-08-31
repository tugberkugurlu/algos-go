package main

import (
	"github.com/deliveroo/assert-go"
	"testing"
	"time"
)

func TestTrip_GetClosestVisitedLocations(t *testing.T) {
	t.Run("Should return closest locations as expected", func(t *testing.T) {
		trip := Trip{
			Name:      "Barcelona 2020",
			TripStart: func() time.Time { val, _ := time.Parse("2006-Jan-02", "2020-Sep-05"); return val }(),
			TripEnd: func() time.Time { val, _ := time.Parse("2006-Jan-02", "2020-Sep-10"); return val }(),
			Members: map[string]bool{
				"Alice": true,
				"Bob": true,
				"Steve": true,
			},
			Locations: []Location{
				{
					Name: "Park Güell",
					Latitude: 41.4145,
					Longitude: 2.1527,
				},
				{
					Name: "La Rambla",
					Latitude: 41.380775,
					Longitude: 2.173661,
				},
				{
					Name: "Camp Nou",
					Latitude: 41.3809,
					Longitude: 2.1228,
				},
			},
		}

		target := Location{
			Name: "Josep Tarradellas Barcelona-El Prat Airport",
			Latitude: 41.2974,
			Longitude: 2.0833,
		}

		closestLocs := trip.GetClosestVisitedLocations(target)

		assert.Equal(t, func() []Location {
			result := make([]Location, len(closestLocs))
			for i, v := range closestLocs {
				result[i] = v.Destination
			}
			return result
		}(), []Location{
			{
				Name: "Camp Nou",
				Latitude: 41.3809,
				Longitude: 2.1228,
			},
			{
				Name: "La Rambla",
				Latitude: 41.380775,
				Longitude: 2.173661,
			},
			{
				Name: "Park Güell",
				Latitude: 41.4145,
				Longitude: 2.1527,
			},
		})
	})
}

func TestTrip_Length(t *testing.T) {
	t.Run("Should return the length correctly", func(t *testing.T) {
		trip := Trip{
			Name:      "Barcelona 2020",
			TripStart: func() time.Time { val, _ := time.Parse("2006-Jan-02", "2020-Sep-05"); return val }(),
			TripEnd: func() time.Time { val, _ := time.Parse("2006-Jan-02", "2020-Sep-10"); return val }(),
			Members: map[string]bool{
				"Alice": true,
				"Bob": true,
				"Steve": true,
			},
			Locations: []Location{
				{
					Name: "Park Güell",
					Latitude: 41.4145,
					Longitude: 2.1527,
				},
			},
		}

		length := trip.Length()

		expected, err := time.ParseDuration("120h")
		assert.Must(t, err)
		assert.Equal(t, length, expected)
	})
}
