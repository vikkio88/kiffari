import { onMount } from "svelte";
import { navigate } from "svelte-routing";
import { LOGIN_TOKEN_KEY } from "./const";
export const protectedRoute = () => {
    onMount(() => {
        if (!window.localStorage.getItem(LOGIN_TOKEN_KEY)) {
            navigate("/login", { replace: true });
        }
    });
};