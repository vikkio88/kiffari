import { KATO_API_URL, LOGIN_TOKEN_KEY } from "../const";

function getAuthToken() {
  return window.localStorage.getItem(LOGIN_TOKEN_KEY) ?? false;
}

export function createNote(note) {
  return fetch(`${KATO_API_URL}/notes`, {
    method: "POST",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(note),
  });
}

export function getLatestNotes() {
  let token = null;
  if (!(token = getAuthToken())) {
    return
  }
  return fetch(`${KATO_API_URL}/notes?latest=true`, {
    headers: { Authorization: `${token}` },
  });
}

export function getNoteDetails(id) {
  return fetch(`${KATO_API_URL}/notes/${id}`);
}

export function updateNote(id, note) {
  return fetch(`${KATO_API_URL}/notes/${id}`, {
    method: "PUT",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(note),
  });
}

export function deleteNote(id) {
  return fetch(`${KATO_API_URL}/notes/${id}`, {
    method: "DELETE",
  });
}
