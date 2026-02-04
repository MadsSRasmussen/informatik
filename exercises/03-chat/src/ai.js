let DEBUG = true;

const AI = (function() {
    const SERVER_URL = "https://informatik.mads-studsgaard.com";

    const converter = new MarkdownConverter();

    const aiStatusElement = document.querySelector("#ai-status");

    let pending = false;

    function createMessage(role, content) {
        const message = {
            role: role,
            content: content,
        }

        return message;
    }

    async function getCompletion(model, messages) {
        try {
            if (pending) return;
            pending = true;

            if (DEBUG) aiStatusElement.style.visibility = "visible";

            let requestBody;
            if (Array.isArray(messages)) {
                requestBody = {
                    model: model,
                    messages: messages,
                }
            } else {
                requestBody = {
                    model: model,
                    messages: [messages],
                }
            }

            const response = await fetch(`${SERVER_URL}/completions`, {
                method: "POST",
                body: JSON.stringify(requestBody),
            });

            if (!response.ok) {
                throw new Error(`Completions request failed with status-code: ${response.status}`);
            }

            const body = await response.json();

            const assistantMessage = {
                role: body.role,
                content: converter.toHTMLString(body.content),
            }

            aiStatusElement.style.visibility = "hidden";
            pending = false;

            return assistantMessage;
        } catch (error) {

            aiStatusElement.style.visibility = "hidden";
            pending = false;

            throw error;
        } 

    }

    return {
        createMessage,
        getCompletion,
    }
})();
