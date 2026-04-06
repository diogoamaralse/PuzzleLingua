package game

import (
	"fmt"
	"math/rand"
	"strings"
)

type RoundState struct {
	Finished   bool
	Message    string
	Current    Puzzle
	Round      int
	MaxRounds  int
	Score      int
	CanAdvance bool
}

func (g *Game) LoadNextPuzzle() RoundState {
	if g.Round >= g.MaxRounds {
		return g.EndState()
	}

	idx := g.randomUnusedPuzzleIndex()
	if idx == -1 {
		return g.EndState()
	}

	g.UsedIndexes[idx] = true
	g.CurrentIdx = idx
	g.Current = g.Puzzles[idx]
	g.HintsUsed = 0
	g.Round++

	return RoundState{
		Finished:   false,
		Message:    "Type the correct translation and press Check.",
		Current:    g.Current,
		Round:      g.Round,
		MaxRounds:  g.MaxRounds,
		Score:      g.Score,
		CanAdvance: false,
	}
}

func (g *Game) CheckAnswer(answer string) RoundState {
	answer = Normalize(answer)
	target := Normalize(g.Current.Target)

	if answer == "" {
		return g.currentState("Please type a translation first.", false)
	}

	if answer == target {
		points := scoreForDifficulty(g.Current.Difficulty) - (g.HintsUsed * 2)
		if points < 3 {
			points = 3
		}

		g.Score += points
		return g.currentState(
			fmt.Sprintf("Correct! \"%s\" = \"%s\". +%d points", g.Current.Source, g.Current.Target, points),
			true,
		)
	}

	return g.currentState("Incorrect. Try again.", false)
}

func (g *Game) ShowHint() RoundState {
	g.HintsUsed++

	targetRunes := []rune(g.Current.Target)
	if len(targetRunes) == 0 {
		return g.currentState("No hint available.", false)
	}

	revealCount := g.HintsUsed
	if revealCount > len(targetRunes) {
		revealCount = len(targetRunes)
	}

	revealed := string(targetRunes[:revealCount])
	masked := strings.Repeat("_", len(targetRunes)-revealCount)

	return g.currentState(
		fmt.Sprintf("Hint: %s%s (%d letters)", strings.ToUpper(revealed), masked, len(targetRunes)),
		false,
	)
}

func (g *Game) SkipPuzzle() RoundState {
	return g.currentState(fmt.Sprintf("Skipped. Correct answer: %s", g.Current.Target), true)
}

func (g *Game) Restart() RoundState {
	g.Score = 0
	g.Round = 0
	g.HintsUsed = 0
	g.UsedIndexes = map[int]bool{}
	return g.LoadNextPuzzle()
}

func (g *Game) EndState() RoundState {
	return RoundState{
		Finished:   true,
		Message:    fmt.Sprintf("Final score: %d. Press Restart to play again.", g.Score),
		Current:    Puzzle{},
		Round:      g.Round,
		MaxRounds:  g.MaxRounds,
		Score:      g.Score,
		CanAdvance: false,
	}
}

func (g *Game) currentState(message string, canAdvance bool) RoundState {
	return RoundState{
		Finished:   false,
		Message:    message,
		Current:    g.Current,
		Round:      g.Round,
		MaxRounds:  g.MaxRounds,
		Score:      g.Score,
		CanAdvance: canAdvance,
	}
}

func (g *Game) randomUnusedPuzzleIndex() int {
	if len(g.UsedIndexes) >= len(g.Puzzles) {
		return -1
	}

	for {
		idx := rand.Intn(len(g.Puzzles))
		if !g.UsedIndexes[idx] {
			return idx
		}
	}
}

func scoreForDifficulty(level string) int {
	switch level {
	case "Hard":
		return 18
	case "Medium":
		return 12
	default:
		return 8
	}
}
