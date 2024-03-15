/**
 * @param {String} body 
 */
export function removeComments(body) {
    return body.replaceAll(/<\!--.+?-->/sg, "").trim();
}

export function removeMd(body) {
    return removeComments(body).replaceAll(/[^\d\w\s]/sg,"").trim()
}


export function previewMd(body) {
    const PREVIEW_MAX_LENGTH = 30;
    const newBody = removeMd(body);

    if(newBody.length > PREVIEW_MAX_LENGTH) {
        return `${newBody.slice(0, PREVIEW_MAX_LENGTH)}...`
    }

    return newBody;
}