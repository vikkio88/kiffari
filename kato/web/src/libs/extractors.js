import { removeComments } from "./renderers/cleanup";

export function parseLinks(body) {
    const cleanBody = removeComments(body);
    const regex = /(?:^|\n)(\S*)(?:\s+((?:https?:\/\/)?\S+))?(?=\n|$)/g;
    const links = [];
    const matches = cleanBody.matchAll(regex);

    for (const match of matches) {
        const label = (match[1] || '').trim() || (match[2] || '').trim() || (match[3] || '').trim();
        const href = match[2] || match[1];
        links.push({ href, label });
    }
    return links;
}