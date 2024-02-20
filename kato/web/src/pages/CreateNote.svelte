<script>
  import { navigate } from "svelte-routing";
  import NoteEditor from "../components/NoteEditor.svelte";
  import { protectedRoute } from "../libs/routes";
  import { catchLogout, createNote } from "../libs/api";
  import Error from "../components/shared/ErrorToast.svelte";
  protectedRoute();

  let error = null;

  let postPromise = null;
  async function onSave(note) {
    postPromise = createNote(note);

    const resp = await postPromise;
    if (resp.status == 401) {
      catchLogout();
      return;
    }

    if (resp.status == 400) {
      postPromise = null;
      error = `Could not create.`;
    }

    if (resp.status == 401) {
      catchLogout();
    }

    if (resp.status == 200) {
      const { id } = await resp.json();
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

{#if Boolean(error)}
  <Error {error} onDismiss={() => (error = null)} />
{/if}

<style>
  h2 {
    margin: 0;
  }
</style>
