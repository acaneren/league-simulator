package main

import "math/rand"

func EstimateChampions() map[string]float64 {
	var matches []Match
	DB.Where("is_played = false").Order("week asc").Find(&matches)
	if len(matches) > 4 {
		return nil
	}
	var teams []Team
	DB.Find(&teams)
	winCount := make(map[string]int)
	simCount := 50

	for i := 0; i < simCount; i++ {
		tempTeams := cloneTeams(teams)
		teamMap := make(map[uint]*Team)
		for j := range tempTeams {
			teamMap[tempTeams[j].ID] = &tempTeams[j]
		}
		for j := 0; j < 4 && j < len(matches); j++ {
			m := matches[j]
			home := teamMap[m.HomeTeamID]
			away := teamMap[m.AwayTeamID]
			hg := rand.Intn(4)
			ag := rand.Intn(4)
			updateStats(home, away, hg, ag)
		}
		winner := getTopTeam(teamMap)
		winCount[winner.Name]++
	}

	result := make(map[string]float64)
	for _, t := range teams {
		result[t.Name] = float64(winCount[t.Name]) * 100.0 / float64(simCount)
	}
	return result
}

func cloneTeams(original_teams []Team) []Team {
	cloned_teams := make([]Team, len(original_teams))
	copy(cloned_teams, original_teams)
	return cloned_teams
}

func getTopTeam(m map[uint]*Team) *Team {
	var top *Team
	for _, t := range m {
		if top == nil || t.Points > top.Points ||
			(t.Points == top.Points && (t.GoalsFor-t.GoalsAgainst) > (top.GoalsFor-top.GoalsAgainst)) {
			top = t
		}
	}
	return top
}
