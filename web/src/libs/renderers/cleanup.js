/**
 * @param {String} body 
 */
export function removeComments(body) {
    return body.replaceAll(/<\!--.+?-->/sg, "").trim();
}