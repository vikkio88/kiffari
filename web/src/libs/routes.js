import { onMount } from "svelte";
import { navigate } from "svelte-routing";
import { LOGIN_TOKEN_KEY } from "../const";
import { userToken } from "../store";
export const protectedRoute = () => {
    onMount(() => {
        if (!window.localStorage.getItem(LOGIN_TOKEN_KEY)) {
            navigate("/login", { replace: true });
            return;
        }
    });

    userToken.set(window.localStorage.getItem(LOGIN_TOKEN_KEY));
};