export function hasCodeSnippets(body) {
    return /```[\s\S]*?```|`[^`\n]+`/.test(body);
}

export const PLUGIN_SETUP_STRING = "<!--\nPlugin: \n-->\n";
export const generatePlugin = (name, content = '') => `<!--\nPlugin: ${name}\n-->\n${content}`;