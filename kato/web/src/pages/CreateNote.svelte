<script>
  import { navigate } from "svelte-routing";
  import NoteEditor from "../components/NoteEditor.svelte";
  import { protectedRoute } from "../libs/routes";
  import { createNote } from "../libs/api";
  protectedRoute();
  
  let postPromise = null;
  async function onSave(title, body, tags) {
    postPromise = createNote({ title, body, tags })

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
