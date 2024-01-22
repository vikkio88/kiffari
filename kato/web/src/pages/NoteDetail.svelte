<script>
  import { KATO_API_URL } from "../const";
  import SvelteMarkdown from "svelte-markdown";
  import TagsList from "../components/TagsList.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import { navigate } from "svelte-routing";
  import ConfirmButton from "../components/shared/ConfirmButton.svelte";

  export let id = "";
  let notePromise = fetch(`${KATO_API_URL}/notes/${id}`).then((resp) =>
    resp.json(),
  );

  async function onDelete() {
    const resp = await fetch(`${KATO_API_URL}/notes/${id}`, {
      method: "DELETE",
    });
    if (resp.ok) {
      navigate("/", { replace: true });
    }
  }

  function getDate({ created_at, updated_at }) {
    if (created_at != updated_at) {
      return `updated: ${new Date(updated_at).toLocaleString()}`;
    }

    return new Date(created_at).toLocaleString();
  }

</script>

{#await notePromise then note}
  <div class="note">
    <h2>{note.title}</h2>
    <span class="date">{getDate(note)}</span>
    <div class="body">
      <SvelteMarkdown source={note.body} />
    </div>
    <div class="tags">
      <h3>Tags</h3>
      <TagsList tags={note.tags} />
    </div>
  </div>
  <Controls>
    <button on:click={() => navigate(`/edit-note/${id}`)}>Edit</button>
    <ConfirmButton onConfirmed={onDelete}>Delete</ConfirmButton>
  </Controls>
{/await}

<style>
  .note {
    border: 2px white solid;
    border-radius: 10px;
    padding: 0 2rem;
  }

  .note h2 {
    margin: 1rem 0 0 0;
  }

  .note span.date {
    font-size: smaller;
  }

  .body {
    text-align: left;
  }

  .tags {
    margin-top: 3rem;
    border-top: 2px white solid;
  }

  .tags h3 {
    margin: 0;
    padding: 0 1rem;
    display: inline-block;
    transform: translateY(-50%);
    background-color: #242424;
  }
</style>
