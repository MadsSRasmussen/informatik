// Her kan du skrive din kode, du har adgang til objektet AI i denne fil.

const chatInput = document.querySelector("#chat-input");
const chatSendBtn = document.querySelector("#chat-send-btn");

chatInput.addEventListener("keypress", (e) => {
    if (e.code === "Enter") {
        handleMessageSend();
    }
});

chatSendBtn.addEventListener("click", () => {
    handleMessageSend();
})

function handleMessageSend() {
    const content = chatInput.value;
    if (!content) {
        return;
    }
    AI.sendMessage(content);
    chatInput.value = "";
}
