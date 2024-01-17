<script>
  import NoteEditor from "../components/NoteEditor.svelte";
  import { KATO_API_URL } from "../const";
  let postPromise = null;
  async function onSave(body, tags) {
    const data = await fetch(`${KATO_API_URL}/notes`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ body, tags }),
    });

    postPromise = data.json();
  }
</script>

{#if postPromise == null}
  <NoteEditor {onSave} />
{:else}
  {#await postPromise}
    <h2>Loading...</h2>
  {:then data}
    <pre>
  {JSON.stringify(data, null, 2)}
    </pre>
  {/await}
{/if}
