package main

type LeagueService interface {
	PlayNextWeek()
	PlayAll()
	ResetLeague()
	EstimateChampions() map[string]float64
}

type MyLeagueService struct{}

func (s *MyLeagueService) PlayNextWeek() {
	PlayNextWeek()
}

func (s *MyLeagueService) PlayAll() {
	PlayAll()
}

func (s *MyLeagueService) ResetLeague() {
	InitLeague()
}

func (s *MyLeagueService) EstimateChampions() map[string]float64 {
	return EstimateChampions()
}
