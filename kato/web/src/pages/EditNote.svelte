<script>
  import { navigate } from "svelte-routing";
  import NoteEditor from "../components/NoteEditor.svelte";
  import { KATO_API_URL } from "../const";
  import { protectedRoute } from "../libs";
  protectedRoute();
  
  export let id = "";
  let notePromise = fetch(`${KATO_API_URL}/notes/${id}`).then((resp) =>
    resp.json(),
  );
  let putPromise = null;

  async function onSave(title, body, tags) {
    putPromise = fetch(`${KATO_API_URL}/notes/${id}`, {
      method: "PUT",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ id, title, body, tags }),
    });

    const data = await putPromise;
    if (data.status == 200) {
      navigate(`/notes/${id}`, { replace: true });
    }
  }
</script>

{#await notePromise then note}
  <NoteEditor title={note.title} text={note.body} tags={note.tags} {onSave} />
{/await}
