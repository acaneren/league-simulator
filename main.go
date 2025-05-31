package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var leagueService LeagueService = &MyLeagueService{}

func main() {
	ConnectDB()
	InitLeague()

	http.HandleFunc("/play-week", handlePlayWeek)
	http.HandleFunc("/play-all", handlePlayAll)
	http.HandleFunc("/standings", handleStandings)
	http.HandleFunc("/reset", handleReset)
	http.HandleFunc("/estimate", handleEstimate)
	http.HandleFunc("/edit-match", handleEditMatch)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handlePlayWeek(w http.ResponseWriter, r *http.Request) {
	var m []Match
	DB.Where("is_played = false").Order("week asc").Limit(2).Find(&m)
	if len(m) == 0 {
		w.Write([]byte("No more matches.\n"))
		return
	}
	week := m[0].Week
	PlayNextWeek()
	var results []Match
	DB.Where("week = ? AND is_played = true", week).Find(&results)
	out := ""
	for _, x := range results {
		var h, a Team
		DB.First(&h, x.HomeTeamID)
		DB.First(&a, x.AwayTeamID)
		out += fmt.Sprintf("Week %d: %s %d-%d %s\n", x.Week, h.Name, x.HomeGoals, x.AwayGoals, a.Name)
	}
	w.Write([]byte(out))
}

func handlePlayAll(w http.ResponseWriter, r *http.Request) {
	PlayAll()
	var matches []Match
	DB.Where("is_played = true").Order("week asc, id asc").Find(&matches)

	result := ""
	for _, m := range matches {
		var h, a Team
		DB.First(&h, m.HomeTeamID)
		DB.First(&a, m.AwayTeamID)
		result += fmt.Sprintf("Week %d: %s %d-%d %s\n", m.Week, h.Name, m.HomeGoals, m.AwayGoals, a.Name)
	}
	w.Write([]byte(result))
}

func handleStandings(w http.ResponseWriter, r *http.Request) {
	var teams []Team
	DB.Order("points desc, goals_for - goals_against desc").Find(&teams)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

func handleReset(w http.ResponseWriter, r *http.Request) {
	leagueService.ResetLeague()
	w.Write([]byte("The league has been reset\n"))
}

func handleEstimate(w http.ResponseWriter, r *http.Request) {
	result := leagueService.EstimateChampions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func handleEditMatch(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	home := toInt(r.URL.Query().Get("home"))
	away := toInt(r.URL.Query().Get("away"))

	if id == "" {
		http.Error(w, "id required", http.StatusBadRequest)
		return
	}

	message := EditMatch(id, home, away)
	w.Write([]byte(message + "\n"))
}
