export function hasCodeSnippets(body) {
    return /```[\s\S]*?```|`[^`\n]+`/.test(body);
}