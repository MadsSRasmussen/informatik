const Sandt = (function() {

    const SERVER_URL = "https://informatik.mads-studsgaard.com";

    async function ping() {
        const response = await fetch(`${SERVER_URL}/ping`);
        
        if (response.status === 200) {
            return true
        } else {
            return false
        }
    }

    async function listPosts() {
        const response = await fetch(`${SERVER_URL}/posts`);

        if (response.status === 200) {
            return await response.json()
        } else {
            throw new Error(`HTTP request failed with status: ${response.status}`);
        }
    }

    async function getPost(id) {
        const response = await fetch(`${SERVER_URL}/posts/${id}`);

        if (response.status === 200) {
            return await response.json()
        } else {
            throw new Error(`HTTP request failed with status: ${response.status}`);
        }
    }

    async function createPost(content) {
        const requestBody = {
            content: content
        }

        const response = await fetch(`${SERVER_URL}/posts`, {
            method: "POST",
            body: JSON.stringify(requestBody)
        });

        if (response.status === 201) {
            const responseBody = await response.json();
            return responseBody.id;
        } else {
            throw new Error(`HTTP request failed with status: ${response.status}`);
        }
    }

    async function deletePost(id) {
        const response = await fetch(`${SERVER_URL}/posts/${id}`, { method: "DELETE" });

        if (response.status === 204) {
            return;
        } else {
            throw new Error(`HTTP request failed with status: ${response.status}`);
        }
    }

    return {
        ping,
        listPosts,
        getPost,
        createPost,
        deletePost,
    }

})();
