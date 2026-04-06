package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Puzzle struct {
	Source     string
	Target     string
	Direction  string // PT->EN or EN->PT
	Difficulty string
	Category   string
}

type Game struct {
	puzzles     []Puzzle
	usedIndexes map[int]bool

	current    Puzzle
	currentIdx int

	score     int
	round     int
	maxRounds int
	hintsUsed int

	titleLabel      *widget.Label
	roundLabel      *widget.Label
	scoreLabel      *widget.Label
	directionLabel  *widget.Label
	difficultyLabel *widget.Label
	categoryLabel   *widget.Label
	wordLabel       *widget.Label
	feedbackLabel   *widget.Label

	input *widget.Entry

	checkButton *widget.Button
	hintButton  *widget.Button
	skipButton  *widget.Button
	nextButton  *widget.Button
	restartBtn  *widget.Button
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a := app.New()
	w := a.NewWindow("PuzzleLingua")
	w.Resize(fyne.NewSize(820, 560))

	game := NewGame()
	w.SetContent(game.UI())
	game.LoadNextPuzzle()

	w.ShowAndRun()
}

func NewGame() *Game {
	g := &Game{
		maxRounds:   10,
		usedIndexes: map[int]bool{},
		puzzles: []Puzzle{
			{Source: "casa", Target: "house", Direction: "PT->EN", Difficulty: "Easy", Category: "Home"},
			{Source: "livro", Target: "book", Direction: "PT->EN", Difficulty: "Easy", Category: "Objects"},
			{Source: "gato", Target: "cat", Direction: "PT->EN", Difficulty: "Easy", Category: "Animals"},
			{Source: "janela", Target: "window", Direction: "PT->EN", Difficulty: "Easy", Category: "Home"},
			{Source: "escola", Target: "school", Direction: "PT->EN", Difficulty: "Easy", Category: "Places"},
			{Source: "cadeira", Target: "chair", Direction: "PT->EN", Difficulty: "Easy", Category: "Home"},
			{Source: "estrada", Target: "road", Direction: "PT->EN", Difficulty: "Easy", Category: "Places"},
			{Source: "praia", Target: "beach", Direction: "PT->EN", Difficulty: "Easy", Category: "Places"},
			{Source: "cidade", Target: "city", Direction: "PT->EN", Difficulty: "Easy", Category: "Places"},
			{Source: "comida", Target: "food", Direction: "PT->EN", Difficulty: "Easy", Category: "Daily Life"},
			{Source: "tempo", Target: "weather", Direction: "PT->EN", Difficulty: "Easy", Category: "Daily Life"},
			{Source: "porta", Target: "door", Direction: "PT->EN", Difficulty: "Easy", Category: "Home"},

			{Source: "amizade", Target: "friendship", Direction: "PT->EN", Difficulty: "Medium", Category: "Abstract"},
			{Source: "trabalho", Target: "work", Direction: "PT->EN", Difficulty: "Medium", Category: "Daily Life"},
			{Source: "mercado", Target: "market", Direction: "PT->EN", Difficulty: "Medium", Category: "Places"},
			{Source: "caminho", Target: "path", Direction: "PT->EN", Difficulty: "Medium", Category: "Places"},
			{Source: "tempestade", Target: "storm", Direction: "PT->EN", Difficulty: "Medium", Category: "Nature"},
			{Source: "biblioteca", Target: "library", Direction: "PT->EN", Difficulty: "Medium", Category: "Places"},
			{Source: "descoberta", Target: "discovery", Direction: "PT->EN", Difficulty: "Medium", Category: "Abstract"},
			{Source: "montanha", Target: "mountain", Direction: "PT->EN", Difficulty: "Medium", Category: "Nature"},
			{Source: "desafio", Target: "challenge", Direction: "PT->EN", Difficulty: "Medium", Category: "Abstract"},
			{Source: "lembranca", Target: "memory", Direction: "PT->EN", Difficulty: "Medium", Category: "Abstract"},
			{Source: "viagem", Target: "journey", Direction: "PT->EN", Difficulty: "Medium", Category: "Travel"},
			{Source: "cozinha", Target: "kitchen", Direction: "PT->EN", Difficulty: "Medium", Category: "Home"},

			{Source: "conhecimento", Target: "knowledge", Direction: "PT->EN", Difficulty: "Hard", Category: "Abstract"},
			{Source: "ferramenta", Target: "tool", Direction: "PT->EN", Difficulty: "Hard", Category: "Objects"},
			{Source: "desenvolvimento", Target: "development", Direction: "PT->EN", Difficulty: "Hard", Category: "Work"},
			{Source: "aprendizagem", Target: "learning", Direction: "PT->EN", Difficulty: "Hard", Category: "Education"},
			{Source: "possibilidade", Target: "possibility", Direction: "PT->EN", Difficulty: "Hard", Category: "Abstract"},
			{Source: "relacionamento", Target: "relationship", Direction: "PT->EN", Difficulty: "Hard", Category: "Abstract"},
			{Source: "responsabilidade", Target: "responsibility", Direction: "PT->EN", Difficulty: "Hard", Category: "Abstract"},
			{Source: "comportamento", Target: "behavior", Direction: "PT->EN", Difficulty: "Hard", Category: "Abstract"},
			{Source: "conversacao", Target: "conversation", Direction: "PT->EN", Difficulty: "Hard", Category: "Communication"},
			{Source: "oportunidade", Target: "opportunity", Direction: "PT->EN", Difficulty: "Hard", Category: "Abstract"},
			{Source: "investigacao", Target: "research", Direction: "PT->EN", Difficulty: "Hard", Category: "Education"},
			{Source: "construcao", Target: "construction", Direction: "PT->EN", Difficulty: "Hard", Category: "Work"},

			{Source: "house", Target: "casa", Direction: "EN->PT", Difficulty: "Easy", Category: "Home"},
			{Source: "book", Target: "livro", Direction: "EN->PT", Difficulty: "Easy", Category: "Objects"},
			{Source: "cat", Target: "gato", Direction: "EN->PT", Difficulty: "Easy", Category: "Animals"},
			{Source: "window", Target: "janela", Direction: "EN->PT", Difficulty: "Easy", Category: "Home"},
			{Source: "school", Target: "escola", Direction: "EN->PT", Difficulty: "Easy", Category: "Places"},
			{Source: "chair", Target: "cadeira", Direction: "EN->PT", Difficulty: "Easy", Category: "Home"},
			{Source: "road", Target: "estrada", Direction: "EN->PT", Difficulty: "Easy", Category: "Places"},
			{Source: "beach", Target: "praia", Direction: "EN->PT", Difficulty: "Easy", Category: "Places"},
			{Source: "city", Target: "cidade", Direction: "EN->PT", Difficulty: "Easy", Category: "Places"},
			{Source: "food", Target: "comida", Direction: "EN->PT", Difficulty: "Easy", Category: "Daily Life"},
			{Source: "weather", Target: "tempo", Direction: "EN->PT", Difficulty: "Easy", Category: "Daily Life"},
			{Source: "door", Target: "porta", Direction: "EN->PT", Difficulty: "Easy", Category: "Home"},

			{Source: "friendship", Target: "amizade", Direction: "EN->PT", Difficulty: "Medium", Category: "Abstract"},
			{Source: "work", Target: "trabalho", Direction: "EN->PT", Difficulty: "Medium", Category: "Daily Life"},
			{Source: "market", Target: "mercado", Direction: "EN->PT", Difficulty: "Medium", Category: "Places"},
			{Source: "path", Target: "caminho", Direction: "EN->PT", Difficulty: "Medium", Category: "Places"},
			{Source: "storm", Target: "tempestade", Direction: "EN->PT", Difficulty: "Medium", Category: "Nature"},
			{Source: "library", Target: "biblioteca", Direction: "EN->PT", Difficulty: "Medium", Category: "Places"},
			{Source: "discovery", Target: "descoberta", Direction: "EN->PT", Difficulty: "Medium", Category: "Abstract"},
			{Source: "mountain", Target: "montanha", Direction: "EN->PT", Difficulty: "Medium", Category: "Nature"},
			{Source: "challenge", Target: "desafio", Direction: "EN->PT", Difficulty: "Medium", Category: "Abstract"},
			{Source: "memory", Target: "lembranca", Direction: "EN->PT", Difficulty: "Medium", Category: "Abstract"},
			{Source: "journey", Target: "viagem", Direction: "EN->PT", Difficulty: "Medium", Category: "Travel"},
			{Source: "kitchen", Target: "cozinha", Direction: "EN->PT", Difficulty: "Medium", Category: "Home"},

			{Source: "knowledge", Target: "conhecimento", Direction: "EN->PT", Difficulty: "Hard", Category: "Abstract"},
			{Source: "tool", Target: "ferramenta", Direction: "EN->PT", Difficulty: "Hard", Category: "Objects"},
			{Source: "development", Target: "desenvolvimento", Direction: "EN->PT", Difficulty: "Hard", Category: "Work"},
			{Source: "learning", Target: "aprendizagem", Direction: "EN->PT", Difficulty: "Hard", Category: "Education"},
			{Source: "possibility", Target: "possibilidade", Direction: "EN->PT", Difficulty: "Hard", Category: "Abstract"},
			{Source: "relationship", Target: "relacionamento", Direction: "EN->PT", Difficulty: "Hard", Category: "Abstract"},
			{Source: "responsibility", Target: "responsabilidade", Direction: "EN->PT", Difficulty: "Hard", Category: "Abstract"},
			{Source: "behavior", Target: "comportamento", Direction: "EN->PT", Difficulty: "Hard", Category: "Abstract"},
			{Source: "conversation", Target: "conversacao", Direction: "EN->PT", Difficulty: "Hard", Category: "Communication"},
			{Source: "opportunity", Target: "oportunidade", Direction: "EN->PT", Difficulty: "Hard", Category: "Abstract"},
			{Source: "research", Target: "investigacao", Direction: "EN->PT", Difficulty: "Hard", Category: "Education"},
			{Source: "construction", Target: "construcao", Direction: "EN->PT", Difficulty: "Hard", Category: "Work"},
		},
	}

	g.titleLabel = widget.NewLabel("PuzzleLingua")
	g.titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	g.roundLabel = widget.NewLabel("Round: 0/0")
	g.scoreLabel = widget.NewLabel("Score: 0")
	g.directionLabel = widget.NewLabel("Direction: -")
	g.difficultyLabel = widget.NewLabel("Difficulty: -")
	g.categoryLabel = widget.NewLabel("Category: -")
	g.wordLabel = widget.NewLabel("Translate: -")
	g.feedbackLabel = widget.NewLabel("Welcome! Translate the word shown on screen.")

	g.input = widget.NewEntry()
	g.input.SetPlaceHolder("Type the translation here")

	g.checkButton = widget.NewButton("Check", func() {
		g.CheckAnswer()
	})

	g.hintButton = widget.NewButton("Hint", func() {
		g.ShowHint()
	})

	g.skipButton = widget.NewButton("Skip", func() {
		g.SkipPuzzle()
	})

	g.nextButton = widget.NewButton("Next", func() {
		g.LoadNextPuzzle()
	})
	g.nextButton.Disable()

	g.restartBtn = widget.NewButton("Restart", func() {
		g.RestartGame()
	})

	g.input.OnSubmitted = func(_ string) {
		g.CheckAnswer()
	}

	return g
}

