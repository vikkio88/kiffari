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

export function getDashNotes() {
  return fetch(`${KATO_API_URL}/dash/notes`, makeHeaders()).then(parseOrThrow).catch(catchLogout);
}

export function getAllNotes() {
  return fetch(`${KATO_API_URL}/notes`, makeHeaders()).then(parseOrThrow).catch(catchLogout);
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

export function archive(resource = "notes", id) {
  return fetch(`${KATO_API_URL}/${resource}/${id}/archive`, {
    method: "POST",
    ...makeHeaders(),
  });
}

export function unarchive(resource = "notes", id) {
  return fetch(`${KATO_API_URL}/${resource}/${id}/archive`, {
    method: "DELETE",
    ...makeHeaders(),
  });
}

export function pinNote(id) {
  return fetch(`${KATO_API_URL}/notes/${id}/pin`, {
    method: "POST",
    ...makeHeaders(),
  });
}

export function unpinNote(id) {
  return fetch(`${KATO_API_URL}/notes/${id}/pin`, {
    method: "DELETE",
    ...makeHeaders(),
  });
}

export function del(resource = "notes", id) {
  return fetch(`${KATO_API_URL}/${resource}/${id}`, {
    method: "DELETE",
    ...makeHeaders(),
  });
}

export function getTagDetails(id) {
  return fetch(`${KATO_API_URL}/tags/${id}`, makeHeaders());
}

export function filterTags(value, abortCtrl) {
  return fetch(`${KATO_API_URL}/tags?q=${value}`, {
    method: "GET",
    signal: abortCtrl.signal,
    ...makeHeaders(),
  }).then(parseOrThrow)
    .catch(catchLogout);
}

export function trendingTags() {
  return fetch(`${KATO_API_URL}/tags?trending=true`, {
    method: "GET",
    ...makeHeaders(),
  }).then(parseOrThrow).catch(catchLogout);
}

export function filterNotes(value, abortCtrl) {
  return fetch(`${KATO_API_URL}/notes?q=${value}`, {
    method: "GET",
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

export function getProjectWithArchivedTasks(id) {
  return fetch(`${KATO_API_URL}/projects/${id}?archived=true`, { ...makeHeaders() }).then(parseOrThrow).catch(catchLogout);
}

export function createProject(project) {
  return fetch(`
  ${KATO_API_URL}/projects`,
    {
      method: "POST", ...makeHeaders(),
      body: JSON.stringify(project)
    }
  );
}

export function updateProject(id, project) {
  return fetch(`
  ${KATO_API_URL}/projects/${id}`,
    {
      method: "PUT", ...makeHeaders(),
      body: JSON.stringify(project)
    }
  ).then(parseOrThrow).catch(catchLogout);
}

export function addTask(projectId, task) {
  return fetch(`${KATO_API_URL}/projects/${projectId}/task`, { method: "POST", ...makeHeaders(), body: JSON.stringify(task) });
}

export function updateTask(projectId, task) {
  return fetch(`${KATO_API_URL}/projects/${projectId}/task`, { method: "PUT", ...makeHeaders(), body: JSON.stringify(task) });
}

export function getTask(taskId) {
  return fetch(`${KATO_API_URL}/tasks/${taskId}`, { method: "GET", ...makeHeaders() }).then(parseOrThrow).catch(catchLogout);
}

export function moveTask(projectId, taskId, status) {
  return fetch(`${KATO_API_URL}/projects/${projectId}/task/${taskId}/status`, { method: "PUT", ...makeHeaders(), body: JSON.stringify({ status }) });
}