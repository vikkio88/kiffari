<script>
  import { Router, Route } from "svelte-routing";
  import Nav from "./components/Nav.svelte";
  import CreateNote from "./pages/CreateNote.svelte";
  import KatoDash from "./pages/KatoDash.svelte";
  import NoteDetail from "./pages/NoteDetail.svelte";
  import EditNote from "./pages/EditNote.svelte";
  import TagDetail from "./pages/TagDetail.svelte";
  import Search from "./pages/Search.svelte";
  import Login from "./pages/Login.svelte";
  import Archived from "./pages/Archived.svelte";
  import About from "./pages/About.svelte";
  import Fallback from "./pages/Fallback.svelte";
  import Spinner from "./components/shared/Spinner.svelte";
  import { catchLogout, getConfig, parseOrThrow } from "./libs/api";
  import ProjectDetails from "./pages/ProjectDetails.svelte";
  import { appConfig } from "./store";
  import Main from "./pages/Main.svelte";
  import KiffariDash from "./pages/KiffariDash.svelte";

  let url = "";

  //TODO: this could also check validity of token
  let configPromise = getConfig()
    .then(parseOrThrow)
    .then((config) => {
      if (!config.auth) {
        catchLogout();
      }
      delete config.auth;
      $appConfig = config;
      return config;
    });
</script>

{#await configPromise}
  <div class="frc">
    <Spinner />
  </div>
{:then config}
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
      {#if config.kiffari}
        <Route path="/" component={Main} />
        <Route path="/kiffari" component={KiffariDash} />
        <Route path="/projects/:id" let:params>
          <ProjectDetails id={params.id} />
        </Route>
        <Route path="/kato" component={KatoDash} />
      {:else}
        <Route path="/kato" component={KatoDash} />
        <Route path="/" component={KatoDash} />
      {/if}
      <Route path="/archived" component={Archived} />
      <Route path="/about" component={About} />
      <Route path="*" component={Fallback} />
    </main>
  </Router>
{/await}
