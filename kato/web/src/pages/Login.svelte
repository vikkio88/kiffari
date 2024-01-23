<script>
    import { navigate } from "svelte-routing";
    import { LOGIN_TOKEN_KEY } from "../const";
    import { isLoggedIn } from "../store";

    function login() {
        window.localStorage.setItem(LOGIN_TOKEN_KEY, "user");
        $isLoggedIn = true;
        navigate("/", { replace: true });
    }
</script>

<button on:click={login}>Login</button>
