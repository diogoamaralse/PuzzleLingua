package ui

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"PuzzleLingua/cmd/internal/data"
	"PuzzleLingua/cmd/internal/game"
)

type AppUI struct {
	game *game.Game

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

func Run() {
	rand.Seed(time.Now().UnixNano())

	a := app.New()
	w := a.NewWindow("PuzzleLingua")
	w.Resize(fyne.NewSize(860, 580))

	ui := New()
	w.SetContent(ui.View())

	ui.applyState(ui.game.LoadNextPuzzle())

	w.ShowAndRun()
}

func New() *AppUI {
	g := game.New(data.DefaultPuzzles(), 5)

	ui := &AppUI{
		game:            g,
		titleLabel:      widget.NewLabel("PuzzleLingua"),
		roundLabel:      widget.NewLabel("Round: 0/0"),
		scoreLabel:      widget.NewLabel("Score: 0"),
		directionLabel:  widget.NewLabel("Direction: -"),
		difficultyLabel: widget.NewLabel("Difficulty: -"),
		categoryLabel:   widget.NewLabel("Category: -"),
		wordLabel:       widget.NewLabel("Translate: -"),
		feedbackLabel:   widget.NewLabel("Welcome! Translate the word shown on screen."),
		input:           widget.NewEntry(),
	}

	ui.titleLabel.TextStyle = fyne.TextStyle{Bold: true}
	ui.input.SetPlaceHolder("Type the translation here")

	ui.checkButton = widget.NewButton("Check", func() {
		ui.applyState(ui.game.CheckAnswer(ui.input.Text))
	})

	ui.hintButton = widget.NewButton("Hint", func() {
		ui.applyState(ui.game.ShowHint())
	})

	ui.skipButton = widget.NewButton("Skip", func() {
		ui.applyState(ui.game.SkipPuzzle())
	})

	ui.nextButton = widget.NewButton("Next", func() {
		ui.applyState(ui.game.LoadNextPuzzle())
	})
	ui.nextButton.Disable()

	ui.restartBtn = widget.NewButton("Restart", func() {
		ui.input.SetText("")
		ui.applyState(ui.game.Restart())
	})

	ui.input.OnSubmitted = func(_ string) {
		ui.applyState(ui.game.CheckAnswer(ui.input.Text))
	}

	return ui
}

func (ui *AppUI) View() fyne.CanvasObject {
	header := container.NewVBox(
		ui.titleLabel,
		container.NewGridWithColumns(2, ui.roundLabel, ui.scoreLabel),
	)

	meta := container.NewGridWithColumns(3,
		ui.directionLabel,
		ui.difficultyLabel,
		ui.categoryLabel,
	)

	contentBox := container.NewVBox(
		widget.NewSeparator(),
		meta,
		widget.NewSeparator(),
		ui.wordLabel,
		ui.input,
		ui.feedbackLabel,
	)

	actions := container.NewGridWithColumns(5,
		ui.checkButton,
		ui.hintButton,
		ui.skipButton,
		ui.nextButton,
		ui.restartBtn,
	)

	return container.NewPadded(
		container.NewVBox(
			header,
			contentBox,
			actions,
		),
	)
}

func (ui *AppUI) applyState(state game.RoundState) {
	ui.roundLabel.SetText(fmt.Sprintf("Round: %d/%d", state.Round, state.MaxRounds))
	ui.scoreLabel.SetText(fmt.Sprintf("Score: %d", state.Score))
	ui.feedbackLabel.SetText(state.Message)

	if state.Finished {
		ui.directionLabel.SetText("Direction: -")
		ui.difficultyLabel.SetText("Difficulty: -")
		ui.categoryLabel.SetText("Category: -")
		ui.wordLabel.SetText("Game finished")

		ui.input.Disable()
		ui.checkButton.Disable()
		ui.hintButton.Disable()
		ui.skipButton.Disable()
		ui.nextButton.Disable()
		return
	}

	ui.directionLabel.SetText(fmt.Sprintf("Direction: %s", state.Current.Direction))
	ui.difficultyLabel.SetText(fmt.Sprintf("Difficulty: %s", state.Current.Difficulty))
	ui.categoryLabel.SetText(fmt.Sprintf("Category: %s", state.Current.Category))
	ui.wordLabel.SetText(fmt.Sprintf("Translate: %s", state.Current.Source))

	if state.CanAdvance {
		ui.input.Disable()
		ui.checkButton.Disable()
		ui.hintButton.Disable()
		ui.skipButton.Disable()
		ui.nextButton.Enable()
		return
	}

	ui.input.Enable()
	ui.checkButton.Enable()
	ui.hintButton.Enable()
	ui.skipButton.Enable()
	ui.nextButton.Disable()
}
