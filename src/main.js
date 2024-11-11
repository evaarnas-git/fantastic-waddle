async function fetchGoResult() {
    const goResult = await fetch("/api/execute_go").then(res => res.json());
    document.getElementById("result").textContent = "Go Result: " + goResult.result;
}

async function authenticateUser(token) {
    const isValid = await fetch("/api/verifyToken", { method: "POST", body: JSON.stringify({ token }) })
        .then(res => res.json());
    console.log("Token is valid:", isValid);
}

async function hashUserPassword(password) {
    const hash = await fetch("/api/hashPassword", { method: "POST", body: JSON.stringify({ password }) })
        .then(res => res.json());
    console.log("Password hash:", hash);
}

async function calculateProduct(a, b) {
    const product = await fetch("/api/multiply", { method: "POST", body: JSON.stringify({ a, b }) })
        .then(res => res.json());
    console.log("Product:", product);
}
