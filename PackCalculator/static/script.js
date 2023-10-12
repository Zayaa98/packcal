document.addEventListener("DOMContentLoaded", () => {
    const calculateButton = document.getElementById("calculateButton");
    const orderQuantityInput = document.getElementById("orderQuantity");
    const resultContainer = document.getElementById("result");

    calculateButton.addEventListener("click", async () => {
        resultContainer.textContent = "Calculating...";

        try {
            const orderQuantity = parseInt(orderQuantityInput.value);
            const response = await fetch("/calculate-packs", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(orderQuantity),
            });

            if (response.ok) {
                const data = await response.json();
                displayPacks(data.packsNeeded);
            } else {
                resultContainer.textContent = "Error calculating packs.";
            }
        } catch (error) {
            resultContainer.textContent = "An error occurred.";
        }
    });

    function displayPacks(packsNeeded) {
        resultContainer.innerHTML = "<strong>Packs Needed:</strong>";
        for (const packSize in packsNeeded) {
            resultContainer.innerHTML += `<br>${packsNeeded[packSize]} packs of ${packSize}`;
        }
    }
});
