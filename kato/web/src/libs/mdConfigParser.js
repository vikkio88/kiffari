/**
 * @param {String} body 
 * @returns {object} config
 */
export function mdConfigParser(body) {
    const results = body.match(/<\!--(.+?)-->/s);
    if (!Boolean(results) || results.length < 2) {
        return [];
    }
    const lines = results[1].trim().split("\n");
    const config = {};
    for (const line of lines) {
        const matches = line.match(/\s*^(.+?):\s?(.+?)$/);
        if (Boolean(matches) && matches.length > 2) {
            config[matches[1]] = parseValue(matches[2]);
        }
    }

    return config;
}

const parseValue = (value) => !isNaN(parseFloat(value)) ? (parseFloat(value) % 1 === 0 ? (parseInt(value) === 1 || parseInt(value) === 0 ? parseInt(value) === 1 : parseInt(value)) : parseFloat(value)) : (value.toLowerCase() === 'true' || value.toLowerCase() === 'false' ? value.toLowerCase() === 'true' : value);