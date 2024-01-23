<script>
  import { Router, Link, Route, navigate } from "svelte-routing";
  import CreateNote from "./pages/CreateNote.svelte";
  import Dash from "./pages/Dash.svelte";
  import NoteDetail from "./pages/NoteDetail.svelte";
  import EditNote from "./pages/EditNote.svelte";
  import TagDetail from "./pages/TagDetail.svelte";
  import Search from "./pages/Search.svelte";
  import Login from "./pages/Login.svelte";
  import { LOGIN_TOKEN_KEY } from "./const";
  import { isLoggedIn } from "./store";

  let url = "";

  function logout() {
    $isLoggedIn = false;
    window.localStorage.removeItem(LOGIN_TOKEN_KEY);
    navigate("/login", { replace: true });
  }
</script>

<Router {url}>
  <nav>
    <Link to="/" title="Dashboard">🪣</Link>
    <Link to="/create-note" title="Add New">➕</Link>
    <Link to="/search" title="Search">🔍</Link>
    {#if $isLoggedIn}
      <button on:click={logout} title="Logout">👋</button>
    {/if}
  </nav>
  <main>
    <Route path="/login" component={Login} />
    <Route path="/tags/:id" let:params>
      <TagDetail id={params.id} />
    </Route>
    <Route path="/create-note" component={CreateNote} />
    <Route path="/edit-note/:id" let:params>
      <EditNote id={params.id} />
    </Route>
    <Route path="/notes/:id" let:params>
      <NoteDetail id={params.id} />
    </Route>
    <Route path="/search" component={Search} />
    <Route path="/" component={Dash} />
  </main>
</Router>

<style>
  nav > button {
    position: absolute;
    right: 0;
    margin: 1rem 1rem 0 0;
  }
</style>
