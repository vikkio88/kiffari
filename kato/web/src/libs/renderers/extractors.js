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