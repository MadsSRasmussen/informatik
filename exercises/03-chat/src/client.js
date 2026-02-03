const AI = (function() {
    const messagesContainer = document.querySelector(".messages-container");
    const chatContainer = document.querySelector(".chat-container");

    function generateMessage(content, role) {
        const message = document.createElement("div");
        message.classList.add("message");
        message.setAttribute("role", role);

        const text = document.createElement("div");
        text.classList.add("message-content");
        text.innerText = content; 

        message.appendChild(text);

        return message;
    }

    async function sendMessage(content) {
        const message = generateMessage(content, "user");
        messagesContainer.appendChild(message);

        messagesContainer.classList.add("has-messages");
        chatContainer.classList.add("has-messages");
    }

    return {
        sendMessage,
    }
})();
