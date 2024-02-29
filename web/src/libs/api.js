import { navigate } from "svelte-routing";
import { KATO_API_URL, LOGIN_TOKEN_KEY } from "../const";
import { userToken } from "../store";

export function parseOrThrow(resp) {
  if (resp.status == 401) {
    throw new Error("unauth");
  }
  return resp.json();
}

export function catchLogout() {
  window.localStorage.removeItem(LOGIN_TOKEN_KEY);
  userToken.set(null);
  navigate("/login", { replace: true });
}

export function login(passkey) {
  return fetch(`${KATO_API_URL}/login`, {
    method: 'POST',
    ...makeHeaders(),
    body: JSON.stringify({ passkey })
  });
}

function getAuthToken() {
  return window.localStorage.getItem(LOGIN_TOKEN_KEY) ?? null;
}

function makeHeaders(others = {}) {
  const auth = { Authorization: getAuthToken() };
  return {
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      ...auth,
      ...others,
    }
  };
}

export function getConfig() {
  return fetch(`${KATO_API_URL}/config`, makeHeaders());
}

export function createNote(note) {
  return fetch(`${KATO_API_URL}/notes`, {
    method: "POST",
    ...makeHeaders(),
    body: JSON.stringify(note),
  });
}

export function getLatestNotes() {
  return fetch(`${KATO_API_URL}/notes?latest=true`, makeHeaders()).then(parseOrThrow).catch(catchLogout);
}

export function getReminderNotes() {
  return fetch(`${KATO_API_URL}/notes?reminders=true`, makeHeaders());
}

export function getArchivedNotes() {
  return fetch(`${KATO_API_URL}/notes?archived=true`, makeHeaders()).then(parseOrThrow).catch(catchLogout);
}

export function getNoteDetails(id) {
  return fetch(`${KATO_API_URL}/notes/${id}`, makeHeaders()).then(parseOrThrow).catch(catchLogout);
}

export function updateNote(id, note) {
  return fetch(`${KATO_API_URL}/notes/${id}`, {
    method: "PUT",
    ...makeHeaders(),
    body: JSON.stringify(note),
  });
}

export function archiveNote(id) {
  return fetch(`${KATO_API_URL}/notes/${id}/archive`, {
    method: "POST",
    ...makeHeaders(),
  });
}

export function unArchiveNote(id) {
  return fetch(`${KATO_API_URL}/notes/${id}/archive`, {
    method: "DELETE",
    ...makeHeaders(),
  });
}

export function deleteNote(id) {
  return fetch(`${KATO_API_URL}/notes/${id}`, {
    method: "DELETE",
    ...makeHeaders(),
  });
}

export function getTagDetails(id) {
  return fetch(`${KATO_API_URL}/tags/${id}`, makeHeaders());
}

export function filterTags(value, abortCtrl) {
  return fetch(`${KATO_API_URL}/tags?q=${value}`, {
    method: "get",
    signal: abortCtrl.signal,
    ...makeHeaders(),
  }).then(parseOrThrow)
    .catch(catchLogout);
}

export function trendingTags() {
  return fetch(`${KATO_API_URL}/tags?trending=true`, {
    method: "get",
    ...makeHeaders(),
  }).then(parseOrThrow).catch(catchLogout);
}

export function filterNotes(value, abortCtrl) {
  return fetch(`${KATO_API_URL}/notes?q=${value}`, {
    method: "get",
    signal: abortCtrl.signal,
    ...makeHeaders(),
  }).then(parseOrThrow)
    .catch(catchLogout);
}

export function getProjects() {
  return fetch(`${KATO_API_URL}/projects`, { ...makeHeaders() }).then(parseOrThrow).catch(catchLogout);
}

export function getProject(id) {
  return fetch(`${KATO_API_URL}/projects/${id}`, { ...makeHeaders() }).then(parseOrThrow).catch(catchLogout);
}

export function addTask(projectId, task) {
  return fetch(`${KATO_API_URL}/projects/${projectId}/task`, { method: "POST", ...makeHeaders(), body: JSON.stringify(task) });
}

export function moveTask(projectId, taskId, status) {
  return fetch(`${KATO_API_URL}/projects/${projectId}/task/${taskId}/status`, { method: "PUT", ...makeHeaders(), body: JSON.stringify({ status }) });
}