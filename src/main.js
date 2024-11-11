async function fetchGoResult() {
    const goResult = await fetch("/api/execute_go").then(res => res.json());
    document.getElementById("result").textContent = "Go Result: " + goResult.result;
}
