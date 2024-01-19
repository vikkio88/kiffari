<script>
  import { Router, Link, Route } from "svelte-routing";
  import CreateNote from "./pages/CreateNote.svelte";
  import Dash from "./pages/Dash.svelte";
  import NoteDetail from "./pages/NoteDetail.svelte";
  import EditNote from "./pages/EditNote.svelte";
  let url = "";
</script>

<Router {url}>
  <nav>
    <Link to="/">Dash</Link>
    <Link to="/create-note">Create</Link>
  </nav>
  <main>
    <Route path="/create-note" component={CreateNote} />
    <Route path="/edit-note/:id" let:params>
      <EditNote id={params.id} />
    </Route>
    <Route path="/notes/:id" let:params>
      <NoteDetail id={params.id} />
    </Route>
    <Route path="/" component={Dash} />
  </main>
</Router>
