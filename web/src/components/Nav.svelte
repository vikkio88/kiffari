<script>
  import { Link, navigate } from "svelte-routing";
  import { LOGIN_TOKEN_KEY } from "../const";
  import { appConfig, userToken } from "../store";

  function logout() {
    $userToken = null;
    window.localStorage.removeItem(LOGIN_TOKEN_KEY);
    navigate("/login", { replace: true });
  }
</script>

{#if $userToken !== null}
  <nav>
    {#if !$appConfig.kiffari}
      <Link to="/" title="Dashboard">🪣</Link>
    {:else}
      <Link to="/" title="Home">🏠</Link>
      <Link to="/kato" title="Kato">🪣</Link>
    {/if}
    <Link to="/archived" title="Archived">🗄️</Link>
    <Link to="/search" title="Search">🔍</Link>
    <button on:click={logout} title="Logout" class="logout">👋</button>
  </nav>
{/if}

<style>
  nav > button.logout {
    position: absolute;
    right: 0;
    margin: 1rem 1rem 0 0;
  }
</style>
