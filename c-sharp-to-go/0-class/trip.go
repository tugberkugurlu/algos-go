package main

import (
	"sort"
	"time"

	"github.com/umahmood/haversine"
)

type Location struct {
	Name      string
	Latitude  float64
	Longitude float64
}

type LocationWithDistance struct {
	Destination  Location
	DistanceInKM float64
}

type Trip struct {
	Name      string
	TripStart time.Time
	TripEnd   time.Time
	Members   map[string]bool
	Locations []Location
}

func (t *Trip) Length() time.Duration {
	return t.TripEnd.Sub(t.TripStart)
}

func (t *Trip) IsDuringTrip(at time.Time) bool {
	return t.TripStart.Before(at) && t.TripEnd.After(at)
}

func (t *Trip) IsMemberInTrip(name string) bool {
	return t.Members[name]
}

func (t *Trip) GetClosestVisitedLocations(target Location) []LocationWithDistance {
	var distances []LocationWithDistance
	for _, loc := range t.Locations {
		lh := haversine.Coord{Lat: loc.Latitude, Lon: loc.Longitude}
		th := haversine.Coord{Lat: target.Latitude, Lon: target.Longitude}
		_, km := haversine.Distance(lh, th)
		distances = append(distances, LocationWithDistance{
			Destination:  loc,
			DistanceInKM: km,
		})
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].DistanceInKM < distances[j].DistanceInKM
	})
	return distances
}