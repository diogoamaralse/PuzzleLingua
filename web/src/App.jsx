import { useEffect, useState } from "react";
import { fetchRound, checkAnswer } from "./api.js";

const TOTAL_ROUNDS = 5;

export default function App() {
    const [puzzle, setPuzzle] = useState(null);
    const [answer, setAnswer] = useState("");
    const [score, setScore] = useState(0);
    const [round, setRound] = useState(1);
    const [feedback, setFeedback] = useState("Loading...");
    const [loading, setLoading] = useState(true);
    const [checked, setChecked] = useState(false);
    const [gameOver, setGameOver] = useState(false);
    const [lastResultCorrect, setLastResultCorrect] = useState(null);

    useEffect(() => {
        loadRound();
    }, []);

    async function loadRound() {
        setLoading(true);
        setChecked(false);
        setAnswer("");
        setLastResultCorrect(null);

        try {
            const data = await fetchRound();
            setPuzzle(data.puzzle);
            setFeedback("Translate the word shown below.");
        } catch (err) {
            console.error(err);
            setFeedback("Could not load round. Make sure the Go backend is running on http://localhost:8080.");
        } finally {
            setLoading(false);
        }
    }

    async function handleCheck() {
        if (!puzzle || !answer.trim()) {
            setFeedback("Type an answer first.");
            return;
        }

        try {
            const result = await checkAnswer({
                id: puzzle.id,
                answer
            });

            setFeedback(result.message);
            setChecked(true);
            setLastResultCorrect(result.correct);

            if (result.correct) {
                const gained =
                    puzzle.difficulty === "Hard"
                        ? 18
                        : puzzle.difficulty === "Medium"
                            ? 12
                            : 8;

                setScore((prev) => prev + gained);
            }
        } catch (err) {
            console.error(err);
            setFeedback("Could not check answer.");
        }
    }

    async function handleNext() {
        if (round >= TOTAL_ROUNDS) {
            setGameOver(true);
            return;
        }

        setRound((prev) => prev + 1);
        await loadRound();
    }

    function handleRestart() {
        setPuzzle(null);
        setAnswer("");
        setScore(0);
        setRound(1);
        setFeedback("Loading...");
        setLoading(true);
        setChecked(false);
        setGameOver(false);
        setLastResultCorrect(null);
        loadRound();
    }

    function handleKeyDown(event) {
        if (event.key === "Enter" && !loading && !checked && !gameOver) {
            handleCheck();
        }
    }

    const maxScore = TOTAL_ROUNDS * 18;

    if (gameOver) {
        return (
            <main className="app-shell">
                <section className="game-card final-card">
                    <p className="eyebrow">PuzzleLingua</p>
                    <h1>Fim do jogo</h1>
                    <p className="subtitle">
                        Completaste as {TOTAL_ROUNDS} rondas.
                    </p>

                    <div className="final-score-box">
                        <span className="meta-label">Pontuação final</span>
                        <div className="final-score">
                            {score} <span>/ {maxScore}</span>
                        </div>
                    </div>

                    <div className="summary-grid">
                        <article className="meta-item">
                            <span className="meta-label">Rondas jogadas</span>
                            <strong>{TOTAL_ROUNDS}</strong>
                        </article>

                        <article className="meta-item">
                            <span className="meta-label">Pontuação total</span>
                            <strong>{score}</strong>
                        </article>

                        <article className="meta-item">
                            <span className="meta-label">Resultado</span>
                            <strong>
                                {score >= 120
                                    ? "Excelente"
                                    : score >= 80
                                        ? "Muito bom"
                                        : score >= 40
                                            ? "Bom"
                                            : "Continua a praticar"}
                            </strong>
                        </article>
                    </div>

                    <div className="actions">
                        <button onClick={handleRestart}>Jogar novamente</button>
                    </div>
                </section>
            </main>
        );
    }

    return (
        <main className="app-shell">
            <section className="game-card">
                <header className="topbar">
                    <div>
                        <p className="eyebrow">PuzzleLingua</p>
                        <h1>Portuguese ↔ English</h1>
                        <p className="subtitle">
                            Practice translation with a Go backend and React frontend.
                        </p>
                    </div>

                    <div className="stats">
                        <div className="stat-pill">Round {round}/{TOTAL_ROUNDS}</div>
                        <div className="stat-pill">Score {score}</div>
                    </div>
                </header>

                <section className="meta-grid">
                    <article className="meta-item">
                        <span className="meta-label">Direction</span>
                        <strong>{puzzle?.direction ?? "-"}</strong>
                    </article>

                    <article className="meta-item">
                        <span className="meta-label">Difficulty</span>
                        <strong>{puzzle?.difficulty ?? "-"}</strong>
                    </article>

                    <article className="meta-item">
                        <span className="meta-label">Category</span>
                        <strong>{puzzle?.category ?? "-"}</strong>
                    </article>
                </section>

                <section className="prompt-box">
                    <span className="meta-label">Translate</span>
                    <div className="prompt-word">
                        {loading ? "..." : puzzle?.source ?? "-"}
                    </div>
                </section>

                <section className="answer-box">
                    <label htmlFor="answer" className="meta-label">
                        Your answer
                    </label>
                    <input
                        id="answer"
                        type="text"
                        value={answer}
                        onChange={(e) => setAnswer(e.target.value)}
                        onKeyDown={handleKeyDown}
                        placeholder="Type the translation here"
                        disabled={loading || checked}
                    />
                </section>

                <p
                    className={`feedback ${checked ? "checked" : ""} ${
                        lastResultCorrect === true
                            ? "correct"
                            : lastResultCorrect === false
                                ? "incorrect"
                                : ""
                    }`}
                >
                    {feedback}
                </p>

                <div className="actions">
                    <button onClick={handleCheck} disabled={loading || checked}>
                        Check
                    </button>

                    <button onClick={handleNext} disabled={!checked && !loading ? false : !checked}>
                        {round >= TOTAL_ROUNDS ? "Finish game" : "Next"}
                    </button>
                </div>
            </section>
        </main>
    );
}