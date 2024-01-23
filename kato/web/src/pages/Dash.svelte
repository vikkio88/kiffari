<script>
  import { KATO_API_URL } from "../const";
  import NoteList from "../components/NoteList.svelte";
  import { protectedRoute } from "../libs";
  protectedRoute();

  let notePromise = fetch(`${KATO_API_URL}/notes?latest=true`).then((resp) =>
    resp.json(),
  );
</script>

<h2>Latest Notes</h2>
{#await notePromise then notes}
  <NoteList {notes} />
{/await}
