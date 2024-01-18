<script>
  import { KATO_API_URL } from "../const";

  let notePromise = fetch(`${KATO_API_URL}/notes?latest=true`).then((resp) =>
    resp.json()
  );
</script>

{#await notePromise then notes}
  <ul>
    {#each notes as note}
      <li>{note.title} - {new Date(note.created_at).toLocaleDateString()}</li>
    {/each}
  </ul>
{/await}
