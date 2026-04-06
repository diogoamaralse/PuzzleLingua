# PuzzleLingua

PuzzleLingua is a translation game for practicing **Portuguese ↔ English** vocabulary.

The project uses:
- a **Go backend** for the API and game logic
- a **React + Vite frontend** for the browser UI

## Features

- Portuguese → English and English → Portuguese rounds
- Easy, Medium, and Hard vocabulary
- Category and difficulty metadata per puzzle
- Answer checking via Go API
- React frontend with score and round tracking

## Tech stack

### Backend
- Go
- net/http
- JSON API

### Frontend
- React
- Vite

## Repository structure

```text
PuzzleLingua/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── data/
│   │   └── words.go
│   ├── game/
│   │   ├── model.go
│   │   ├── normalize.go
│   │   └── service.go
│   └── httpapi/
│       └── handler.go
├── web/
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── src/
│       ├── App.jsx
│       ├── api.js
│       ├── main.jsx
│       └── styles.css
├── go.mod
└── README.md
```

## Requirements

### Backend
- Go 1.26.X or newer

### Frontend
- Node.js 18+ recommended
- npm

## Run the backend

From the project root:

```bash
go run ./cmd/api
```

The API should start on:

```text
http://localhost:8080
```

### Available endpoints

#### Health check
```http
GET /api/health
```

#### Get a random round
```http
GET /api/round
```

#### Check an answer
```http
POST /api/check
Content-Type: application/json
```

Example request body:

```json
{
  "id": 1,
  "answer": "house"
}
```

## Run the frontend

Open a second terminal and go into the `web` folder:

```bash
cd web
npm install
npm run dev
```

The frontend should start on:

```text
http://localhost:5173
```

## Full local development

Run both services at the same time.

### Terminal 1
```bash
go run ./cmd/api
```

### Terminal 2
```bash
cd web
npm install
npm run dev
```

Then open:

```text
http://localhost:5173
```

## Troubleshooting

### Blank screen in the frontend
Check the browser console first. Common causes:
- frontend files not copied correctly
- broken import path in `main.jsx`
- Vite config issue
- backend not running on `http://localhost:8080`

### npm dependency conflict
If `npm install` fails because of Vite/plugin versions, remove the lockfile and modules and reinstall:

```bash
cd web
rm -rf node_modules package-lock.json
npm install
```

### Go import cycle
If you see an import cycle error, make sure:
- `internal/data` can import `internal/game`
- `internal/game` does **not** import `internal/data`
- wiring happens in `internal/httpapi`

## Next improvements

- Move the vocabulary from `words.go` to `words.json`
- Add difficulty filters
- Add hints
- Track high scores
- Add authentication and saved progress
- Add Docker support for frontend and backend

## Current status

This project is currently a working baseline with:
- Go REST API
- React browser UI
- translation rounds loaded from backend data

It is designed to be easy to extend as the game grows.