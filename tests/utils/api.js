const API_URL = "http://localhost:5111/api";

let token = null;

export function basicHeaders(additional = {}) {

    if (Boolean(token)) {
        additional = {
            ...additional,
            Authorization: token,
        };
    }
    return {
        Accept: "application/json",
        "Content-Type": "application/json",
        ...additional
    };
}

export async function login() {
    const resp = await post('login', { passKey: "password" });
    if (!resp.ok) {
        throw Error("Could not login");
    }

    const { token: newToken } = await resp.json();
    token = newToken;
}

export function logout() {
    token = null;
}

export function get(api, headers = {}) {
    return fetch(`${API_URL}/${api}`, { method: "GET", headers: { ...basicHeaders(), ...headers } });
}

export function put(api, body = {}, headers = {}) {
    return fetch(`${API_URL}/${api}`, { method: "PUT", headers: { ...basicHeaders(), ...headers }, body: JSON.stringify(body) });
}

export function del(api, headers = {}) {
    return fetch(`${API_URL}/${api}`, { method: "DELETE", headers: { ...basicHeaders(), ...headers } });
}

export function post(api, body = {}, headers = {}) {
    return fetch(`${API_URL}/${api}`, { method: "POST", headers: { ...basicHeaders(), ...headers }, body: JSON.stringify(body) });
} 