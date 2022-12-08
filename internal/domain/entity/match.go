package entity

import (
	"strconv"
	"time"
)

type MatchResult struct {
	TeamAScore int
	TeamBScore int
}

func NewMatchResult(teamAScore, teamBScore int) *MatchResult {
	return &MatchResult{
		TeamAScore: teamAScore,
		TeamBScore: teamBScore,
	}
}

func (m *MatchResult) GetResult() string {
	// 1-0
	return strconv.Itoa(m.TeamAScore) + "-" + strconv.Itoa(m.TeamBScores)
}

type Match struct {
	ID      string
	TeamA   *Team
	TeamB   *Team
	TeamAID string
	TeamBID string
	Date    time.Time
	Status  string
	Result  MatchResult
}

func NewMatch(id string, teamA *Team, teamB *Team, date time.Time) *Match {
	return &Match{
		ID:      id,
		TeamA:   teamA,
		TeamB:   teamB,
		TeamAID: teamA.ID,
		TeamBID: teamB.ID,
		Date:    date,
	}
}