func (g *Game) UI() fyne.CanvasObject {
	header := container.NewVBox(
		g.titleLabel,
		container.NewGridWithColumns(2, g.roundLabel, g.scoreLabel),
	)

	meta := container.NewGridWithColumns(3,
		g.directionLabel,
		g.difficultyLabel,
		g.categoryLabel,
	)

	contentBox := container.NewVBox(
		widget.NewSeparator(),
		meta,
		widget.NewSeparator(),
		g.wordLabel,
		g.input,
		g.feedbackLabel,
	)

	actions := container.NewGridWithColumns(5,
		g.checkButton,
		g.hintButton,
		g.skipButton,
		g.nextButton,
		g.restartBtn,
	)

	root := container.NewVBox(
		header,
		contentBox,
		actions,
	)

	return container.NewPadded(root)
}

func (g *Game) LoadNextPuzzle() {
	if g.round >= g.maxRounds {
		g.EndGame()
		return
	}

	idx := g.randomUnusedPuzzleIndex()
	if idx == -1 {
		g.EndGame()
		return
	}

	g.usedIndexes[idx] = true
	g.currentIdx = idx
	g.current = g.puzzles[idx]
	g.hintsUsed = 0
	g.round++

	g.roundLabel.SetText(fmt.Sprintf("Round: %d/%d", g.round, g.maxRounds))
	g.scoreLabel.SetText(fmt.Sprintf("Score: %d", g.score))
	g.directionLabel.SetText(fmt.Sprintf("Direction: %s", g.current.Direction))
	g.difficultyLabel.SetText(fmt.Sprintf("Difficulty: %s", g.current.Difficulty))
	g.categoryLabel.SetText(fmt.Sprintf("Category: %s", g.current.Category))
	g.wordLabel.SetText(fmt.Sprintf("Translate: %s", strings.ToUpper(g.current.Source)))
	g.feedbackLabel.SetText("Type the correct translation and press Check.")
	g.input.SetText("")
	g.input.Enable()

	g.checkButton.Enable()
	g.hintButton.Enable()
	g.skipButton.Enable()
	g.nextButton.Disable()
}

