-- Tables for the football league project

CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE,
    strength INTEGER,
    played INTEGER,
    won INTEGER,
    drawn INTEGER,
    lost INTEGER,
    goals_for INTEGER,
    goals_against INTEGER,
    points INTEGER
);

CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    week INTEGER,
    home_team_id INTEGER REFERENCES teams(id),
    away_team_id INTEGER REFERENCES teams(id),
    home_goals INTEGER,
    away_goals INTEGER,
    is_played BOOLEAN
);

-- League standings
SELECT * FROM teams ORDER BY points DESC, goals_for - goals_against DESC;

-- All matches with their scores
SELECT * FROM matches ORDER BY week, id;

-- Lists matches for a specific team
SELECT * FROM matches WHERE home_team_id = 1 OR away_team_id = 1 ORDER BY week;

-- Only playd matches
SELECT * FROM matches WHERE is_played = true ORDER BY week;

-- Matches with team names using JOIN
SELECT m.week, th.name AS home_team, ta.name AS away_team,
       m.home_goals, m.away_goals, m.is_played
FROM matches m
JOIN teams th ON m.home_team_id = th.id
JOIN teams ta ON m.away_team_id = ta.id
ORDER BY m.week, m.id;
