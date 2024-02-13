<script>
  import { Link, navigate } from "svelte-routing";
  import { LOGIN_TOKEN_KEY } from "../const";
  import { userToken } from "../store";

  function logout() {
    $userToken = null;
    window.localStorage.removeItem(LOGIN_TOKEN_KEY);
    navigate("/login", { replace: true });
  }

  function create() {
    navigate("/create-note");
  }
</script>

{#if $userToken !== null}
  <nav>
    <Link to="/" title="Dashboard">🪣</Link>
    <Link to="/archived" title="Archived">🗄️</Link>
    <Link to="/search" title="Search">🔍</Link>
    <button class="add" title="New Note" on:click={create}>➕</button>
    <button on:click={logout} title="Logout" class="logout">👋</button>
  </nav>
{/if}

<style>
  nav > button.logout {
    position: absolute;
    right: 0;
    margin: 1rem 1rem 0 0;
  }

  .add {
    font-size: 1.2rem;
    padding: 0.3rem 1rem;
  }
</style>
