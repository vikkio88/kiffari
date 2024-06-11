import { removeComments } from "./cleanup";

export function parseLinks(body) {
    const cleanBody = removeComments(body);
    const regex = /(.+)\s*(https?:\/\/\S+)/g;
    const links = [];
    const matches = cleanBody.matchAll(regex);

    for (const match of matches) {
        const label = (match[1] ?? match[2]).trim();
        const href = match[2];
        links.push({ href, label });
    }
    return links;
}

export function extractYTVideos(body) {
    const cleanBody = removeComments(body);
    const embedBaseUrl = "https://www.youtube.com/embed/";

    const ytVideos = [];
    const urls = cleanBody.split("\n");

    for (const original of urls) {
        const pattern = /(?:https?:\/\/)?(?:www\.)?youtu(?:\.be\/|be\.com\/(?:watch\?.*?v=|.*\/|embed\/|v\/))([^&?/#]+)/;
        const match = original.match(pattern);
        const videoId = match ? match[1] : null;

        if (!videoId) continue;

        ytVideos.push({
            original: original,
            embed: embedBaseUrl + videoId
        });
    }
    return ytVideos;
}

export function extractTodos(body) {
    const cleanBody = removeComments(body);
    const todos = [];
    const lines = cleanBody.split("\n");
    for (const line of lines) {
        const match = line.match(/-\s*\[\s*(x?)\s*\]\s*(.+)?/);
        if (!match) continue;
        const [, doneStatus, rawLabel] = match;
        const label = rawLabel.trim();
        const done = doneStatus.trim().toLowerCase() === 'x';
        todos.push({ label, done });
    }

    return todos;
}

export function exportTodos(todos) {
    let todoLines = [];
    for (const todo of todos) {
        todoLines.push(`- [${todo.done ? "x" : ""}] ${todo.label}`);
    }
    return `<!--
plugin: todo
-->
${todoLines.join("\n")}
`;
}

export function extractTrackers(body) {
    const cleanBody = removeComments(body);
    const regex = /\s?(\d+)\/(\d+):?(\d+)? ?(.+)?/;
    const lines = cleanBody.split("\n");
    const trackers = [];
    for (const line of lines) {
        const matches = line.match(regex);
        if (!matches) continue;
        const [_, value = 1, max, step = 1, label = null] = matches;
        trackers.push({
            label, range: { value: parseInt(`${value}`), max: parseInt(max), min: 0, step: parseInt(`${step}`) }
        });
    }

    return trackers;
}

export function exportTrackers(trackers) {
    let markdown = "";
    for (const tracker of trackers) {
        markdown += `${tracker.range.value}/${tracker.range.max}:${tracker.range.step} ${tracker.label}\n`;
    }
    return `<!--
Plugin: tracker
-->
${markdown.trim()}`;
}