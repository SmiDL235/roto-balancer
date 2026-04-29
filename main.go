
package main

import "fmt"






func main() {

	roster := GetMyRoster()

	team_totals := map[string]float64{
		"R":		0,
		"HR":		0,
		"RBI":		0,
		"SB":		0,
		"Hits":		0,
		"AtBats":	0,
		"IP":		0,
		"HitsAllowed":	0,
		"BB":		0,
		"ER":		0,
		"SO":		0,
		"W":		0,
		"SV":		0,
	}


	for _, p := range roster {
		if p.IsPitcher {
			team_totals["TotalOuts"] += p.GetTotalOuts()
			team_totals["IP"] += p.IP
			team_totals["HitsAllowed"] += p.HitsAllowed
			team_totals["BB"] += p.BB
			team_totals["ER"] += p.ER
			team_totals["SO"] += p.SO
			team_totals["W"] += p.W
			team_totals["SV"] += p.SV
		} else {
			team_totals["R"] += (p.R)
			team_totals["HR"] += (p.HR)
			team_totals["RBI"] += (p.RBI)
			team_totals["SB"] += (p.SB)
			team_totals["Hits"] += (p.Hits)
			team_totals["AtBats"] += (p.AtBats)
		}
	}

	teamAvg := 0.0
	if team_totals["AtBats"] > 0 {
		teamAvg = team_totals["Hits"] / team_totals["AtBats"]
	}


	teamTotalOuts := team_totals["TotalOuts"]
	trueInnings := teamTotalOuts / 3.0


	teamERA := 0.0
	teamWHIP := 0.0

	if trueInnings > 0 {
		// Formula: (Earned Runs / True Innings) * 9
		teamERA = (team_totals["ER"] / trueInnings) * 9

		// Formula: ( Hits + Walks) / True Innings 
		teamWHIP = (team_totals["HitsAllowed"] + team_totals["BB"]) / trueInnings
	}

	fmt.Printf("--- Hitting Totals ---")
	fmt.Printf("R: %.0f | HR: %.0f | RBI: %.0f | SB: %.0f | AVG: %.3f\n", team_totals["R"], team_totals["HR"], team_totals["RBI"], team_totals["SB"], teamAvg)


	// Convert outs back to standard .1 / .2 notation
	displayIP := float64(int(team_totals["TotalOuts"]) / 3) + float64(int(team_totals["TotalOuts"]) % 3) / 10.0

	fmt.Printf("\n--- Pitching Totals ---\n")
	fmt.Printf("W: %.0f | SO: %.0f | SV: %.0f | Total IP: %.1f | ERA: %.2f | WHIP: %.3f\n", team_totals["W"], team_totals["SO"], team_totals["SV"], displayIP, teamERA, teamWHIP)



	// 1. Define a potential Free Agent
	freeAgents := []Player{
		{Name: "Luis Arraez", Positions: []string{"1B", "2B"}, IsPitcher: false,  Hits: 40, AtBats: 120, R: 15, HR: 1, RBI: 10, SB: 0},
		{Name: "Steven Kwan", Positions: []string{"OF"}, IsPitcher: false, Hits: 35, AtBats: 110, R: 18, HR: 2, RBI: 8, SB: 5},
		{Name: "Giancarlo Stanton", Positions: []string{"OF"}, IsPitcher: false, Hits: 15, AtBats: 80, R: 10, HR: 12, RBI: 25, SB: 0},
		{Name: "Josh Hader", Positions: []string{"RP"}, IsPitcher: true, IP: 1.0, ER: 0, HitsAllowed: 0, BB: 0, SO: 2, SV: 1, W: 0},
		{Name: "Logan Webb", Positions: []string{"SP"}, IsPitcher: true, IP: 7, ER: 2, HitsAllowed: 5, BB: 1, SO: 8, W: 1},
	}

	bo := roster[4]


	for _, fa := range freeAgents {
		if fa.IsPitcher {
			// 1. Calculate New Total Outs and Stats
			newTotalOuts := team_totals["TotalOuts"] + fa.GetTotalOuts()
			newTrueInnings := newTotalOuts / 3.0

			newER := team_totals["ER"] + fa.ER
			newH := team_totals["HitsAllowed"] + fa.HitsAllowed
			newBB := team_totals["BB"] + fa.BB

			// 2. Calculate New Ratios
			newERA := (newER / newTrueInnings) * 9
			newWHIP := (newH + newBB) / newTrueInnings

			fmt.Printf("\n--- Pitcher Swap Analysis: %s ---\n", fa.Name)
			fmt.Printf("  -> New ERA: %.2f (Change: %+.2f)\n", newERA, newERA - teamERA)
                	fmt.Printf("  -> New WHIP: %.3f (Change: %+.3f)\n", newWHIP, newWHIP - teamWHIP)
                	fmt.Printf("  -> Gains: +%.0f SO, +%.0f SV, +%.0f W\n\n", fa.SO, fa.SV, fa.W)

		} else {
			 newTotalHits := team_totals["Hits"] - bo.Hits + fa.Hits
        		newTotalAtBats := team_totals["AtBats"] - bo.AtBats + fa.AtBats
        		newAvg := newTotalHits / newTotalAtBats

        		avgChange := newAvg - teamAvg

			fmt.Printf("\n--- Hitter Swap Analysis: %s for %s ---\n", fa.Name, bo.Name)
        		fmt.Printf(" -> New Team AVG: %.3f (Change: %+.3f)\n", newAvg, avgChange)
        		fmt.Printf(" -> New Team HR:  %.0f | SB: %.0f\n", team_totals["HR"]-bo.HR+fa.HR, team_totals["SB"]-bo.SB+fa.SB)
    
		}
	}

}
