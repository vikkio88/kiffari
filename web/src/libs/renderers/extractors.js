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