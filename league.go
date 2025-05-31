package main

import (
	"math/rand"
	"strconv"
)

func PlayNextWeek() {
	var matches []Match
	DB.Where("is_played = false").Order("week asc").Limit(2).Find(&matches)

	for _, m := range matches {
		var home, away Team
		DB.First(&home, m.HomeTeamID)
		DB.First(&away, m.AwayTeamID)
		hg := rand.Intn(home.Strength + 1)
		ag := rand.Intn(away.Strength + 1)

		m.HomeGoals = hg
		m.AwayGoals = ag
		m.IsPlayed = true

		updateStats(&home, &away, hg, ag)
		DB.Save(&m)
		DB.Save(&home)
		DB.Exec("UPDATE teams SET played=?, won=?, drawn=?, lost=?, goals_for=?, goals_against=?, points=? WHERE id=?",
			away.Played, away.Won, away.Drawn, away.Lost, away.GoalsFor, away.GoalsAgainst, away.Points, away.ID)
	}
}

func PlayAll() {
	for {
		var left int64
		DB.Model(&Match{}).Where("is_played = false").Count(&left)
		if left == 0 {
			break
		}
		PlayNextWeek()
	}
}

func EditMatch(id string, newHome, newAway int) string {
	var match Match
	DB.First(&match, id)
	if match.ID == 0 {
		return "Match not found"
	}

	var home, away Team
	DB.First(&home, match.HomeTeamID)
	DB.First(&away, match.AwayTeamID)

	undoStats(&home, &away, match.HomeGoals, match.AwayGoals)

	match.HomeGoals = newHome
	match.AwayGoals = newAway
	match.IsPlayed = true

	updateStats(&home, &away, newHome, newAway)

	DB.Save(&match)
	DB.Save(&home)
	DB.Exec("UPDATE teams SET played=?, won=?, drawn=?, lost=?, goals_for=?, goals_against=?, points=? WHERE id=?",
		away.Played, away.Won, away.Drawn, away.Lost, away.GoalsFor, away.GoalsAgainst, away.Points, away.ID)

	return "Match updated"
}

func updateStats(home, away *Team, hg, ag int) {
	home.Played++
	away.Played++
	home.GoalsFor += hg
	home.GoalsAgainst += ag
	away.GoalsFor += ag
	away.GoalsAgainst += hg

	if hg > ag {
		home.Won++
		home.Points += 3
		away.Lost++
	} else if ag > hg {
		away.Won++
		away.Points += 3
		home.Lost++
	} else {
		home.Drawn++
		away.Drawn++
		home.Points++
		away.Points++
	}
}

func undoStats(home, away *Team, hg, ag int) {
	home.Played--
	away.Played--
	home.GoalsFor -= hg
	home.GoalsAgainst -= ag
	away.GoalsFor -= ag
	away.GoalsAgainst -= hg

	if hg > ag {
		home.Won--
		home.Points -= 3
		away.Lost--
	} else if ag > hg {
		away.Won--
		away.Points -= 3
		home.Lost--
	} else {
		home.Drawn--
		away.Drawn--
		home.Points--
		away.Points--
	}
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
