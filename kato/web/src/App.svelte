<script>
  import { Router, Link, Route } from "svelte-routing";
  import CreateNote from "./pages/CreateNote.svelte";
  import Dash from "./pages/Dash.svelte";
  let url = ""
</script>


  <Router {url}>
    <nav>
      <Link to="/">Dash</Link>
      <Link to="/create-note">Create</Link>
    </nav>
    <main>
      <Route path="/create-note" component={CreateNote} />
      <Route path="/" component={Dash} />
    </main>
  </Router>

