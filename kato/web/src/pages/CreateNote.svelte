<script>
  import { navigate } from "svelte-routing";
  import NoteEditor from "../components/NoteEditor.svelte";
  import { KATO_API_URL } from "../const";
  import { protectedRoute } from "../libs";
  protectedRoute();
  
  let postPromise = null;
  async function onSave(title, body, tags) {
    postPromise = fetch(`${KATO_API_URL}/notes`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title, body, tags }),
    });

    const data = await postPromise;
    if (data.status == 200) {
      const { id } = await data.json();
      navigate(`/notes/${id}`, { replace: true });
    }
  }
</script>

<h2>Add Note</h2>
{#if postPromise == null}
  <NoteEditor {onSave} />
{:else}
  <h2>Creating...</h2>
{/if}

<style>
  h2 {
    margin: 0;
  }
</style>
