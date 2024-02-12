<script>
  import { navigate } from "svelte-routing";
  import NoteEditor from "../components/NoteEditor.svelte";
  import { getNoteDetails, updateNote } from "../libs/api";
  import { protectedRoute } from "../libs/routes";
  protectedRoute();

  export let id = "";
  let notePromise = getNoteDetails(id).then((resp) => resp.json());
  let putPromise = null;

  async function onSave(note) {
    putPromise = updateNote(id, { id, ...note });

    const data = await putPromise;
    if (data.status == 200) {
      navigate(`/notes/${id}`, { replace: true });
    }
  }
</script>

{#await notePromise then note}
  <NoteEditor
    title={note.title}
    text={note.body}
    tags={note.tags}
    dueDate={Boolean(note.due_date) ? new Date(note.due_date) : null}
    {onSave}
  />
{/await}
