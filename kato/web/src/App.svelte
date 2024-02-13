<script>
  import { Router, Route } from "svelte-routing";
  import Nav from "./components/Nav.svelte";
  import CreateNote from "./pages/CreateNote.svelte";
  import Dash from "./pages/Dash.svelte";
  import NoteDetail from "./pages/NoteDetail.svelte";
  import EditNote from "./pages/EditNote.svelte";
  import TagDetail from "./pages/TagDetail.svelte";
  import Search from "./pages/Search.svelte";
  import Login from "./pages/Login.svelte";
  import Archived from "./pages/Archived.svelte";
  import Footer from "./components/Footer.svelte";

  let url = "";
</script>

<Router {url}>
  <Nav />
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
    <Route path="/archived" component={Archived} />
  </main>
  <Footer />
</Router>
