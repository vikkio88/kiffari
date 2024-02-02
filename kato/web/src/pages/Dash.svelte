<script>
  import NoteList from "../components/NoteList.svelte";
  import { catchLogout, getLatestNotes, parseOrThrow } from "../libs/api";
  import { protectedRoute } from "../libs/routes";

  protectedRoute();
  let notePromise = getLatestNotes()
    .then(parseOrThrow)
    .catch(catchLogout);
</script>

<h2>Latest Notes</h2>
{#await notePromise then notes}
  <NoteList {notes} />
{/await}
