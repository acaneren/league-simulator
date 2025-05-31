# Football League Simulator

A simple backend project simulating a 4-team football league, written in Go.  
The league follows Premier League rules and includes API endpoints to play matches, see results, edit scores, and get championship estimates.

## Features

- 4 teams: Chelsea, Arsenal, Manchester City, Liverpool
- 6 weeks, 2 matches per week, each team plays 6 matches
- Premier League points system (3 win, 1 draw, 0 loss, goal difference)
- Play one week (`/play-week`) or the whole season (`/play-all`)
- View standings (`/standings`)
- Edit any match result (`/edit-match`)
- Get simple championship probability after week 4/5 (`/estimate`)
- Data stored in PostgreSQL, deployed on Railway
- No frontend, only REST API

## Live Demo

[https://league-simulator-production-0018.up.railway.app/](https://league-simulator-production-0018.up.railway.app/)

## API Endpoints

| Endpoint       | Method | Description                                        |
|----------------|--------|----------------------------------------------------|
| /play-week     | GET    | Plays next week, returns results as plain text     |
| /play-all      | GET    | Plays all weeks, returns all results as plain text |
| /standings     | GET    | Returns the league table in JSON                   |
| /edit-match    | GET    | Edits a match (id, home, away params)              |
| /estimate      | GET    | Championship probability (after week 4/5)          |
| /reset         | GET    | Resets the league                                  |

### Example: Edit a match

/edit-match?id=1&home=2&away=1


## Database

- PostgreSQL tables and sample queries are included in `schema.sql`
- To manually create tables, use `schema.sql` file

## Run Locally

1. Install Go 1.20+ and PostgreSQL
2. Clone this repo: git clone https://github.com/acaneren/league-simulator.git
3. Set database config in `db.go` (or use Railway's DATABASE_URL)
4. Run the server: go run .
5. Use Postman, curl or browser to test endpoints

## Project Files

- `main.go`       — HTTP server and endpoints
- `db.go`         — Database connection
- `models.go`     — Structs for Team and Match
- `fixtures.go`   — Fixture generator
- `league.go`     — Main league logic (play, edit, update)
- `estimate.go`   — Simple champion estimator
- `service.go`    — Interface layer for league functions
- `schema.sql`    — Table definitions and sample queries

## Project-Links

- Github: [github.com/acaneren/league-simulator](https://github.com/acaneren/league-simulator)
- Live API: [https://league-simulator-production-0018.up.railway.app/](https://league-simulator-production-0018.up.railway.app/)
