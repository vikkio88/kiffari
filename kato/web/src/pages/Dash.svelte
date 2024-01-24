<script>
  import NoteList from "../components/NoteList.svelte";
  import { getLatestNotes } from "../libs/api";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();
  let notePromise = getLatestNotes().then((resp) => resp.json());
</script>

<h2>Latest Notes</h2>
{#await notePromise then notes}
  <NoteList {notes} />
{/await}
