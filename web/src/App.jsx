import { useEffect, useState } from "react";
import { fetchRound, checkAnswer } from "./api.js";

export default function App() {
    const [puzzle, setPuzzle] = useState(null);
    const [answer, setAnswer] = useState("");
    const [score, setScore] = useState(0);
    const [round, setRound] = useState(1);
    const [feedback, setFeedback] = useState("Loading...");
    const [loading, setLoading] = useState(true);
    const [checked, setChecked] = useState(false);

    useEffect(() => {
        loadRound();
    }, []);

    async function loadRound() {
        setLoading(true);
        setChecked(false);
        setAnswer("");

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
        setRound((prev) => prev + 1);
        await loadRound();
    }

    function handleKeyDown(event) {
        if (event.key === "Enter" && !loading && !checked) {
            handleCheck();
        }
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
                        <div className="stat-pill">Round {round}</div>
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

                <p className={`feedback ${checked ? "checked" : ""}`}>{feedback}</p>

                <div className="actions">
                    <button onClick={handleCheck} disabled={loading || checked}>
                        Check
                    </button>

                    <button onClick={handleNext}>
                        Next
                    </button>
                </div>
            </section>
        </main>
    );
}