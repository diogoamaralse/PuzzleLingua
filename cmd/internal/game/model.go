package game

type Puzzle struct {
	Source     string
	Target     string
	Direction  string
	Difficulty string
	Category   string
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
