import { writable } from "svelte/store";

export const userToken = writable(null);


export const appConfig = writable({ kiffari: false });