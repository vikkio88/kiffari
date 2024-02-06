<script>
  import SvelteMarkdown from "svelte-markdown";
  import TagsList from "../components/TagsList.svelte";
  import Controls from "../components/shared/Controls.svelte";
  import { navigate } from "svelte-routing";
  import ConfirmButton from "../components/shared/ConfirmButton.svelte";
  import Plugin from "../components/Plugin.svelte";
  import { protectedRoute } from "../libs/routes";
  import {
    getNoteDetails,
    deleteNote,
    parseOrThrow,
    catchLogout,
  } from "../libs/api";
  protectedRoute();

  export let id = "";
  let notePromise = getNoteDetails(id).then(parseOrThrow).catch(catchLogout);

  async function onDelete() {
    const resp = await deleteNote(id);
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
      <Plugin config={note.body} />
      <SvelteMarkdown source={note.body} />
    </div>
    <div class="tags">
      <h3>Tags</h3>
      <div class="tagList">
        <TagsList tags={note.tags} />
      </div>
    </div>
  </div>
  <Controls>
    <button on:click={() => navigate(`/edit-note/${id}`)}>Edit</button>
    <ConfirmButton onConfirmed={onDelete}>Delete</ConfirmButton>
  </Controls>
{/await}

<style>
  .note {
    /* border: 2px white solid; */
    border-radius: 10px;
    padding: 0 2rem;
  }

  .note h2 {
    margin: 1rem 0 0 0;
  }

  .note span.date {
    font-size: smaller;
    color: #a3a3a3;
  }

  .body {
    border-radius: 10px;
    text-align: left;
    padding: 3rem;
    margin: 1em 0;
    min-height: 50vh;
  }

  .tags {
    border-top: 2px #a3a3a3 solid;
  }

  .tagList {
    margin: 1rem 0;
    display: flex;
    flex-direction: row;
    padding-bottom: 0.5rem;
  }

  .tags h3 {
    margin: 0;
    padding: 0 1rem;
    display: inline-block;
    transform: translateY(-50%);
    background-color: #242424;
    color: #a3a3a3;
  }
</style>
