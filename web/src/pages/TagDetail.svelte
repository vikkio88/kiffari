<script>
  import NoteList from "../components/notes/NoteList.svelte";
  import Header from "../components/shared/Header.svelte";
  import { getTagDetails, parseOrThrow, catchLogout } from "../libs/api";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();

  export let id = "";
  let tagPromise = getTagDetails(id).then(parseOrThrow).catch(catchLogout);
</script>

{#await tagPromise then tag}
  <div class="header">
    <h1 class="hash">#</h1>
    <h1>
      {tag.label}
    </h1>
  </div>
  <div class="results">
    <Header text="Notes" />
    <NoteList notes={tag.notes} />
  </div>
{/await}

<style>
  .header {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }
  .header > h1 {
    margin: 0;
    font-size: 3rem;
  }

  h1.hash {
    color: var(--link-color);
    font-size: 3.2rem;
    margin-right: 0.3rem;
  }

  .results {
    margin-top: 5rem;
  }
</style>
