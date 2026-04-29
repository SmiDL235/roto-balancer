package main


type Player struct {
	Name		string
	Positions []	string  // "C", "1B', "2B", "3B", "SS", "OF", "UTIL", "SP", "RP"
	IsPitcher 	bool // true for pitchers, false for hitters


	// Hitter Stats
	R, HR, RBI, SB 	float64
	Hits, AtBats 	float64

	// Pitcher Stats
	W, SO, SV 	float64
	ER, IP	 	float64
	HitsAllowed 	float64 //Hits given up
	BB		float64 // Walks
}


// GetTotalOuts converts IP like 13.2 into 41 total outs
func (p Player) GetTotalOuts() float64 {
	innings := float64(int(p.IP))
	outs := (p.IP - innings + 0.001) * 10
	return (innings * 3) + float64(int(outs))
}

// GetWHIP calculates (Hits + BB) / IP
func (p Player) GetWHIP() float64 {
	trueInnings := p.GetTotalOuts() / 3.0

	if trueInnings == 0 {
		return 0
	}
	return (p.HitsAllowed + p.BB) / trueInnings
}

// GetERA calculates (ER / IP) * 9
func (p Player) GetERA() float64 {
	trueInnings := p.GetTotalOuts() / 3.0

	if trueInnings == 0 {
		return 0
	}
	return (p.ER / trueInnings) * 9 
}


func (p Player) GetAVG() float64 {
	if p.AtBats == 0 {
		return 0
	}
	return p.Hits / p.AtBats
}
