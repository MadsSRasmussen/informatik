const Client = (function() {
    const messagesContainer = document.querySelector(".messages-container");
    const chatContainer = document.querySelector(".chat-container");

    let loader = null;

    function generateMessageHTML(message) {
        const messageElement = document.createElement("div");
        messageElement.classList.add("message");
        messageElement.setAttribute("role", message.role);

        const text = document.createElement("div");
        text.classList.add("message-content");
        text.innerHTML = message.content; 

        messageElement.appendChild(text);

        return messageElement;
    }

    function pushMessage(message) {
        const messageElement = generateMessageHTML(message);

        messagesContainer.appendChild(messageElement);

        if (message.role !== "system") {
            messagesContainer.classList.add("has-messages");
            chatContainer.classList.add("has-messages");
        }

        if (message.role === "user") {
            setScroll();
        }
    }

    function setLoading() {
        if (loader !== null) {
            return;
        }

        loader = generateMessageHTML({ 
            role: "assistant", 
            content: "<span class='loader'></span>" 
        });

        messagesContainer.appendChild(loader);
        setScroll();
    }

    function replaceLoader(element) {
        if (loader === null) {
            return;
        } 

        loader.replaceWith(element);
        loader = null;
        setScroll();
    }

    function reset() {
        messagesContainer.innerHTML = "";

        messagesContainer.classList.remove("has-messages");
        chatContainer.classList.remove("has-messages");
    }

    function setScroll() {
        messagesContainer.scrollTo({
            top: messagesContainer.scrollHeight,
            behavior: "smooth",
        })
    }

    return {
        generateMessageHTML,
        pushMessage,
        reset,
        setLoading,
        replaceLoader,
    }
})();
