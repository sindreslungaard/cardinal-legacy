import { writable } from "svelte/store";

export const user = writable(localStorage.getItem("user") ? JSON.parse(localStorage.getItem("user")) : null);