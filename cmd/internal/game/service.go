package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Service struct {
	puzzles []Puzzle
	byID    map[int]Puzzle
	rng     *rand.Rand
}

func NewService(puzzles []Puzzle) *Service {
	byID := make(map[int]Puzzle, len(puzzles))
	for _, p := range puzzles {
		byID[p.ID] = p
	}

	return &Service{
		puzzles: puzzles,
		byID:    byID,
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Service) RandomPuzzle() Puzzle {
	p := s.puzzles[s.rng.Intn(len(s.puzzles))]
	p.Target = ""
	return p
}

func (s *Service) Check(req CheckRequest) CheckResponse {
	p, ok := s.byID[req.ID]
	if !ok {
		return CheckResponse{
			Correct:  false,
			Expected: "",
			Message:  "Puzzle not found.",
		}
	}

	answer := Normalize(req.Answer)
	expected := Normalize(p.Target)

	if answer == expected {
		return CheckResponse{
			Correct:  true,
			Expected: p.Target,
			Message:  fmt.Sprintf("Correct! %s = %s", p.Source, p.Target),
		}
	}

	return CheckResponse{
		Correct:  false,
		Expected: p.Target,
		Message:  fmt.Sprintf("Incorrect. Correct answer: %s", p.Target),
	}
}
