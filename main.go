package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type AgencyInfo struct {
	strAgencyID       string
	strAgencyName     string
	strAgencyURL      string
	strAgencyTimeZone string
}
type Routes struct {
	strRouteID          string
	strAgencyID         string
	strRouteShortName   string
	strRouteLongName    string
	strRouteDescription string
	strRouteType        string
	strRouteURL         string
	strRouteColor       string
	strRouteTextColor   string
	intRouteSortOrder   int
	intContinousPickup  int
	intContinousDropoff int
}
type Stops struct {
	strStopID string
}
type Trips struct {
	strRouteID   string
	strServiceID string
	strTripId    string
	strHeadSign  string
	strShortName string
	intDirection int
	strBlockID   string
}

func main() {

	fmt.Println("Loading schedule")
	strScheduleDir := os.Args[1]

	sAgencyInfo := LoadAgencyInformation(strScheduleDir)
	sRoutes := LoadRoutes(strScheduleDir)
	fmt.Println(sRoutes)
	fmt.Println(sAgencyInfo.strAgencyName)

}

func LoadSchedule(strPath string) {

}

func ListStopsbyRoute(strRouteId string) {

}

func LoadStops(strPath string) (s []Stops) {

}
func LoadRoutes(strPath string) (r []Routes) {

	csvFile, err := os.Open(strPath + "/routes.txt")
	if err != nil {
		fmt.Println(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		//throw away the first line
		if line[0] != "route_id" {
			temp := Routes{
				strRouteID:        line[0],
				strAgencyID:       line[1],
				strRouteShortName: line[2],
				strRouteLongName:  line[3],
			}
			r = append(r, temp)

		}
	}
	return r
}
func LoadAgencyInformation(strPath string) (a AgencyInfo) {

	csvFile, err := os.Open(strPath + "/agency.txt")
	if err != nil {
		fmt.Println(err)
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		//throw away the first line
		if line[0] != "agency_id" {

			a.strAgencyID = line[0]
			a.strAgencyName = line[1]
			a.strAgencyURL = line[2]
			fmt.Println(a)
		}
	}
	return a
}
