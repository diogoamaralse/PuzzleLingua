const API_BASE = "http://localhost:8080";

export async function fetchRound() {
    const res = await fetch(`${API_BASE}/api/round`);

    if (!res.ok) {
        throw new Error(`Failed to load round: ${res.status}`);
    }

    return res.json();
}

export async function checkAnswer(payload) {
    const res = await fetch(`${API_BASE}/api/check`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
    });

    if (!res.ok) {
        throw new Error(`Failed to check answer: ${res.status}`);
    }

    return res.json();
}