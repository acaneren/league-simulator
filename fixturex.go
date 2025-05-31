package main

import "fmt"

func InitLeague() {

	DB.Exec("TRUNCATE matches, teams RESTART IDENTITY;")

	teams := []Team{
		{Name: "Chelsea", Strength: 5},
		{Name: "Arsenal", Strength: 4},
		{Name: "Manchester City", Strength: 3},
		{Name: "Liverpool", Strength: 2},
	}
	for _, t := range teams {
		DB.Create(&t)
	}

	var savedTeams []Team
	DB.Find(&savedTeams)

	matchups := [][2]int{
		{0, 1}, {2, 3},
		{0, 2}, {1, 3},
		{0, 3}, {1, 2},
		{1, 0}, {3, 2},
		{2, 0}, {3, 1},
		{3, 0}, {2, 1},
	}
	for i, pair := range matchups {
		match := Match{
			Week:       i/2 + 1,
			HomeTeamID: savedTeams[pair[0]].ID,
			AwayTeamID: savedTeams[pair[1]].ID,
		}
		DB.Create(&match)
	}
	fmt.Println("League initialized: 6 weeks / 12 matches")
}
