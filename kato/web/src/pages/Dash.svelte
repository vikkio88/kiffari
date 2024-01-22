<script>
  import { Link } from "svelte-routing";
  import { KATO_API_URL } from "../const";
  import NoteList from "../components/NoteList.svelte";

  let notePromise = fetch(`${KATO_API_URL}/notes?latest=true`).then((resp) =>
    resp.json(),
  );
</script>

<h2>Latest Notes</h2>
{#await notePromise then notes}
  <NoteList {notes} />
{/await}
