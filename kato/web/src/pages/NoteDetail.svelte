<script>
  import { KATO_API_URL } from "../const";
  import SvelteMarkdown from "svelte-markdown";
  import TagsList from "../components/TagsList.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import { navigate } from "svelte-routing";
  import ConfirmButton from "../components/shared/ConfirmButton.svelte";

  export let id = "";
  let notePromise = fetch(`${KATO_API_URL}/notes/${id}`).then((resp) =>
    resp.json()
  );

  async function onDelete() {
    const resp = await fetch(`${KATO_API_URL}/notes/${id}`, {
      method: "DELETE",
    });
    if (resp.ok) {
      navigate("/", { replace: true });
    }
  }
</script>

{#await notePromise then note}
  <h2>{note.title}</h2>
  <SvelteMarkdown source={note.body} />
  <TagsList tags={note.tags} />
  <Controls>
    <button on:click={() => navigate(`/edit-note/${id}`)}>Edit</button>
    <ConfirmButton onConfirmed={onDelete}>Delete</ConfirmButton>
  </Controls>
{/await}
