package main

import "fmt"


type Player struct {
	Name string
	R int
	HR int
	RBI int
	SB int
	Hits float64
	AtBats float64
}



func main() {

	roster := []Player{
		{Name: "Ivan Herrera", R: 11, HR: 4, RBI: 13, SB: 0, Hits: 23.0, AtBats: 99.0},
		{Name: "Raphael Devers", R: 8, HR: 2, RBI: 10, SB: 0, Hits: 24.0, AtBats: 111.0},
		{Name: "Luke Keaschall", R: 12, HR: 1, RBI: 13, SB: 6, Hits: 24.0, AtBats: 110.0},
		{Name: "Kazuma Okamoto", R: 12, HR: 5, RBI: 11, SB: 0, Hits: 22.0, AtBats: 99.0},
		{Name: "Bo Bichette", R: 12, HR: 1, RBI: 12, SB: 1, Hits: 27.0, AtBats: 116.0},
		{Name: "James Wood", R: 27, HR: 10, RBI: 21, SB: 5, Hits: 27.0, AtBats: 111.0},
		{Name: "Julio Rodriguez", R: 12, HR: 2, RBI: 10, SB: 3, Hits: 28.0, AtBats: 116.0},
		{Name: "Konnor Griffin", R: 8, HR: 1, RBI: 12, SB: 6, Hits: 17.0, AtBats: 80.0},
		{Name: "Oneil Cruz", R: 19, HR: 8, RBI: 24, SB: 10, Hits: 29.0, AtBats: 112.0},
		{Name: "Jose Caballero", R: 12, HR: 3, RBI: 11, SB: 11, Hits: 27.0, AtBats: 99.0},
		{Name: "Dalton Rushing", R: 12, HR: 7, RBI: 16, SB: 0, Hits: 15.0, AtBats: 39.0},
		{Name: "Jeremiah Jackson", R: 9, HR: 5, RBI: 19, SB: 1, Hits: 24.0, AtBats: 89.0},
	}

	team_totals := map[string]float64{
		"R":		0,
		"HR":		0,
		"RBI":		0,
		"SB":		0,
		"Hits":		0,
		"AtBats":	0,
	}


	for _, p := range roster {
		team_totals["R"] += float64(p.R)
		team_totals["HR"] += float64(p.HR)
		team_totals["RBI"] += float64(p.RBI)
		team_totals["SB"] += float64(p.SB)
		team_totals["Hits"] += float64(p.Hits)
		team_totals["AtBats"] += float64(p.AtBats)
	}

	teamAvg := team_totals["Hits"] / team_totals["AtBats"]


	fmt.Printf("Team Totals: R:%.0f, HR:%.0f, RBI:%.0f, SB:%.0f\n", team_totals["R"], team_totals["HR"], team_totals["RBI"], team_totals["SB"])
	fmt.Printf("Team Batting Average: %.3f\n", teamAvg)



	// 1. Define a potential Free Agent
	freeAgents := []Player{
		{Name: "Luis Arraez", Hits: 40, AtBats: 120, R: 15, HR: 1, RBI: 10, SB: 0},
		{Name: "Steven Kwan", Hits: 35, AtBats: 110, R: 18, HR: 2, RBI: 8, SB: 5},
		{Name: "Giancarlo Stanton", Hits: 15, AtBats: 80, R: 10, HR: 12, RBI: 25, SB: 0},
	}
	//2. Calculate "What If" Team Totals
	// We take our current totals and subtract Bo's stats, then add Arraez's
	// (Bo Bichette was roster[4] in my list)

	bo := roster[4]

	fmt.Printf("\n--- Swap Analysis for %s ---\n", bo.Name)


	for _, fa := range freeAgents {

	newTotalHits := team_totals["Hits"] - bo.Hits + fa.Hits
	newTotalAtBats := team_totals["AtBats"] - bo.AtBats + fa.AtBats
	newAvg := newTotalHits / newTotalAtBats

	avgChange := newAvg - teamAvg

	fmt.Printf("Pick up %s: New Team AVG: %.3f (Change: %+.3f)\n", fa.Name, newAvg, avgChange)

	}
}