func (g *Game) CheckAnswer() {
	answer := normalize(g.input.Text)
	target := normalize(g.current.Target)

	if answer == "" {
		g.feedbackLabel.SetText("Please type a translation first.")
		return
	}

	if answer == target {
		points := scoreForDifficulty(g.current.Difficulty) - (g.hintsUsed * 2)
		if points < 3 {
			points = 3
		}

		g.score += points
		g.scoreLabel.SetText(fmt.Sprintf("Score: %d", g.score))
		g.feedbackLabel.SetText(fmt.Sprintf("Correct! \"%s\" = \"%s\". +%d points", g.current.Source, g.current.Target, points))
		g.endRound()
		return
	}

	g.feedbackLabel.SetText("Incorrect. Try again.")
}

func (g *Game) ShowHint() {
	g.hintsUsed++

	targetRunes := []rune(g.current.Target)
	if len(targetRunes) == 0 {
		return
	}

	revealCount := g.hintsUsed
	if revealCount > len(targetRunes) {
		revealCount = len(targetRunes)
	}

	revealed := string(targetRunes[:revealCount])
	masked := strings.Repeat("_", len(targetRunes)-revealCount)

	g.feedbackLabel.SetText(fmt.Sprintf("Hint: %s%s (%d letters)", strings.ToUpper(revealed), masked, len(targetRunes)))
}

