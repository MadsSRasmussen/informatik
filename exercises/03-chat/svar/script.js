// Referencer til html-elementer
const chatInput = document.querySelector("#chat-input");
const chatResetBtn = document.querySelector("#chat-reset-btn");
const chatSendBtn = document.querySelector("#chat-send-btn");

DEBUG = false;

// Konfiguration
const model = "mistral-medium-latest";
let messages = [];

const systemText = `
Du er en max gangster swag chatbot. Du svarer med den sygeste swag og bruger mange emojis i dine svar.

Dit top swaggede svar må du formatere i markdown, hvis det er swag nok til dig.
`.trim();

const systemMessage = AI.createMessage("system", systemText);
messages.push(systemMessage);

// Funktioner
async function sendMessage() {

    // TODO-01: Generér en bruger-besked
    const content = chatInput.value;
    const userMessage = AI.createMessage("user", content);

    // TODO-01.5 -xxx-
    messages.push(userMessage);

    // TODO-02: Opdater UI
    Client.pushMessage(userMessage);
    Client.setLoading();

    chatInput.value = "";

    // TODO-03: Send beskeden til en sprogmodel:
    const responseMessage = await AI.getCompletion(model, messages);

    // TODO-04: Opdater UI
    const responseElement = Client.generateMessageHTML(responseMessage);
    Client.replaceLoader(responseElement);
}

function resetConversation() {
    messages = [];
    messages.push(systemMessage);
    Client.reset();
}

// Events – Denne sektion behøver I ikke at bekymre jer om.
chatSendBtn.addEventListener("click", sendMessage);
chatResetBtn.addEventListener("click", resetConversation);
chatInput.addEventListener("keypress", e => e.code === "Enter" ? sendMessage() : undefined);
