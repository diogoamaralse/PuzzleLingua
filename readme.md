# PuzzleLingua

PuzzleLingua is a Fyne desktop game for practicing Portuguese ↔ English translations.

## Run

```bash
go mod tidy
go run ./cmd/puzzlelingua
```

## Structure

- `cmd/puzzlelingua`: application entrypoint
- `internal/game`: game rules and state
- `internal/data`: built-in vocabulary
- `internal/ui`: Fyne desktop interface