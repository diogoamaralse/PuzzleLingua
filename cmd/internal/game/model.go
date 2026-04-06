package game

type Puzzle struct {
	ID         int    `json:"id"`
	Source     string `json:"source"`
	Target     string `json:"target,omitempty"`
	Direction  string `json:"direction"`
	Difficulty string `json:"difficulty"`
	Category   string `json:"category"`
}

type Game struct {
	Puzzles     []Puzzle
	UsedIndexes map[int]bool

	Current    Puzzle
	CurrentIdx int

	Score     int
	Round     int
	MaxRounds int
	HintsUsed int
}

func New(puzzles []Puzzle, maxRounds int) *Game {
	if maxRounds <= 0 {
		maxRounds = 10
	}

	return &Game{
		Puzzles:     puzzles,
		UsedIndexes: map[int]bool{},
		MaxRounds:   maxRounds,
	}
}

type CheckRequest struct {
	ID     int    `json:"id"`
	Answer string `json:"answer"`
}

type CheckResponse struct {
	Correct  bool   `json:"correct"`
	Expected string `json:"expected"`
	Message  string `json:"message"`
}

type RoundResponse struct {
	Puzzle Puzzle `json:"puzzle"`
}
