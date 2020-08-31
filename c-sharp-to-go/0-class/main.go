package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/hako/durafmt"
)

func main() {
	trip := Trip{
		Name:      "Barcelona 2020",
		TripStart: func() time.Time { val, _ := time.Parse("2006-Jan-02", "2020-Sep-05"); return val }(),
		TripEnd:   func() time.Time { val, _ := time.Parse("2006-Jan-02", "2020-Sep-10"); return val }(),
		Members: map[string]bool{
			"Alice": true,
			"Bob":   true,
			"Steve": true,
		},
		Locations: []Location{
			{
				Name:      "Park Güell",
				Latitude:  41.4145,
				Longitude: 2.1527,
			},
			{
				Name:      "La Rambla",
				Latitude:  41.380775,
				Longitude: 2.173661,
			},
			{
				Name:      "Camp Nou",
				Latitude:  41.3809,
				Longitude: 2.1228,
			},
		},
	}

	target := Location{
		Name:      "Josep Tarradellas Barcelona-El Prat Airport",
		Latitude:  41.2974,
		Longitude: 2.0833,
	}

	closestLocs := trip.GetClosestVisitedLocations(target)
	length := trip.Length()
	xmasTime, _ := time.Parse("2006-Jan-02", "2020-Dec-25")
	isDuringXMasTime := trip.IsDuringTrip(xmasTime)
	membersToCheck := []string{"Bob", "Barbara"}

	printTrip(trip, length, isDuringXMasTime, membersToCheck, target, closestLocs)

	trip.Members["Barbara"] = true
	printTrip(trip, length, isDuringXMasTime, membersToCheck, target, closestLocs)
}

func printTrip(trip Trip, length time.Duration, isDuringXMasTime bool, membersToCheck []string, target Location, closestLocs []LocationWithDistance) {
	fmt.Println("====================")
	fmt.Printf("Trip '%s' will start at %s and end at %s, and it will last %s.\n",
		trip.Name,
		trip.TripStart.Format("2006-Jan-02"),
		trip.TripEnd.Format("2006-Jan-02"),
		durafmt.ParseShort(length))
	if isDuringXMasTime {
		fmt.Println("The trip is going to be during christmas time.")
	} else {
		fmt.Println("The trip is not going to be during christmas time. ")
	}
	for _, memberToCheck := range membersToCheck {
		printIsInTheTrip(memberToCheck, trip.IsMemberInTrip(memberToCheck))
	}
	fmt.Printf("The trip will start at '%s', and the following locations will be visited in order:\n", target.Name)
	fmt.Printf("%s\n", func() string {
		result := make([]string, len(closestLocs))
		for i, v := range closestLocs {
			result[i] = v.Destination.Name
		}
		return strings.Join(result, ", ")
	}())
	fmt.Println("====================")
}

func printIsInTheTrip(name string, isInTrip bool) {
	if isInTrip {
		fmt.Printf("%s joins this trip.\n", name)
	} else {
		fmt.Printf("%s will not join this trip.\n", name)
	}
}

/*

 - this shows how a value update may not propagate as expected https://play.golang.org/p/9X6AgO2CRaW
 - it doesn't matter whether the value is a pointer here, change won't have an impact: https://play.golang.org/p/Q-2y0uIRlMs
	- as seen here, even if the method signatures accept the pointer, they can still be called on the value type itself. This is super unintuitive.
 	- you can make this work by assigning the value to be updated in a nested struct, which is managed by a pointer

 - you can make this safer by ensuring that your struct is not allowed to be constructed as a value type: https://play.golang.org/p/bd-2dI1_eEx
 - note that this is still safe to pass around as it cannot be modified: https://play.golang.org/p/Y8LWzsi7jzG

 - example: Trip
    - Trip name, trip start and end time, trip members, locations to visit
    - methods:
		- Members() []string
		- Locations() []Location
		- IsDuringTrip(time time.Time) bool
		- AddMember(string name) error
		- RemoveMember(string name) error
		- IsMemberInTrip(string name) bool
		- AddLocation(loc Location) error
		- GetClosestVisitedLocations(loc Location) ([]LocationWithDistance)

	- LocationWithDistance:
		- location Location, DistanceInKM float64

	- Location:
		- name string, lat, lon float64

	--------------------------------------------------------
	### Improvements
	--------------------------------------------------------

	The slices returned from these members can be modified (e.g. members in the indcies can be mutated)
	see https://play.golang.org/p/a1mFqYxKyps
	Wouldn't this be cool? https://github.com/golang/go/issues/20443

		- Members() []string
		- Locations() []Location

	Instead of directly returning the slice, we can encapsulate it for read-only purposes, and hide the source behind
		an interface:

		- ReadOnlyLocations
			- Each(func(loc Location))
			- GetByIndex(int index) (Location, error)
			- Len() int

		- ReadOnlyMembers
			- Each(func(loc string))
			- GetByIndex(int index) (string, error)
			- Len() int

	For the below one, we don't have to abstract this away

		- GetClosestVisitedLocations(loc Location) ([]LocationWithDistance)

	--------------------------------------------------------

 - construction, and validation
 - methods

 - construction, and validation
 - methods
 - private member hiding
 - prevent uncontrolled mutability
 - pass by reference, not by value
 - static construction

 - immutability

*/
