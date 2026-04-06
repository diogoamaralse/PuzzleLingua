package data

import "PuzzleLingua/cmd/internal/game"

func DefaultPuzzles() []game.Puzzle {
	return []game.Puzzle{
		{ID: 1, Source: "casa", Target: "house", Direction: "PT->EN", Difficulty: "Easy", Category: "Home"},
		{ID: 2, Source: "livro", Target: "book", Direction: "PT->EN", Difficulty: "Easy", Category: "Objects"},
		{ID: 3, Source: "gato", Target: "cat", Direction: "PT->EN", Difficulty: "Easy", Category: "Animals"},
		{ID: 4, Source: "janela", Target: "window", Direction: "PT->EN", Difficulty: "Easy", Category: "Home"},
		{ID: 5, Source: "escola", Target: "school", Direction: "PT->EN", Difficulty: "Easy", Category: "Places"},
		{ID: 6, Source: "amizade", Target: "friendship", Direction: "PT->EN", Difficulty: "Medium", Category: "Abstract"},
		{ID: 7, Source: "biblioteca", Target: "library", Direction: "PT->EN", Difficulty: "Medium", Category: "Places"},
		{ID: 8, Source: "descoberta", Target: "discovery", Direction: "PT->EN", Difficulty: "Medium", Category: "Abstract"},
		{ID: 9, Source: "conhecimento", Target: "knowledge", Direction: "PT->EN", Difficulty: "Hard", Category: "Abstract"},
		{ID: 10, Source: "desenvolvimento", Target: "development", Direction: "PT->EN", Difficulty: "Hard", Category: "Work"},
		{ID: 11, Source: "house", Target: "casa", Direction: "EN->PT", Difficulty: "Easy", Category: "Home"},
		{ID: 12, Source: "book", Target: "livro", Direction: "EN->PT", Difficulty: "Easy", Category: "Objects"},
		{ID: 13, Source: "cat", Target: "gato", Direction: "EN->PT", Difficulty: "Easy", Category: "Animals"},
		{ID: 14, Source: "school", Target: "escola", Direction: "EN->PT", Difficulty: "Easy", Category: "Places"},
		{ID: 15, Source: "friendship", Target: "amizade", Direction: "EN->PT", Difficulty: "Medium", Category: "Abstract"},
		{ID: 16, Source: "library", Target: "biblioteca", Direction: "EN->PT", Difficulty: "Medium", Category: "Places"},
		{ID: 17, Source: "knowledge", Target: "conhecimento", Direction: "EN->PT", Difficulty: "Hard", Category: "Abstract"},
		{ID: 18, Source: "development", Target: "desenvolvimento", Direction: "EN->PT", Difficulty: "Hard", Category: "Work"},
	}
}
