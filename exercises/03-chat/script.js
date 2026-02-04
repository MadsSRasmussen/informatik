// Referencer til html-elementer
const chatInput = document.querySelector("#chat-input");
const chatResetBtn = document.querySelector("#chat-reset-btn");
const chatSendBtn = document.querySelector("#chat-send-btn");

// Konfiguration
const model = "mistral-small-latest";
let messages = [];

// Funktioner
async function sendMessage() {
    console.log("Vi prøver at sende en besked.");

    // TODO-01: Generér en bruger-besked

    // TODO-01.5 -xxx-

    // TODO-02: Opdater UI

    // TODO-03: Send beskeden til en sprogmodel:

    // TODO-04: Opdater UI
}

function resetConversation() {
    console.log("Vi prøver at starter samtalen forfra.");

    // TODO-05: Genstart samtalen
}

// Events – Denne sektion behøver I ikke at bekymre jer om.
chatSendBtn.addEventListener("click", sendMessage);
chatResetBtn.addEventListener("click", resetConversation);
chatInput.addEventListener("keypress", e => e.code === "Enter" ? sendMessage() : undefined);
