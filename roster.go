package main


func GetMyRoster() []Player {
	return []Player{
                {Name: "Ivan Herrera", Positions: []string{"C"}, IsPitcher: false, R: 11, HR: 4, RBI: 13, SB: 0, Hits: 23.0, AtBats: 99.0},
                {Name: "Raphael Devers", Positions: []string{"1B"}, IsPitcher: false, R: 8, HR: 2, RBI: 10, SB: 0, Hits: 24.0, AtBats: 111.0},
                {Name: "Luke Keaschall", Positions: []string{"2B"}, IsPitcher: false, R: 12, HR: 1, RBI: 13, SB: 6, Hits: 24.0, AtBats: 110.0},
                {Name: "Kazuma Okamoto", Positions: []string{"1B", "3B"}, IsPitcher: false, R: 12, HR: 5, RBI: 11, SB: 0, Hits: 22.0, AtBats: 99.0},
                {Name: "Bo Bichette", Positions: []string{"3B", "SS"}, IsPitcher: false, R: 12, HR: 1, RBI: 12, SB: 1, Hits: 27.0, AtBats: 116.0},
                {Name: "James Wood", Positions: []string{"OF"}, IsPitcher: false, R: 27, HR: 10, RBI: 21, SB: 5, Hits: 27.0, AtBats: 111.0},
                {Name: "Julio Rodriguez", Positions: []string{"OF"}, IsPitcher: false, R: 12, HR: 2, RBI: 10, SB: 3, Hits: 28.0, AtBats: 116.0},
                {Name: "Konnor Griffin", Positions: []string{"SS", "OF"}, IsPitcher: false, R: 8, HR: 1, RBI: 12, SB: 6, Hits: 17.0, AtBats: 80.0},
                {Name: "Oneil Cruz", Positions: []string{"OF"}, IsPitcher: false, R: 19, HR: 8, RBI: 24, SB: 10, Hits: 29.0, AtBats: 112.0},
                {Name: "Jose Caballero", Positions: []string{"2B", "3B", "SS", "OF"}, IsPitcher: false, R: 12, HR: 3, RBI: 11, SB: 11, Hits: 27.0, AtBats: 99.0},
                {Name: "Dalton Rushing", Positions: []string{"C"}, IsPitcher: false, R: 12, HR: 7, RBI: 16, SB: 0, Hits: 15.0, AtBats: 39.0},
                {Name: "Jeremiah Jackson", Positions: []string{"2B", "3B", "OF"}, IsPitcher: false, R: 9, HR: 5, RBI: 19, SB: 1, Hits: 24.0, AtBats: 89.0},
		{Name: "Chase Burns", Positions: []string{"SP", "RP"}, IsPitcher: true, IP: 34.0, HitsAllowed: 26, ER: 10, BB: 12, SO: 39, W: 3},
		{Name: "Braxton Ashcraft", Positions: []string{"SP", "RP"}, IsPitcher: true, IP: 34, HitsAllowed: 27, ER: 14, BB: 12, SO: 39, W: 1},
		{Name: "Dennis Santana", Positions: []string{"RP"}, IsPitcher: true, IP: 13.2, HitsAllowed: 9, ER: 5, BB: 8, SO: 10, W: 2, SV: 2},
		{Name: "Alex Vesia", Positions: []string{"RP"}, IsPitcher: true, IP: 10.1, HitsAllowed: 4, ER: 2, BB: 5, SO: 14, W: 0, SV: 2},
		{Name: "Garret Crochet", Positions: []string{"SP"}, IsPitcher: true, IP: 30, HitsAllowed: 33, ER: 21, BB: 11, SO: 37, W: 3},
		{Name: "Reid Detmers", Positions: []string{"SP", "RP"}, IsPitcher: true, IP: 33.2, HitsAllowed: 28, ER: 16, BB: 9, SO: 36, W: 1},
		{Name: "Chase Dollander", Positions: []string{"SP", "RP"}, IsPitcher: true, IP: 32, HitsAllowed: 23, ER: 8, BB: 9, SO: 39, W: 3},
		{Name: "Corbin Martin", Positions: []string{"RP"}, IsPitcher: true, IP: 4, HitsAllowed: 1, ER: 0, BB: 1, SO: 2, W: 0, SV: 1},
		{Name: "Freddy Peralta", Positions: []string{"SP"}, IsPitcher: true, IP: 32.1, HitsAllowed: 26, ER: 14, BB: 13, SO: 36, W: 1},
		{Name: "MacKenzie Gore", Positions: []string{"SP"}, IsPitcher: true, IP: 31, HitsAllowed: 25, ER: 15, BB: 15, SO: 42, W: 2},
		{Name: "Phil Maton", Positions: []string{"RP"}, IsPitcher: true, IP: 5, HitsAllowed: 9, ER: 8, BB: 4, SO: 7, W: 0, SV: 1},
	}
}
