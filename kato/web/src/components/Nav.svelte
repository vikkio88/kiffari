<script>
    import { Link, navigate } from "svelte-routing";
    import { LOGIN_TOKEN_KEY } from "../const";
    import { userToken } from "../store";

    function logout() {
        $userToken = null;
        window.localStorage.removeItem(LOGIN_TOKEN_KEY);
        navigate("/login", { replace: true });
    }
</script>

<nav>
    <Link to="/" title="Dashboard">🪣</Link>
    <Link to="/archived" title="Archived">🗄️</Link>
    <Link to="/create-note" title="Add New">➕</Link>
    <Link to="/search" title="Search">🔍</Link>
    {#if $userToken !== null}
        <button on:click={logout} title="Logout">👋</button>
    {/if}
</nav>

<style>
    nav > button {
        position: absolute;
        right: 0;
        margin: 1rem 1rem 0 0;
    }
</style>
