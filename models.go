package main

type Team struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"unique"`
	Strength     int
	Played       int
	Won          int
	Drawn        int
	Lost         int
	GoalsFor     int
	GoalsAgainst int
	Points       int
}

type Match struct {
	ID         uint `gorm:"primaryKey"`
	Week       int
	HomeTeamID uint
	AwayTeamID uint
	HomeGoals  int
	AwayGoals  int
	IsPlayed   bool
}
