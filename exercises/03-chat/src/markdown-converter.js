class MarkdownConverter {

    /**
     * @typedef {Object} ConvertionRule
     * @property {RegEx} regex
     * @property {string | Function} replacement
     */

    constructor() {

        /** @type {ConvertionRule[]} */
        this.rules = [
            { regex: /^######[ \t]+([^\n]+)$/gm, replacement: "<h6>$1</h6>" },
            { regex: /^#####[ \t]+([^\n]+)$/gm, replacement: "<h5>$1</h5>" },
            { regex: /^####[ \t]+([^\n]+)$/gm, replacement: "<h4>$1</h4>" },
            { regex: /^###[ \t]+([^\n]+)$/gm, replacement: "<h3>$1</h3>" },
            { regex: /^##[ \t]+([^\n]+)$/gm, replacement: "<h2>$1</h2>" },
            { regex: /^#[ \t]+([^\n]+)$/gm, replacement: "<h1>$1</h1>" },
            { regex: /^\s*(?:---|\*\*\*|___)\s*$/gm, replacement: "<hr>" },
            { regex: /\*\*(.*?)\*\*/g, replacement: "<strong>$1</strong>" },
            { regex: /\*(.*?)\*/g, replacement: "<em>$1</em>" },
            { regex: /`(.*?)`/g, replacement: "<code>$1</code>" },
            { regex: /\[(.*?)\]\((.*?)\)/g, replacement: '<a href="$2">$1</a>' },
        ];

        // Code block processing
        /** @type {ConvertionRule} */
        this.codeBlockRule = {
            regex: /```(\w+)?\n([\s\S]*?)```/g,
            replacement: (match, language, code) => {
            const langClass = language ? ` class="language-${language} code-block"` : "";
            const encodedCode = btoa(
                `<pre><code${langClass}>${code.trim()}</code></pre>`
            );
            return `{{CODEBLOCK:${encodedCode}}}`;
            },
        };

        this.listRule = {
            regex: /(^|\n)\s*([-*+] .*(\n|$))+/gm, // Match blocks of unordered list items
            replacement: (match) => {
                const items = match
                    .trim()
                    .split("\n")
                    .map(item => item.replace(/^[-*+] (.*)/, "<li>$1</li>"))
                    .join("");
                return `<ul>${items}</ul>`;
            },
        };

        this.orderedListRule = {
            regex: /(^|\n)(\d+\. .*(\n\s*)*)+/gm, // Allow for blank lines within the list
            replacement: (match) => {
                const items = match
                    .trim()
                    .split(/\n+/) // Split on any sequence of newline characters
                    .filter(item => item.trim() !== "") // Ignore empty lines
                    .map(item => item.replace(/^\d+\. (.*)/, "<li>$1</li>")) // Wrap list items in <li>
                    .join("");
                return `<ol>${items}</ol>`; // Wrap the result in <ol>
            },
        };

    }

    /**
     * 
     * @param {string} markdown 
     * @returns 
     */
    toHTMLString(markdown) {
        // 1. Code blocks (protected)
        markdown = markdown.replace(this.codeBlockRule.regex, this.codeBlockRule.replacement);

        // 2. Horizontal rules
        markdown = markdown.replace(
            /^\s*(?:---|\*\*\*|___)\s*$/gm,
            "<hr>"
        );

        // 3. Headers (block-level, line-based)
        markdown = markdown.replace(/^######[ \t]+([^\n]+)$/gm, "<h6>$1</h6>");
        markdown = markdown.replace(/^#####[ \t]+([^\n]+)$/gm, "<h5>$1</h5>");
        markdown = markdown.replace(/^####[ \t]+([^\n]+)$/gm, "<h4>$1</h4>");
        markdown = markdown.replace(/^###[ \t]+([^\n]+)$/gm, "<h3>$1</h3>");
        markdown = markdown.replace(/^##[ \t]+([^\n]+)$/gm, "<h2>$1</h2>");
        markdown = markdown.replace(/^#[ \t]+([^\n]+)$/gm, "<h1>$1</h1>");

        // 4. Lists (ONLY now)
        markdown = markdown.replace(this.orderedListRule.regex, this.orderedListRule.replacement);
        markdown = markdown.replace(this.listRule.regex, this.listRule.replacement);

        // 5. Inline formatting (safe after blocks)
        markdown = markdown.replace(/\*\*(.*?)\*\*/g, "<strong>$1</strong>");
        markdown = markdown.replace(/\*(.*?)\*/g, "<em>$1</em>");
        markdown = markdown.replace(/`(.*?)`/g, "<code>$1</code>");
        markdown = markdown.replace(/\[(.*?)\]\((.*?)\)/g, '<a href="$2">$1</a>');

        // 6. Line breaks LAST
        markdown = markdown.replace(/\n/g, "<br>");

        // 7. Restore code blocks
        markdown = markdown.replace(/{{CODEBLOCK:(.*?)}}/g, (_, encoded) =>
            atob(encoded)
        );

        return markdown;
    }

}