func (g *Game) SkipPuzzle() {
	g.feedbackLabel.SetText(fmt.Sprintf("Skipped. Correct answer: %s", g.current.Target))
	g.endRound()
}

func (g *Game) endRound() {
	g.input.Disable()
	g.checkButton.Disable()
	g.hintButton.Disable()
	g.skipButton.Disable()
	g.nextButton.Enable()
}

func (g *Game) EndGame() {
	g.wordLabel.SetText("Game finished")
	g.directionLabel.SetText("Direction: -")
	g.difficultyLabel.SetText("Difficulty: -")
	g.categoryLabel.SetText("Category: -")
	g.feedbackLabel.SetText(fmt.Sprintf("Final score: %d. Press Restart to play again.", g.score))
	g.input.Disable()
	g.checkButton.Disable()
	g.hintButton.Disable()
	g.skipButton.Disable()
	g.nextButton.Disable()
}

func (g *Game) RestartGame() {
	g.score = 0
	g.round = 0
	g.hintsUsed = 0
	g.usedIndexes = map[int]bool{}
	g.scoreLabel.SetText("Score: 0")
	g.LoadNextPuzzle()
}

func (g *Game) randomUnusedPuzzleIndex() int {
	if len(g.usedIndexes) >= len(g.puzzles) {
		return -1
	}

	for {
		idx := rand.Intn(len(g.puzzles))
		if !g.usedIndexes[idx] {
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

func normalize(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))

	replacer := strings.NewReplacer(
		"á", "a",
		"à", "a",
		"ã", "a",
		"â", "a",
		"ä", "a",
		"é", "e",
		"ê", "e",
		"è", "e",
		"ë", "e",
		"í", "i",
		"ì", "i",
		"ï", "i",
		"ó", "o",
		"ô", "o",
		"õ", "o",
		"ò", "o",
		"ö", "o",
		"ú", "u",
		"ù", "u",
		"ü", "u",
		"ç", "c",
	)
	return replacer.Replace(s)
}
